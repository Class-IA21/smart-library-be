package router

import (
	"fmt"
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookRoutes(path string, app *fiber.App, bc *controllers.BookController, bcc *controllers.BookCardController) {
	app.Get(fmt.Sprintf("/%s", path), bc.GetBooks)
	app.Get(fmt.Sprintf("/%s/:id", path), bc.GetBookByID)
	app.Delete(fmt.Sprintf("/%s/:id", path), bc.DeleteBookByID)
	app.Put(fmt.Sprintf("/%s/:id", path), bcc.UpdateBook)
	app.Post(fmt.Sprintf("/%s", path), bcc.InsertBook)
}
