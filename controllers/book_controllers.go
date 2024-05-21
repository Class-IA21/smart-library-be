package controllers

import (
	"context"
	"strconv"

	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
)

type BookControllerInterface interface {
	GetBookByID(c *fiber.Ctx) error
	GetAllBooks(c *fiber.Ctx) error
	DeleteBookByID(c *fiber.Ctx) error
	UpdateBookByID(c *fiber.Ctx) error
	GetBookByCardID(c *fiber.Ctx) error
}

type BookController struct {
	service services.BookServiceInterface
}

func NewBookController(service *services.BookService) *BookController {
	return &BookController{
		service: service,
	}
}

func (c *BookController) GetBookByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	book, err := c.service.GetBookByID(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := entity.ResponseWeb{
		Error:   false,
		Message: "OK",
		Data:    book,
	}

	return ctx.JSON(response)
}

func (c *BookController) GetAllBooks(ctx *fiber.Ctx) error {
	books, err := c.service.GetAllBooks(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := entity.ResponseWeb{
		Error:   false,
		Message: "OK",
		Data:    books,
	}

	return ctx.JSON(response)
}

func (c *BookController) DeleteBookByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil || id <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	err = c.service.DeleteBookByID(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *BookController) UpdateBookByID(ctx *fiber.Ctx) error {
	var book entity.Book
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	err := c.service.UpdateBookByID(context.Background(), &book)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *BookController) GetBookByCardID(ctx *fiber.Ctx) error {
	cardID, err := strconv.Atoi(ctx.Query("card_id"))

	if err != nil || cardID <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	book, err := c.service.GetBookByCardID(context.Background(), cardID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := entity.ResponseWeb{
		Error:   false,
		Message: "OK",
		Data:    book,
	}

	return ctx.JSON(response)
}
