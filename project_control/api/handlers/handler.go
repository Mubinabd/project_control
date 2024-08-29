package handlers

import (
	s "github.com/Mubinabd/project_control/internal/usecase/service"
)

type Handlers struct {
	Group   *s.GroupService
	Private *s.PrivateService
}

func NewHandler(group *s.GroupService, private *s.PrivateService) *Handlers {
	return &Handlers{
		Group:   group,
		Private: private,
	}
}
