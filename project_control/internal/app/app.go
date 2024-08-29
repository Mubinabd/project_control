package app

import (
	a "github.com/Mubinabd/project_control/api"
	"github.com/Mubinabd/project_control/api/handlers"
	"github.com/Mubinabd/project_control/internal/repository/postgresql"
	s "github.com/Mubinabd/project_control/internal/usecase/service"
	"github.com/Mubinabd/project_control/pkg/config"
	"golang.org/x/exp/slog"
)

func Run(cfg *config.Config) {
	// Postgres Connection
	db, err := postgresql.New(cfg)
	if err != nil {
		slog.Error("can't connect to db: %v", err)
		return
	}
	defer db.Db.Close()
	slog.Info("Connected to Postgres")

	groupService := s.NewGroupService(db)
	privateService := s.NewPrivateService(db)

	// HTTP Server
	h := handlers.NewHandler(groupService, privateService)

	router := a.NewGin(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.GRPCPort); err != nil {
		slog.Error("can't start server: %v", err)
	}

	slog.Info("REST server started on port %s", cfg.GRPCPort)
}
