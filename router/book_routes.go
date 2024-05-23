package router

import (
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoutes(app *fiber.App, controller *controllers.BookController) {
	app.Get("/books/:id", controller.GetBookByID)
	app.Delete("/books/:id", controller.DeleteBookByID)
	app.Put("/books/:id", controller.UpdateBookByID)
	app.Get("/books/card/:cardId", controller.GetBookByCardID)
	app.Get("/books", controller.GetBooks)
}
