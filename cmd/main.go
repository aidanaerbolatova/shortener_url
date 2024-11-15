package main

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"log/slog"
	"shortener-link/internal/config"
	"shortener-link/internal/handler"
	"shortener-link/internal/repository"
	"shortener-link/internal/server"
	"shortener-link/internal/service"
	"time"
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
	if err := srv.DeleteExpiredShortenerLinks(cancelCtx, 24*time.Hour); err != nil {
		slog.Error(err.Error())
		return err
	}

	h := handler.NewHandler(srv)
	if err := server.RunServer(cfg, h.Register()); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
