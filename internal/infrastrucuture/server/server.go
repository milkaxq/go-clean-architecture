package server

import (
	"booking/internal/application/interfaces"
	"booking/internal/domain/repositories"
	"booking/internal/infrastrucuture/configuration"
	"booking/internal/interfaces/http/controllers"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Logger interfaces.Logger
	Config *configuration.Configs
	DBConn repositories.DBConnection
	Router *gin.Engine
}

func NewServer(logger interfaces.Logger, config *configuration.Configs, dbConn repositories.DBConnection) *Server {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	return &Server{
		Logger: logger,
		Config: config,
		DBConn: dbConn,
		Router: router,
	}
}

func (s *Server) StartServer() error {
	err := s.DBConn.Connect(context.Background())
	if err != nil {
		s.Logger.Error("Failed to connect to database", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}

	controllers.RegisterRoutes(s.Router)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":..%d", s.Config.Server.Port),
		Handler: s.Router.Handler(),
	}

	s.Logger.Info("Server start running", map[string]interface{}{"server port": srv.Addr})

	errChan := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Error("Server failed to start", map[string]interface{}{
				"error": err.Error(),
			})
			errChan <- err
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		s.Logger.Info("Shutdown Server ...", map[string]interface{}{})
	case err := <-errChan:
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Error("Server forced to shutdown", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}

	if err := s.DBConn.Close(ctx); err != nil {
		s.Logger.Error("Failed to close database connection", map[string]interface{}{
			"error": err.Error(),
		})
		return err
	}

	s.Logger.Info("Server exiting", map[string]interface{}{})

	return nil
}
