package controllers

import (
	"context"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type CardControllerInterface interface {
	GetCardByID(c *fiber.Ctx) error
	GetCardTypeByUID(ctx *fiber.Ctx) error
}

type CardController struct {
	service services.CardServiceInterface
}

func NewCardController(service *services.CardServices) *CardController {
	return &CardController{
		service: service,
	}
}

func (c *CardController) GetCardByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	book, err := c.service.GetCardByID(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := entity.ResponseWeb{
		Error:   false,
		Message: "OK",
		Data:    book,
	}

	return ctx.JSON(response)
}

func (c *CardController) GetCardTypeByUID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("uid"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	book, cardType, err := c.service.GetCardTypeByUID(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := entity.ResponseWeb{
		Error:   false,
		Message: "OK",
		Data: map[string]interface{}{
			"type":  cardType,
			"value": book,
		},
	}

	return ctx.JSON(response)
}
