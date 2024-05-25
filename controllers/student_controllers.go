package controllers

import (
	"net/http"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/entity"
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
	service *services.StudentServices
}

func NewStudentController(service *services.StudentServices) *StudentController {
	return &StudentController{
		service: service,
	}
}

func (c *StudentController) GetStudents(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	students, errorResponse := c.service.GetStudents(ctx.Context(), page, pageSize)
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

	student, errorResponse := c.service.GetStudentByID(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithData(http.StatusOK, "OK", student)
	return ctx.JSON(response)
}

func (c *StudentController) InsertStudent(ctx *fiber.Ctx) error {
	var student entity.Student
	if err := ctx.BodyParser(&student); err != nil {
		errorResponse := helper.ErrorResponse(http.StatusBadRequest, err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	if errorResponse := helper.ValidateStruct(&student); errorResponse != nil {
		return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	errorResponse := c.service.InsertStudent(ctx.Context(), &student)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Data student successfully created.")
	return ctx.Status(http.StatusCreated).JSON(response)
}

func (c *StudentController) UpdateStudent(ctx *fiber.Ctx) error {
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

	errorResponse := c.service.UpdateStudent(ctx.Context(), &student)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)

	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Data student successfully updated.")
	return ctx.JSON(response)
}

func (c *StudentController) DeleteStudent(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(helper.ErrorResponse(http.StatusBadRequest, "invalid id student"))
	}

	errorResponse := c.service.DeleteStudent(ctx.Context(), id)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	response := helper.SuccessResponseWithoutData(http.StatusOK, "Data student successfully deleted.")
	return ctx.Status(http.StatusOK).JSON(response)
}
