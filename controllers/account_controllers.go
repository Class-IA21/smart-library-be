package controllers

import (
	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type AccountControllerInterface interface {
	RegisterAccount(c *fiber.Ctx) error
	LoginAccount(c *fiber.Ctx) error
	UpdateAccount(c *fiber.Ctx) error
	DeleteAccount(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	GetAccountByID(c *fiber.Ctx) error
}

type AccountController struct {
	*services.AccountServices
}

func NewAccountController(as *services.AccountServices) *AccountController {
	return &AccountController{
		AccountServices: as,
	}
}

func (c *AccountController) RegisterAccount(ctx *fiber.Ctx) error {
	var account entity.AccountRequest
	if err := ctx.BodyParser(&account); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid request payload")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	if errorResponse := helper.ValidateStruct(&account); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.AccountServices.RegisterAccount(ctx.Context(), &account)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusCreated, "Account successfully registered.")
	return ctx.JSON(response)
}

func (c *AccountController) LoginAccount(ctx *fiber.Ctx) error {
	var login entity.Login
	if err := ctx.BodyParser(&login); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid request payload")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	if errorResponse := helper.ValidateStruct(&login); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	userDetail, errorResponse := c.AccountServices.LoginAccount(ctx.Context(), &login)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusCreated, "Login Successfully.", userDetail)
	return ctx.JSON(response)
}

func (c *AccountController) UpdateAccount(ctx *fiber.Ctx) error {
	accountId, err := strconv.Atoi(ctx.Params("accountId"))
	if err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "invalid format accountId")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	var account entity.AccountRequest
	if err := ctx.BodyParser(&account); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid request payload")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}
	account.ID = accountId

	if errorResponse := helper.ValidateStruct(&account); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.AccountServices.UpdateAccount(ctx.Context(), &account)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusCreated, "Account successfully updated")
	return ctx.JSON(response)
}

func (c *AccountController) DeleteAccount(ctx *fiber.Ctx) error {
	accountId, err := strconv.Atoi(ctx.Params("accountId"))
	if err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "invalid format accountId")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.AccountServices.DeleteAccount(ctx.Context(), accountId)
	if errorResponse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusCreated, "Account successfully deleted")
	return ctx.JSON(response)
}

func (c *AccountController) ChangePassword(ctx *fiber.Ctx) error {
	accountId, err := strconv.Atoi(ctx.Params("accountId"))
	if err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "invalid format accountId")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	var account entity.AccountChangePasswordRequest
	if err := ctx.BodyParser(&account); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid request payload")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.AccountServices.ChangePassword(ctx.Context(), &account, accountId)
	if errorResponse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusCreated, "Account password successfully updated")
	return ctx.JSON(response)
}

func (c *AccountController) GetAccountByID(ctx *fiber.Ctx) error {
	accountId, err := strconv.Atoi(ctx.Params("accountId"))
	if err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "invalid format accountId")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	account, errorResponse := c.AccountServices.GetAccountByID(ctx.Context(), accountId)
	if errorResponse != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusCreated, "OK", account)
	return ctx.JSON(response)
}
