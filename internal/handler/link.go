package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	models2 "shortener-link/internal/models"
)

func (h Handler) Create(c *fiber.Ctx) error {
	var url models2.LinkRequest
	if err := c.BodyParser(&url); err != nil {
		c.JSON(models2.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return err
	}
	_, err := h.service.Create(c.Context(), url.Link)
	if err != nil {
		c.JSON(models2.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h Handler) GetAll(c *fiber.Ctx) error {

	links, err := h.service.GetAll(c.Context())
	if err != nil {
		c.JSON(models2.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return err
	}

	c.Status(http.StatusOK).JSON(links)
	return nil
}

func (h Handler) GetByShortener(c *fiber.Ctx) error {

	link := c.Query("link")

	fullLink, err := h.service.GetByShortenerLink(c.Context(), link)
	if err != nil {
		if err == models2.ErrLinkNotFound {
			c.JSON(models2.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "сокращённая ссылка не найдена",
			})
			return models2.ErrLinkNotFound
		}
		c.JSON(models2.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return err
	}

	c.Status(http.StatusOK).JSON(fullLink)
	return nil

}

func (h Handler) Delete(c *fiber.Ctx) error {

	link := c.Query("link")

	if err := h.service.DeleteShortenerLink(c.Context(), link); err != nil {
		c.JSON(models2.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (h Handler) GetStats(c *fiber.Ctx) error {

	link := c.Query("link")

	stat, lastVisitTime, err := h.service.GetStatsByShortenerLink(c.Context(), link)
	if err != nil {
		if err == models2.ErrLinkNotFound {
			c.JSON(models2.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "сокращённая ссылка не найдена",
			})
			return models2.ErrLinkNotFound
		}
		c.JSON(models2.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return err
	}

	c.Status(http.StatusOK).JSON(models2.GetStatsResponse{
		StatsCount:    stat,
		LastVisitTime: lastVisitTime,
	})
	return nil
}
