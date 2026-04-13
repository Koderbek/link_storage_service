package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Koderbek/link_storage_service/internal/model"
	"github.com/jmoiron/sqlx"
)

type LinkRepository interface {
	Create(url, code string) error
	Link(code string) (*model.Link, error)
	Links(limit, offset int) ([]model.Link, error)
	Delete(code string) error
	Stats(code string) (*model.LinkStats, error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Create(url, code string) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (original_url, short_code) VALUES ($1, $2);",
		LinkTable,
	)

	_, err := repo.db.Exec(query, url, code)

	return err
}

func (repo *Repository) Link(code string) (*model.Link, error) {
	var link model.Link
	query := fmt.Sprintf(
		"SELECT original_url, visits FROM %s WHERE short_code = $1",
		LinkTable,
	)

	if err := repo.db.Get(&link, query, code); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &link, nil
}

func (repo *Repository) Links(limit, offset int) ([]model.Link, error) {
	var links []model.Link
	query := fmt.Sprintf(
		"SELECT original_url, visits FROM %s ORDER BY id LIMIT $1 OFFSET $2",
		LinkTable,
	)

	if err := repo.db.Select(&links, query, limit, offset); err != nil {
		return nil, err
	}

	return links, nil
}

func (repo *Repository) Delete(code string) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE short_code = $1",
		LinkTable,
	)

	_, err := repo.db.Exec(query, code)

	return err
}

func (repo *Repository) Stats(code string) (*model.LinkStats, error) {
	var stats model.LinkStats
	query := fmt.Sprintf(
		"SELECT original_url, short_code, created_at, visits FROM %s WHERE short_code = $1",
		LinkTable,
	)

	if err := repo.db.Get(&stats, query, code); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &stats, nil
}
