package controllers

import (
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type StudentCardControllerInterface interface {
	InsertStudent(c *fiber.Ctx) error
	UpdateStudent(c *fiber.Ctx) error
}

type StudentCardController struct {
	*services.StudentCardServices
}

func NewStudentCardController(sc *services.StudentCardServices) *StudentCardController {
	return &StudentCardController{
		StudentCardServices: sc,
	}
}

func (c *StudentCardController) InsertStudent(ctx *fiber.Ctx) error {
	var student entity.Student
	if err := ctx.BodyParser(&student); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	if errorResponse := helper.ValidateStruct(&student); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.StudentCardServices.InsertStudent(ctx.Context(), &student)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Data student successfully created.")
	return ctx.Status(http.StatusCreated).JSON(response)
}

func (c *StudentCardController) UpdateStudent(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(helper.ErrorResponse(http.StatusBadRequest, "invalid id student"))
	}

	var student entity.Student
	if err := ctx.BodyParser(&student); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "Invalid payload request")
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}
	student.ID = id

	if errorResponse := helper.ValidateStruct(&student); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.StudentCardServices.UpdateStudent(ctx.Context(), &student)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)

	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Data student successfully updated.")
	return ctx.JSON(response)
}
