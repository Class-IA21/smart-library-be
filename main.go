package main

import (
	"log"

	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/dimassfeb-09/smart-library-be/db"
	"github.com/dimassfeb-09/smart-library-be/repository"
	"github.com/dimassfeb-09/smart-library-be/router"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	database, _ := db.Connection()

	bookRepository := repository.NewBookRepository()
	bookService := services.NewBookService(bookRepository, database)
	bookController := controllers.NewBookController(bookService)

	studentRepository := repository.NewStudentRepository()

	cardRepository := repository.NewCardRepository()
	cardService := services.NewCardServices(database, bookRepository, cardRepository, studentRepository)
	cardController := controllers.NewCardController(cardService)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "PID ${pid} | [${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	router.RegisterBookRoutes(app, bookController)
	router.RegisterCardRoutes(app, cardController)
	log.Fatal(app.Listen(":3000"))
}
