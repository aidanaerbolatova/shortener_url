package service

import (
	"context"
	"errors"
	"shortener-link/internal/models"
	"shortener-link/internal/repository"
	"time"
)

type LinkService interface {
	Create(ctx context.Context, link string) (string, error)
	GetAll(ctx context.Context) ([]models.Link, error)
	GetByShortenerLink(ctx context.Context, shortenerLink string) (string, error)
	DeleteShortenerLink(ctx context.Context, shortenerLink string) error
	GetStatsByShortenerLink(ctx context.Context, shortenerLink string) (int, time.Time, error)
	DeleteExpiredShortenerLink(ctx context.Context) error
	DeleteExpiredShortenerLinks(ctx context.Context, interval time.Duration) error
}

type Service struct {
	repo repository.LinkRepository
}

func NewLinkServiceImpl(repo repository.LinkRepository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("link repository is nil")
	}
	return &Service{
		repo: repo,
	}, nil
}

func (s *Service) Create(ctx context.Context, link string) (string, error) {

	shortenerLink, err := GenerateShortenerLink(link)
	if err != nil {
		return "", err
	}

	linkID, err := s.repo.Create(ctx, models.Link{
		FullLink:      link,
		ShortenerLink: shortenerLink,
		Visits:        0,
		CreatedAt:     time.Now().UTC(),
	})
	if err != nil {
		return "", err
	}

	return linkID, nil
}

func (s *Service) GetAll(ctx context.Context) ([]models.Link, error) {

	links, err := s.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func (s *Service) GetByShortenerLink(ctx context.Context, shortenerLink string) (string, error) {

	linkInfo, err := s.repo.GetByShortenerLink(ctx, shortenerLink)
	if err != nil {
		return "", nil
	}

	if err = s.repo.UpdateVisitorsByShortenerLink(ctx, shortenerLink); err != nil {
		return "", err
	}

	return linkInfo.FullLink, nil
}

func (s *Service) DeleteShortenerLink(ctx context.Context, shortenerLink string) error {

	if err := s.repo.DeleteByShortenerLink(ctx, shortenerLink); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetStatsByShortenerLink(ctx context.Context, shortenerLink string) (int, time.Time, error) {

	link, err := s.repo.GetByShortenerLink(ctx, shortenerLink)
	if err != nil {
		return 0, time.Time{}, err
	}

	return link.Visits, link.LastVisitTime, nil
}

func (s *Service) DeleteExpiredShortenerLink(ctx context.Context) error {

	if err := s.repo.DeleteExpiredShortenerLink(ctx); err != nil {
		return err
	}

	return nil
}
