package router

import (
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterCardRoutes(app *fiber.App, controller *controllers.CardController) {
	app.Get("/cards", controller.GetCards)
	app.Get("/cards/:id", controller.GetCardByID)
	app.Get("/check_card", controller.GetCardTypeByUID)
	app.Delete("/cards/:id", controller.DeleteCard)
	app.Post("/cards", controller.InsertCard)
	app.Put("/cards/:id", controller.UpdateCard)
}
