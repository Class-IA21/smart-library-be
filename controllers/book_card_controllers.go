package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type BookCardControllerInterface interface {
	InsertBook(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
}

type BookCardController struct {
	*services.BookCardServices
}

func NewBookCardController(sc *services.BookCardServices) *BookCardController {
	return &BookCardController{
		BookCardServices: sc,
	}
}

func (c *BookCardController) InsertBook(ctx *fiber.Ctx) error {
	var book entity.Book
	if err := ctx.BodyParser(&book); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	if errorResponse := helper.ValidateStruct(&book); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	id, errorResponse := c.BookCardServices.InsertBook(ctx.Context(), &book)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "Data Book successfully created.", map[string]any{
		"id": id,
	})
	return ctx.Status(http.StatusCreated).JSON(response)
}

func (c *BookCardController) UpdateBook(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(helper.ErrorResponse(http.StatusBadRequest, "invalid id Book"))
	}

	var Book entity.Book
	if err := ctx.BodyParser(&Book); err != nil {
		fmt.Println(err)
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid payload request")
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}
	Book.ID = id

	if errorResponse := helper.ValidateStruct(&Book); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.BookCardServices.UpdateBook(ctx.Context(), &Book)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)

	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Data Book successfully updated.")
	return ctx.JSON(response)
}
