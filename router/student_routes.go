package router

import (
	"fmt"
	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterStudentRoutes(path string, app *fiber.App, controller *controllers.StudentController, controllerStudentCard *controllers.StudentCardController) {
	app.Get(fmt.Sprintf("/%s", path), controller.GetStudents)
	app.Get(fmt.Sprintf("/%s/:id", path), controller.GetStudentByID)
	app.Delete(fmt.Sprintf("/%s/:id", path), controller.DeleteStudent)
	app.Put(fmt.Sprintf("/%s/:id", path), controllerStudentCard.UpdateStudent)
	app.Post(fmt.Sprintf("/%s", path), controllerStudentCard.InsertStudent)
}
