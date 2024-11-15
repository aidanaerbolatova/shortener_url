package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"shortener-link/internal/service"
	"time"
)

type Handler struct {
	service service.LinkService
}

func NewHandler(service service.LinkService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register() *fiber.App {

	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,PUT,POST,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 1 * time.Minute,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ping")
	})

	app.Post("/shortener", h.Create)
	app.Get("/shortener", h.GetAll)
	app.Get("/:link<string>", h.GetByShortener)
	app.Delete("/:link<string>", h.Delete)
	app.Get("/stats/:link<string>", h.GetStats)

	return app
}
