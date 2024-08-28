package app

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/Mubinabd/project_control/internal/grpc"
	"github.com/Mubinabd/project_control/internal/http"
	"github.com/Mubinabd/project_control/internal/http/handlers"
	"github.com/Mubinabd/project_control/internal/pkg/config"
	"github.com/Mubinabd/project_control/internal/pkg/logger"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Run(cfg config.Config) {
	logger := logger.NewLogger(basepath, cfg.LogPath)
	clients, err := grpc.NewClients(&cfg)
	if err != nil {
		logger.ERROR.Println("Failed to create gRPC clients", err)
		log.Fatal(err)
		return
	}

	// make handler
	h := handlers.NewHandler(*clients, logger)

	// make gin
	router := http.NewGin(h)

	// start server
	router.Run(":8080")
}
