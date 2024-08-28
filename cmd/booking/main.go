package main

import (
	"booking/internal/infrastrucuture/configuration"
	"booking/internal/infrastrucuture/logging"
	"booking/internal/infrastrucuture/persistence"
	"booking/internal/infrastrucuture/server"
)

func main() {
	logger := logging.NewLogger()

	config, err := configuration.LoadConfig(logger)
	if err != nil {
		logger.Fatal("Failed to load configuration", map[string]interface{}{
			"error": err.Error(),
		})
	}

	dbConn := persistence.NewDBConnection(config, logger)

	app := server.NewServer(logger, config, dbConn)

	if err := app.StartServer(); err != nil {
		logger.Fatal("Server encountered a critical error", map[string]interface{}{
			"error": err.Error(),
		})
	}
}
