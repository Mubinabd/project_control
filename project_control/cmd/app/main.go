package main

import (
	"github.com/Mubinabd/project_control/internal/app"
	"github.com/Mubinabd/project_control/internal/pkg/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}

