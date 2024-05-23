package controllers

import (
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type CardControllerInterface interface {
	GetCards(c *fiber.Ctx) error
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

func (c *CardController) GetCards(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	cards, errorResponse := c.service.GetCards(ctx.Context(), page, pageSize)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := entity.ResponseWebWithData{
		Error:   false,
		Message: "OK",
		Data:    cards,
	}

	return ctx.JSON(response)

}

func (c *CardController) GetCardByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(helper.ErrorResponse(http.StatusBadRequest, "invalid card id"))
	}

	card, errorResponse := c.service.GetCardByID(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := entity.ResponseWebWithData{
		Error:   false,
		Message: "OK",
		Data:    card,
	}

	return ctx.JSON(response)
}

func (c *CardController) GetCardTypeByUID(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	card, cardType, errorResponse := c.service.GetCardTypeByUID(ctx.Context(), uid)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := entity.ResponseWebWithData{
		Error:   false,
		Message: "OK",
		Data: map[string]interface{}{
			"card_type": cardType,
			"value":     card,
		},
	}

	return ctx.JSON(response)
}
