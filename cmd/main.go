package main

import (
	"github.com/gofiber/fiber/v2/log"
	"shortener-link/internal/config"
	handler2 "shortener-link/internal/handler"
	"shortener-link/internal/repository"
	"shortener-link/internal/server"
	service2 "shortener-link/internal/service"
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
		log.Trace(err)
		return err
	}

	db, err := repository.NewConnection(&cfg)
	if err != nil {
		log.Trace(err)
		return err
	}
	repository := repository.NewRepository(db)

	service, err := service2.NewLinkServiceImpl(repository)
	if err != nil {
		log.Trace(err)
		return err
	}

	handler := handler2.NewHandler(service)

	if err := server.RunServer(cfg, handler.Register()); err != nil {
		log.Trace(err)
		return err
	}

	return nil
}
