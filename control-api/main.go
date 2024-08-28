package main

import (
	"github.com/Mubinabd/project_control/internal/app"
	"github.com/Mubinabd/project_control/internal/pkg/config"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}

