package router

import (
	"fmt"
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(path string, app *fiber.App, controller *controllers.AccountController) {
	app.Post(fmt.Sprintf("/%s/register", path), controller.RegisterAccount)
	app.Post(fmt.Sprintf("/%s/login", path), controller.LoginAccount)
}
