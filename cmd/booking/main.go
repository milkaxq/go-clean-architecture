package main

import (
	"booking/internal/infrastrucuture/configuration"
	"booking/internal/infrastrucuture/database"
	"booking/internal/infrastrucuture/logging"
	"context"
)

func main() {
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

}
