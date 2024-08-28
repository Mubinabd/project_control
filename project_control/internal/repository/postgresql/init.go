package postgresql

import (
	"database/sql"
	"github.com/Mubinabd/project_control/internal/repository"
)

type Storage struct {
	PrivateS repository.PrivateI
	GroupS   repository.GroupI
	db       *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		PrivateS: NewPrivateManager(db),
		GroupS:   NewGroupManager(db),
		db:       db,
	}
}
func (s *Storage) Group() repository.GroupI {
	if s.GroupS == nil {
		s.GroupS = NewGroupManager(s.db)
	}
	return s.GroupS
}
func (s *Storage) Private() repository.PrivateI {
	if s.PrivateS == nil {
		s.PrivateS = NewPrivateManager(s.db)
	}
	return s.PrivateS
}
