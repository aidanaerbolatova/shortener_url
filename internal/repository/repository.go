package repository

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	"shortener-link/internal/models"
	"time"
)

type LinkRepository interface {
	Create(ctx context.Context, linkInfo models.Link) (string, error)
	Get(ctx context.Context) ([]models.Link, error)
	GetByShortenerLink(ctx context.Context, shortenerLink string) (models.Link, error)
	UpdateVisitorsByShortenerLink(ctx context.Context, shortenerLink string) error
	DeleteByShortenerLink(ctx context.Context, shortenerLink string) error
	DeleteExpiredShortenerLink(ctx context.Context) error
}

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, linkInfo models.Link) (string, error) {
	var linkID string

	err := r.db.QueryRow(ctx, InsertQuery, linkInfo.FullLink, linkInfo.ShortenerLink, linkInfo.Visits, linkInfo.LastVisitTime, linkInfo.CreatedAt).Scan(&linkID)
	if err != nil {
		return "", err
	}

	return linkID, nil
}

func (r *Repository) Get(ctx context.Context) ([]models.Link, error) {
	rows, err := r.db.Query(ctx, GetQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]models.Link, 0, 2)
	for rows.Next() {
		var link models.Link
		if err := rows.Scan(&link.ID, &link.FullLink, &link.ShortenerLink, &link.Visits, &link.LastVisitTime, &link.CreatedAt); err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}

func (r *Repository) UpdateVisitorsByShortenerLink(ctx context.Context, shortenerLink string) error {
	_, err := r.db.Exec(ctx, UpdateVisitorsByShortenerLinkQuery, time.Now().UTC(), shortenerLink)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteByShortenerLink(ctx context.Context, shortenerLink string) error {
	_, err := r.db.Exec(ctx, DeleteByShortenerLinkQueryQuery, shortenerLink)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByShortenerLink(ctx context.Context, shortenerLink string) (models.Link, error) {
	var link models.Link
	if err := r.db.QueryRow(ctx, GetByShortenerLinkQuery, shortenerLink).Scan(
		&link.ID,
		&link.FullLink,
		&link.ShortenerLink,
		&link.Visits,
		&link.LastVisitTime,
		&link.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return models.Link{}, models.ErrLinkNotFound
		}
		return models.Link{}, err
	}
	return link, nil
}

func (r *Repository) DeleteExpiredShortenerLink(ctx context.Context) error {
	_, err := r.db.Exec(ctx, DeleteExpiredShortenerLinkQuery)
	if err != nil {
		return err
	}

	return nil
}
