package resthttp

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"net"
	"net/http"
	"shortener-link/internal/config"
	"time"
)

func RunServer(cfg config.Config, c *fiber.App) error {
	listener, err := net.Listen("tcp", cfg.Host+":"+cfg.Port)
	if err != nil {
		return err
	}
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		if err := c.Listener(listener); err != nil {
			log.Fatalf("Fiber listener error: %v", err)
		}
	}()

	return server.Serve(listener)
}
