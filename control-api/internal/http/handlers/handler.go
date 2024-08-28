package handlers

import (
	"github.com/Mubinabd/project_control/internal/grpc"
	"github.com/Mubinabd/project_control/internal/pkg/logger"
)

type Handler struct {
	Clients grpc.Clients
	Logger  *logger.Logger
}

func NewHandler(clients grpc.Clients, logger *logger.Logger) *Handler {
	return &Handler{Clients: clients, Logger: logger}
}
