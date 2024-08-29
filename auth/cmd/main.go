package main

import (
	"github.com/Mubinabd/project_control/internal/config"
	"github.com/Mubinabd/project_control/pkg/app"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration from environment variables or a file.
	cfg := config.Load()

	// Run the application with the loaded configuration.
	app.Run(&cfg)
}
