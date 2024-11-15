package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"net/http"
	"shortener-link/internal/models"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	var request models.LinkRequest
	if err := c.BodyParser(&request); err != nil {
		jsonErr := c.JSON(models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return errors.Wrap(err, jsonErr.Error())
	}
	_, err := h.service.Create(c.Context(), request.Link)
	if err != nil {
		jsonErr := c.JSON(models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return errors.Wrap(err, jsonErr.Error())
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) GetAll(c *fiber.Ctx) error {

	links, err := h.service.GetAll(c.Context())
	if err != nil {
		jsonErr := c.JSON(models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return errors.Wrap(err, jsonErr.Error())
	}

	c.Status(http.StatusOK).JSON(links)
	return nil
}

func (h *Handler) GetByShortener(c *fiber.Ctx) error {

	link := c.Query("link")

	fullLink, err := h.service.GetByShortenerLink(c.Context(), link)
	if err != nil {
		if err == models.ErrLinkNotFound {
			jsonErr := c.JSON(models.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return errors.Wrap(models.ErrLinkNotFound, jsonErr.Error())
		}
		jsonErr := c.JSON(models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return errors.Wrap(err, jsonErr.Error())
	}

	c.Status(http.StatusOK).JSON(fullLink)
	return nil

}

func (h *Handler) Delete(c *fiber.Ctx) error {

	link := c.Query("link")

	if err := h.service.DeleteShortenerLink(c.Context(), link); err != nil {
		jsonErr := c.JSON(models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return errors.Wrap(err, jsonErr.Error())
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) GetStats(c *fiber.Ctx) error {

	link := c.Query("link")

	stat, lastVisitTime, err := h.service.GetStatsByShortenerLink(c.Context(), link)
	if err != nil {
		if err == models.ErrLinkNotFound {
			jsonErr := c.JSON(models.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return errors.Wrap(models.ErrLinkNotFound, jsonErr.Error())
		}
		jsonErr := c.JSON(models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return errors.Wrap(err, jsonErr.Error())
	}

	c.Status(http.StatusOK).JSON(models.GetStatsResponse{
		StatsCount:    stat,
		LastVisitTime: lastVisitTime,
	})
	return nil
}
