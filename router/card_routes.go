package router

import (
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterCardRoutes(app *fiber.App, controller *controllers.CardController) {
	app.Get("/cards", controller.GetCardTypeByUID)
	//app.Get("/cards", controller.GetCardByID)
}
