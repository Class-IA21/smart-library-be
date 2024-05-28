package controllers

import (
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type NotificationControllerInterface interface {
	GetNotificationByAccountID(ctx *fiber.Ctx) error
}

type NotificationController struct {
	*services.NotificationServices
}

func NewNotificationController(ns *services.NotificationServices) *NotificationController {
	return &NotificationController{
		NotificationServices: ns,
	}
}

func (nc *NotificationController) GetNotificationByAccountID(ctx *fiber.Ctx) error {
	accountId, err := strconv.Atoi(ctx.Params("accountId"))
	if err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid format account_id")
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	response, errorResponse := nc.NotificationServices.GetNotificationByAccountID(ctx.Context(), accountId)
	if errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	return ctx.JSON(helper.SuccessResponseWithData(http.StatusOK, "OK", response))
}
