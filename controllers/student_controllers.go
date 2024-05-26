package controllers

import (
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/helper"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type StudentControllerInterface interface {
	GetStudents(c *fiber.Ctx) error
	GetStudentByID(c *fiber.Ctx) error
	InsertStudent(c *fiber.Ctx) error
	UpdateStudent(c *fiber.Ctx) error
	DeleteStudent(c *fiber.Ctx) error
}

type StudentController struct {
	*services.StudentServices
}

func NewStudentController(ss *services.StudentServices) *StudentController {
	return &StudentController{
		StudentServices: ss,
	}
}

func (c *StudentController) GetStudents(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	students, errorResponse := c.StudentServices.GetStudents(ctx.Context(), page, pageSize)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", students)
	return ctx.JSON(response)
}

func (c *StudentController) GetStudentByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, "ID mahasiswa tidak valid")
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	student, errorResponse := c.StudentServices.GetStudentByID(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", student)
	return ctx.JSON(response)
}

func (c *StudentController) DeleteStudent(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(helper.ErrorResponse(http.StatusBadRequest, "invalid id student"))
	}

	errorResponse := c.StudentServices.DeleteStudent(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Data student successfully deleted.")
	return ctx.Status(http.StatusOK).JSON(response)
}
