package main

import (
	"github.com/irinaponzi/package-tracker/internal/config"
	"log"
)

// @title Package Tracker API
// @version 1.0
// @description API to track packages and manage batches
// @host localhost:8080
// @BasePath /api/v1
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// - app setup
	app := config.NewAppDefault()
	app.SetDependencies()
	app.SetMappings()

	return app.Run()
}
