package router

import (
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoutes(app *fiber.App, controller *controllers.BookController) {
	app.Get("/books", controller.GetBookByID)
	app.Get("/books", controller.GetAllBooks)
	app.Delete("/books", controller.DeleteBookByID)
	app.Put("/books", controller.UpdateBookByID)
	app.Get("/books", controller.GetBookByCardID)
}
