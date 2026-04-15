package database

import (
	"fmt"
	"github.com/Koderbek/link_storage_service/internal/model"
	"github.com/jmoiron/sqlx"
)

type LinkRepository interface {
	Create(url string) (uint, error)
	Link(id uint) (*model.Link, error)
	Links(limit, offset uint) ([]model.Link, error)
	UpdateVisits(id, visits uint) error
	Delete(id uint) error
	Stats(id uint) (*model.LinkStats, error)
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Create(url string) (uint, error) {
	var id uint
	query := fmt.Sprintf(
		"INSERT INTO %s (original_url) VALUES ($1) RETURNING id;",
		LinkTable,
	)

	err := repo.db.QueryRowx(query, url).Scan(&id)

	return id, err
}

func (repo *Repository) Link(id uint) (*model.Link, error) {
	var link model.Link
	query := fmt.Sprintf(
		"SELECT id, original_url, visits FROM %s WHERE id = $1",
		LinkTable,
	)

	if err := repo.db.Get(&link, query, id); err != nil {
		return nil, err
	}

	return &link, nil
}

func (repo *Repository) Links(limit, offset uint) ([]model.Link, error) {
	var links []model.Link
	query := fmt.Sprintf(
		"SELECT id, original_url, visits FROM %s ORDER BY id LIMIT $1 OFFSET $2",
		LinkTable,
	)

	if err := repo.db.Select(&links, query, limit, offset); err != nil {
		return nil, err
	}

	return links, nil
}

func (repo *Repository) UpdateVisits(id, visits uint) error {
	query := fmt.Sprintf(
		"UPDATE %s SET visits = $1 WHERE id = $2",
		LinkTable,
	)

	if _, err := repo.db.Exec(query, visits, id); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Delete(id uint) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1",
		LinkTable,
	)

	_, err := repo.db.Exec(query, id)

	return err
}

func (repo *Repository) Stats(id uint) (*model.LinkStats, error) {
	var stats model.LinkStats
	query := fmt.Sprintf(
		"SELECT id, original_url, created_at, visits FROM %s WHERE id = $1",
		LinkTable,
	)

	if err := repo.db.Get(&stats, query, id); err != nil {
		return nil, err
	}

	return &stats, nil
}
