package main

import (
	"log"

	"github.com/dimassfeb-09/smart-library-be/controllers"
	"github.com/dimassfeb-09/smart-library-be/db"
	"github.com/dimassfeb-09/smart-library-be/repository"
	"github.com/dimassfeb-09/smart-library-be/router"
	"github.com/dimassfeb-09/smart-library-be/services"
	"github.com/gofiber/fiber/v2"
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
	router.RegisterBookRoutes(app, bookController)
	router.RegisterCardRoutes(app, cardController)
	log.Fatal(app.Listen(":3000"))
}
