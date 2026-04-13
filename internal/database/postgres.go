package database

import (
	"github.com/Koderbek/link_storage_service/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const LinkTable = "link"

func NewPostgresDb(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DbConnection)
	if err != nil {
		return nil, err
	}

	return db, nil
}
