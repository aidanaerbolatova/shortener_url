package main

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"log/slog"
	"shortener-link/internal/config"
	"shortener-link/internal/repository"
	"shortener-link/internal/resthttp"
	"shortener-link/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Error(err)
		return
	}
}

func run() error {
	cfg, err := config.LoadConfig("./.env")
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	db, err := repository.NewConnection(&cfg)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	repo := repository.NewRepository(db)
	srv, err := service.NewLinkServiceImpl(repo)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	cancelCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		if err := srv.DeleteExpiredShortenerLinks(cancelCtx, 24*time.Hour); err != nil {
			slog.Error("Error deleting expired links: ", err)
		}
	}()

	h := resthttp.NewHandler(srv)
	if err := resthttp.RunServer(cfg, h.Register()); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
