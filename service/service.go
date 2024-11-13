package service

import (
	"context"
	"errors"
	"shortener-link/models"
	"shortener-link/repository"
	"time"
)

type LinkService interface {
	Create(ctx context.Context, link string) (string, error)
	GetAll(ctx context.Context) ([]models.Link, error)
	GetByShortenerLink(ctx context.Context, shortenerLink string) (string, error)
	DeleteShortenerLink(ctx context.Context, shortenerLink string) error
	GetStatsByShortenerLink(ctx context.Context, shortenerLink string) (int, error)
}

type LinkServiceImpl struct {
	repo repository.LinkRepository
}

func NewLinkServiceImpl(repo repository.LinkRepository) (*LinkServiceImpl, error) {
	if repo == nil {
		return nil, errors.New("link repository is nil")
	}
	return &LinkServiceImpl{
		repo: repo,
	}, nil
}

func (s *LinkServiceImpl) Create(ctx context.Context, link string) (string, error) {

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

func (s *LinkServiceImpl) GetAll(ctx context.Context) ([]models.Link, error) {

	links, err := s.repo.Get(ctx)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func (s *LinkServiceImpl) GetByShortenerLink(ctx context.Context, shortenerLink string) (string, error) {

	linkInfo, err := s.repo.GetByShortenerLink(ctx, shortenerLink)
	if err != nil {
		return "", nil
	}

	if err = s.repo.UpdateVisitorsByShortenerLink(ctx, shortenerLink); err != nil {
		return "", err
	}

	return linkInfo.FullLink, nil
}

func (s *LinkServiceImpl) DeleteShortenerLink(ctx context.Context, shortenerLink string) error {

	if err := s.repo.DeleteByShortenerLink(ctx, shortenerLink); err != nil {
		return err
	}

	return nil
}

func (s *LinkServiceImpl) GetStatsByShortenerLink(ctx context.Context, shortenerLink string) (int, error) {

	link, err := s.repo.GetByShortenerLink(ctx, shortenerLink)
	if err != nil {
		return 0, err
	}

	return link.Visits, nil
}
