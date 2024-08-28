package server

import (
	"booking/internal/application/interfaces"
	"booking/internal/domain/repositories"
	"booking/internal/infrastrucuture/configuration"
	"booking/internal/interfaces/http/controllers"
	"context"
	"fmt"
	"log"
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
		Addr:    fmt.Sprintf(":%d", s.Config.Server.Port),
		Handler: s.Router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Error("listen: %s\n", map[string]interface{}{
				"err": err.Error(),
			})
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
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
