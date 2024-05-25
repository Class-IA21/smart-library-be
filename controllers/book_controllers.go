package controllers

import (
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type BookControllerInterface interface {
	GetBookByID(c *fiber.Ctx) error
	GetBooks(c *fiber.Ctx) error
	DeleteBookByID(c *fiber.Ctx) error
	UpdateBookByID(c *fiber.Ctx) error
	InsertBook(c *fiber.Ctx) error
}

type BookController struct {
	service *services.BookServices
}

func NewBookController(service *services.BookServices) *BookController {
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

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", book)
	return ctx.JSON(response)
}

func (c *BookController) GetBooks(ctx *fiber.Ctx) error {

	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	books, errorResponse := c.service.GetBooks(ctx.Context(), page, pageSize)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", books)
	return ctx.JSON(response)
}

func (c *BookController) DeleteBookByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid book id")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.service.DeleteBookByID(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Successfully delete book.")
	return ctx.Status(http.StatusOK).JSON(response)
}
