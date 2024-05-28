package router

import (
	"fmt"
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAccountRoutes(path string, app *fiber.App, controller *controllers.AccountController, nc *controllers.NotificationController) {
	app.Get(fmt.Sprintf("/%s/:accountId", path), controller.GetAccountByID)
	app.Get(fmt.Sprintf("/%s/changePassword", path), controller.ChangePassword)
	app.Delete(fmt.Sprintf("/%s/:accountId", path), controller.DeleteAccount)
	app.Put(fmt.Sprintf("/%s/:accountId", path), controller.UpdateAccount)
	app.Get(fmt.Sprintf("/%s/:accountId/notifications", path), nc.GetNotificationByAccountID)
}
