package repository

import (
	"context"
	"database/sql"
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
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, linkInfo models.Link) (string, error) {

	query := `INSERT into link_info
    			(link_id, full_link, shortner_link, visits, last_visit_time, created_at) 
				values (CAST($1 as integer),$2, $3, $4, $5, $6) 
				returning id `

	var linkID string

	row := r.db.QueryRowContext(ctx, query, linkInfo.FullLink, linkInfo.ShortenerLink, linkInfo.Visits, linkInfo.LastVisitTime, linkInfo.CreatedAt)
	if err := row.Scan(&linkID); err != nil {
		return "", err
	}
	return linkID, nil
}

func (r *Repository) Get(ctx context.Context) ([]models.Link, error) {

	query := `SELECT link_id, full_link, shortener_link, visits, last_visit_time,  created_at 
			  FROM link_info`

	rows, err := r.db.QueryContext(ctx, query)
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

	query := `UPDATE link_info
			  SET visits = visits + 1, xf = $1
			  WHERE shortener_link = $2`

	_, err := r.db.ExecContext(ctx, query, time.Now().UTC(), shortenerLink)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteByShortenerLink(ctx context.Context, shortenerLink string) error {

	query := `DELETE FROM link_info WHERE shortener_link = $1`

	_, err := r.db.ExecContext(ctx, query, shortenerLink)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByShortenerLink(ctx context.Context, shortenerLink string) (models.Link, error) {

	query := `SELECT link_id, full_link, shortener_link, visits, last_visit_time, created_at 
			  FROM link_info 
			  WHERE shortener_link = $1`

	if _, err := r.db.Prepare(query); err != nil {
		return models.Link{}, err
	}

	var link models.Link
	if err := r.db.QueryRowContext(ctx, query, shortenerLink).Scan(
		&link.ID,
		&link.FullLink,
		&link.ShortenerLink,
		&link.Visits,
		&link.LastVisitTime,
		&link.CreatedAt,
	); err != nil {
		if err == models.ErrLinkNotFound {
			return models.Link{}, models.ErrLinkNotFound
		}
		return models.Link{}, err
	}
	return link, nil
}

func (r *Repository) DeleteExpiredShortenerLink(ctx context.Context) error {

	query := `DELETE FROM link_info  WHERE created_at < NOW() - INTERVAL 30 DAY`

	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
