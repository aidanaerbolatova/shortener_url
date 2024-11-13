package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"shortener-link/models"
)

func Create(c *fiber.Ctx) {
	var url models.LinkRequest
	if err := c.BodyParser(&url); err != nil {
		c.JSON(models.ErrorResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		})
	}
}

func GetAll(c *fiber.Ctx) {}

func GetByShortener(c *fiber.Ctx) {
	_ = c.Query("link")
	//link:
}

func Delete(c *fiber.Ctx) {
	_ = c.Query("link")
	//link:
}

func GetStats(c *fiber.Ctx) {
	_ = c.Query("link")
	//link:
}
