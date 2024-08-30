package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Mubinabd/project_control/internal/repository"
	"github.com/Mubinabd/project_control/pkg/config"
)

type Storage struct {
	Db       *sql.DB
	GroupS   repository.GroupI
	PrivateS repository.PrivateI
	AuthS    repository.AuthI
	UserS    repository.UserI
}

func New(cfg *config.Config) (*Storage, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Println("can't connect to db: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{
		Db:       db,
		GroupS:   NewGroupManager(db),
		PrivateS: NewPrivateManager(db),
		AuthS:    NewAuthRepo(db),
		UserS:    NewUserRepo(db),
	}, nil
}
func (s *Storage) Group() repository.GroupI {
	if s.GroupS == nil {
		s.GroupS = NewGroupManager(s.Db)
	}
	return s.GroupS
}

func (s *Storage) Private() repository.PrivateI {
	if s.PrivateS == nil {
		s.PrivateS = NewPrivateManager(s.Db)
	}
	return s.PrivateS
}
func (s *Storage) Auth() repository.AuthI {
	if s.AuthS == nil {
		s.AuthS = NewAuthRepo(s.Db)
	}
	return s.AuthS
}

func (s *Storage) User() repository.UserI {
	if s.UserS == nil {
		s.UserS = NewUserRepo(s.Db)
	}
	return s.UserS
}
