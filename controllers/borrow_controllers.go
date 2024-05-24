package controllers

import (
	"net/http"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type BorrowControllerInterface interface {
	GetBorrowByTransactionID(c *fiber.Ctx) error
	InsertBorrow(c *fiber.Ctx) error
}

type BorrowController struct {
	service *services.BorrowServices
}

func NewBorrowController(service *services.BorrowServices) *BorrowController {
	return &BorrowController{
		service: service,
	}
}

func (c *BorrowController) GetBorrowByTransactionID(ctx *fiber.Ctx) error {
	transactionID := ctx.Params("transactionId")

	borrow, errorResponse := c.service.GetBorrowByTransactionID(ctx.Context(), transactionID)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	if borrow == nil {
		response := helper.SuccessResponseWithoutData(http.StatusNotFound, "Transaction not found")
		return ctx.JSON(response)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", borrow)
	return ctx.JSON(response)
}

func (c *BorrowController) InsertBorrow(ctx *fiber.Ctx) error {
	var borrow entity.Borrow
	if err := ctx.BodyParser(&borrow); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid request payload")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	if errorResponse := helper.ValidateStruct(&borrow); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.service.InsertBorrow(ctx.Context(), &borrow)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusCreated, "Borrow successfully inserted")
	return ctx.Status(http.StatusCreated).JSON(response)
}
