package main

import (
	"os"

	"github.com/Royalsspirit/cisco-api/database"
	"github.com/Royalsspirit/cisco-api/internal/api"
	"github.com/Royalsspirit/cisco-api/internal/logger"
	"github.com/Royalsspirit/cisco-api/internal/validator"
)

func main() {
	serverName := "cisco-api"
	// Init database
	db := database.NewDB(os.Getenv("DB"))

	// Init logger
	logger := logger.NewLogger(&logger.Config{
		Name:        serverName,
		Environment: "develop",
	})

	// Init validator
	v := validator.NewValidator()

	// TODO: Get variable from the Environment
	server := api.NewServer(&api.ServerConfig{
		Environment: "develop",
		Port:        "8080",
		Logger:      logger,
		DB:          db,
		Validator:   v,
	})

	server.Run()
}
