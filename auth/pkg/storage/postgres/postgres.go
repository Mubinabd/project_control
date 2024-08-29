package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Mubinabd/project_control/pkg/storage"


	"github.com/Mubinabd/project_control/internal/config"
)

type Storage struct {
	Db      *sql.DB
	OrderDb *sql.DB
	AuthS   storage.AuthI
	UserS   storage.UserI
}

func NewPostgresStorage(config *config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST,
		config.DB_USER,
		config.DB_NAME,
		config.DB_PASSWORD,
		config.DB_PORT,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{
		Db:    db,
		AuthS: NewAuthRepo(db),
		UserS: NewUserRepo(db),
	}, nil
}
