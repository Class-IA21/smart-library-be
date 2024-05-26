package router

import (
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoutes(app *fiber.App, bc *controllers.BookController, bcc *controllers.BookCardController) {
	app.Get("/books", bc.GetBooks)
	app.Get("/books/:id", bc.GetBookByID)
	app.Delete("/books/:id", bc.DeleteBookByID)
	app.Put("/books/:id", bcc.UpdateBook)
	app.Post("/books", bcc.InsertBook)
}
