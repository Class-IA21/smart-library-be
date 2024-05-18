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
	database, err := db.Connection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	bookRepository := repository.NewBookRepositoryImpl()
	bookService := services.NewBookService(bookRepository, database)
	bookController := controllers.NewBookController(bookService)

	app := fiber.New()
	router.RegisterBookRoutes(app, bookController)
	log.Fatal(app.Listen(":3000"))
}
