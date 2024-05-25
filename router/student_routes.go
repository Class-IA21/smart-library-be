package router

import (
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterStudentRoutes(app *fiber.App, controller *controllers.StudentController, controllerStudentCard *controllers.StudentCardController) {
	app.Get("/students", controller.GetStudents)
	app.Get("/students/:id", controller.GetStudentByID)
	app.Delete("/students/:id", controller.DeleteStudent)
	app.Put("/students/:id", controllerStudentCard.UpdateStudent)
	app.Post("/students", controllerStudentCard.InsertStudent)
}
