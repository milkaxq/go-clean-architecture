package server

import (
	"booking/internal/infrastrucuture/configuration"
	"booking/internal/infrastrucuture/database"
	"booking/internal/infrastrucuture/logging"
	"booking/internal/interfaces/http/controllers"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	logger := logging.NewLogger()

	config, err := configuration.LoadConfig(logger)
	if err != nil {
		logger.Fatal("Failed to load configuration", map[string]interface{}{
			"error": err.Error(),
		})
	}

	dbConn := database.NewDBConnection(config, logger)

	err = dbConn.Connect(context.Background())
	if err != nil {
		logger.Fatal("Failed to connect to database", map[string]interface{}{
			"error": err.Error(),
		})
	}

	r := gin.Default()

	r.SetTrustedProxies(nil)

	controllers.RegisterRoutes(r)

	if err := r.Run(fmt.Sprintf(":%d", config.Server.Port)); err != nil {
		logger.Fatal("Failed to start server", map[string]interface{}{"error": err.Error()})
	}
}
