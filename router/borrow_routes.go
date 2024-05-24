package router

import (
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBorrowRoutes(app *fiber.App, controller *controllers.BorrowController) {
	app.Get("/borrow/student/:studentId", controller.GetTransactionsByStudentID)
	app.Post("/borrow", controller.InsertBorrow)
	app.Get("/borrow/:transactionId", controller.GetBorrowByTransactionID)
}
