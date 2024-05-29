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
	InsertCard(ctx *fiber.Ctx) error
	UpdateCard(ctx *fiber.Ctx) error
	DeleteCard(ctx *fiber.Ctx) error
	InsertContainerCard(ctx *fiber.Ctx) error
	GetOnceContainerCardByUID(ctx *fiber.Ctx) error
}

type CardController struct {
	service *services.CardServices
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

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", cards)
	return ctx.JSON(response)

}

func (c *CardController) GetCardByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		response := helper.ErrorResponse(http.StatusBadRequest, "invalid card id")
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	card, errorResponse := c.service.GetCardByID(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", card)
	return ctx.JSON(response)
}

func (c *CardController) GetCardTypeByUID(ctx *fiber.Ctx) error {
	uid := ctx.Query("uid")

	id, cardType, errorResponse := c.service.GetCardTypeByUID(ctx.Context(), uid)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", fiber.Map{
		"id":        id,
		"card_type": cardType,
	})
	return ctx.JSON(response)
}

func (c *CardController) InsertCard(ctx *fiber.Ctx) error {
	var card entity.Card

	if err := ctx.BodyParser(&card); err != nil {
		response := helper.ErrorResponse(fiber.StatusBadRequest, "Invalid request")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if errorResponse := helper.ValidateStruct(&card); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	if errorResponse := c.service.InsertCard(ctx.Context(), &card); errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusCreated, "Card created successfully")
	return ctx.Status(http.StatusCreated).JSON(response)
}

func (c *CardController) UpdateCard(ctx *fiber.Ctx) error {
	var card entity.Card

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		response := helper.ErrorResponse(fiber.StatusBadRequest, "invalid card id")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	card.ID = id

	if err := ctx.BodyParser(&card); err != nil {
		response := helper.ErrorResponse(fiber.StatusBadRequest, "request invalid")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if errorResponse := helper.ValidateStruct(&card); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.service.UpdateCard(ctx.Context(), id, &card)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Card updated successfully")
	return ctx.JSON(response)
}

func (c *CardController) DeleteCard(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		response := helper.ErrorResponse(fiber.StatusBadRequest, "invalid card id")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	errorResponse := c.service.DeleteCard(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Card deleted successfully")
	return ctx.JSON(response)
}

func (c *CardController) InsertContainerCard(ctx *fiber.Ctx) error {
	var containerCard entity.ContainerCard
	if err := ctx.BodyParser(&containerCard); err != nil {
		response := helper.ErrorResponse(fiber.StatusBadRequest, "Invalid request")
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	if errorResponse := helper.ValidateStruct(&containerCard); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	if errorResponse := c.service.InsertContainerCard(ctx.Context(), containerCard.UID); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Container card inserted successfully")
	return ctx.JSON(response)
}

func (c *CardController) GetOnceContainerCardByUID(ctx *fiber.Ctx) error {

	uid, errorResponse := c.service.GetOnceContainerCard(ctx.Context())
	if errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "Container card retrived successfully", entity.ContainerCard{
		UID: uid,
	})
	return ctx.JSON(response)
}
