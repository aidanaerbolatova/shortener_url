package repository

import (
	"context"
	"database/sql"
	"shortener-link/models"
)

type LinkRepository interface {
	Create(ctx context.Context, linkInfo models.Link) (string, error)
	Get(ctx context.Context) ([]models.Link, error)
	GetByShortenerLink(ctx context.Context, shortenerLink string) (models.Link, error)
	UpdateVisitorsByShortenerLink(ctx context.Context, shortenerLink string) error
	DeleteByShortenerLink(ctx context.Context, shortenerLink string) error
}

type LinkPostgresRepository struct {
	db *sql.DB
}

func NewLinkPostgresRepository(db *sql.DB) *LinkPostgresRepository {
	return &LinkPostgresRepository{
		db: db,
	}
}

func (r *LinkPostgresRepository) Create(ctx context.Context, linkInfo models.Link) (string, error) {

	query := `INSERT into link_info
    			(link_id, full_link, shortner_link, visits, created_at) 
				values (CAST($1 as integer),$2, $3, $4, $5) 
				returning id `

	var linkID string

	row := r.db.QueryRowContext(ctx, query, linkInfo.FullLink, linkInfo.ShortenerLink, linkInfo.Visits, linkInfo.CreatedAt)
	if err := row.Scan(&linkID); err != nil {
		return "", err
	}
	return linkID, nil
}

func (r *LinkPostgresRepository) Get(ctx context.Context) ([]models.Link, error) {

	query := `SELECT link_id, full_link, shortener_link, visits, created_at 
			  FROM link_info`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]models.Link, 0, 2)
	for rows.Next() {
		var link models.Link
		if err := rows.Scan(&link.ID, &link.FullLink, &link.ShortenerLink, &link.Visits, &link.CreatedAt); err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}

func (r *LinkPostgresRepository) UpdateVisitorsByShortenerLink(ctx context.Context, shortenerLink string) error {

	query := `UPDATE link_info
			  SET visits = visits + 1
			  WHERE shortener_link = $2`

	_, err := r.db.ExecContext(ctx, query, shortenerLink)
	if err != nil {
		return err
	}
	return nil
}

func (r *LinkPostgresRepository) DeleteByShortenerLink(ctx context.Context, shortenerLink string) error {

	query := `DELETE FROM link_info WHERE shortener_link = $1`

	_, err := r.db.ExecContext(ctx, query, shortenerLink)
	if err != nil {
		return err
	}
	return nil
}

func (r *LinkPostgresRepository) GetByShortenerLink(ctx context.Context, shortenerLink string) (models.Link, error) {

	query := `SELECT link_id, full_link, shortener_link, visits, created_at 
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
		&link.CreatedAt,
	); err != nil {
		if err == models.ErrLinkNotFound {
			return models.Link{}, models.ErrLinkNotFound
		}
		return models.Link{}, err
	}
	return link, nil
}
