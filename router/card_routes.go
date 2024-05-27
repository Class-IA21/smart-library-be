package router

import (
	"fmt"
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterCardRoutes(path string, app *fiber.App, controller *controllers.CardController) {
	app.Get(fmt.Sprintf("/%s/check_card", path), controller.GetCardTypeByUID)
	app.Get(fmt.Sprintf("/%s", path), controller.GetCards)
	app.Get(fmt.Sprintf("/%s/:id", path), controller.GetCardByID)
	app.Delete(fmt.Sprintf("/%s/:id", path), controller.DeleteCard)
	app.Post(fmt.Sprintf("/%s", path), controller.InsertCard)
	app.Put(fmt.Sprintf("/%s/:id", path), controller.UpdateCard)
}
