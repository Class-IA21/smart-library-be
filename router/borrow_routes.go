package router

import (
	"fmt"

	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBorrowRoutes(path string, app *fiber.App, controller *controllers.BorrowController) {
	app.Get(fmt.Sprintf("/%s/student/:studentId", path), controller.GetTransactionsByStudentID)
	app.Get(fmt.Sprintf("/%s", path), controller.GetBorrows)
	app.Get(fmt.Sprintf("/%s/book/:bookId", path), controller.GetBorrowsByBookID)
	app.Post(fmt.Sprintf("/%s", path), controller.InsertBorrow)
	app.Get(fmt.Sprintf("/%s/:transactionId", path), controller.GetBorrowByTransactionID)
	app.Put(fmt.Sprintf("/%s/:transactionId", path), controller.UpdateBorrow)
}
