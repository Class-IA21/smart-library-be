package controllers

import (
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type BookControllerInterface interface {
	GetBookByID(c *fiber.Ctx) error
	GetBooks(c *fiber.Ctx) error
	DeleteBookByID(c *fiber.Ctx) error
	UpdateBookByID(c *fiber.Ctx) error
	GetBookByCardID(c *fiber.Ctx) error
}

type BookController struct {
	service *services.BookService
}

func NewBookController(service *services.BookService) *BookController {
	return &BookController{
		service: service,
	}
}

func (c *BookController) GetBookByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(http.StatusBadRequest, "Invalid book id"))
	}

	book, errorResponse := c.service.GetBookByID(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := entity.ResponseWebWithData{
		Error:   false,
		Message: "OK",
		Data:    book,
	}

	return ctx.JSON(response)
}

func (c *BookController) GetBooks(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	books, errorResponse := c.service.GetBooks(ctx.Context(), page, pageSize)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := entity.ResponseWebWithData{
		Error:   false,
		Message: "OK",
		Data:    books,
	}

	return ctx.JSON(response)
}

func (c *BookController) DeleteBookByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(helper.ErrorResponse(http.StatusBadRequest, "Invalid book id"))
	}

	errorResponse := c.service.DeleteBookByID(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := entity.ResponseWebWithoutData{
		Error:   false,
		Message: "Successfully delete book.",
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (c *BookController) UpdateBookByID(ctx *fiber.Ctx) error {
	var book entity.Book
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(helper.ErrorResponse(http.StatusBadRequest, "Invalid request payload"))
	}

	errorResponse := c.service.UpdateBookByID(ctx.Context(), &book)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := entity.ResponseWebWithoutData{
		Error:   false,
		Message: "Successfully update book.",
	}
	return ctx.Status(http.StatusOK).JSON(response)
}

func (c *BookController) GetBookByCardID(ctx *fiber.Ctx) error {
	cardID, err := strconv.Atoi(ctx.Params("cardId"))
	if err != nil || cardID <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(http.StatusBadRequest, "type card id only integer"))
	}

	book, errorResponse := c.service.GetBookByCardID(ctx.Context(), cardID)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := entity.ResponseWebWithData{
		Error:   false,
		Message: "OK",
		Data:    book,
	}

	return ctx.JSON(response)
}
