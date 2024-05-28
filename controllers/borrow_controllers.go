package controllers

import (
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type BorrowControllerInterface interface {
	GetTransactionsByStudentID(c *fiber.Ctx) error
	GetBorrowByTransactionID(c *fiber.Ctx) error
	GetBorrowsByBookID(c *fiber.Ctx) error
	GetBorrows(c *fiber.Ctx) error
	InsertBorrow(c *fiber.Ctx) error
	UpdateBorrow(c *fiber.Ctx) error
}

type BorrowController struct {
	service *services.BorrowServices
}

func NewBorrowController(service *services.BorrowServices) *BorrowController {
	return &BorrowController{
		service: service,
	}
}

func (c *BorrowController) GetTransactionsByStudentID(ctx *fiber.Ctx) error {
	studentId, err := strconv.Atoi(ctx.Params("studentId"))
	if err != nil || studentId <= 0 {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "ID mahasiswa tidak valid")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	student, errorResponse := c.service.GetBorrowsByStudentID(ctx.Context(), studentId)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", student)
	return ctx.JSON(response)
}

func (c *BorrowController) GetBorrowsByBookID(ctx *fiber.Ctx) error {
	bookId, err := strconv.Atoi(ctx.Params("bookId"))
	if err != nil || bookId <= 0 {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "ID Book tidak valid")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	student, errorResponse := c.service.GetBorrowByBookID(ctx.Context(), bookId)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", student)
	return ctx.JSON(response)
}

func (c *BorrowController) GetBorrows(ctx *fiber.Ctx) error {
	student, errorResponse := c.service.GetBorrows(ctx.Context())
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", student)
	return ctx.JSON(response)
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

func (c *BorrowController) UpdateBorrow(ctx *fiber.Ctx) error {

	transactionID := ctx.Params("transactionId")

	var borrow entity.BorrowUpdate
	if err := ctx.BodyParser(&borrow); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid request payload")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}
	borrow.TransactionID = transactionID

	if errorResponse := helper.ValidateStruct(&borrow); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.service.UpdateBorrow(ctx.Context(), &borrow)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusCreated, "Borrow successfully updated")
	return ctx.JSON(response)
}
