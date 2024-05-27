package main

import (
	"database/sql"
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

type App struct {
	BookController        *controllers.BookController
	StudentController     *controllers.StudentController
	CardController        *controllers.CardController
	BorrowController      *controllers.BorrowController
	BookCardController    *controllers.BookCardController
	StudentCardController *controllers.StudentCardController
	AccountController     *controllers.AccountController
}

func NewApp(database *sql.DB) *App {
	bookRepository := repository.NewBookRepository()
	bookService := services.NewBookServices(bookRepository, database)
	bookController := controllers.NewBookController(bookService)

	studentRepository := repository.NewStudentRepository()
	studentService := services.NewStudentServices(database, studentRepository)
	studentController := controllers.NewStudentController(studentService)

	cardRepository := repository.NewCardRepository()
	cardService := services.NewCardServices(database, cardRepository, studentService, bookService)
	cardController := controllers.NewCardController(cardService)

	borrowRepository := repository.NewBorrowRepository()
	borrowService := services.NewBorrowServices(database, borrowRepository, studentService, bookService)
	borrowController := controllers.NewBorrowController(borrowService)

	bookCardService := services.NewBookCardServices(database, bookService, cardService)
	bookCardController := controllers.NewBookCardController(bookCardService)

	studentCardService := services.NewStudentCardServices(database, cardService, studentService)
	studentCardController := controllers.NewStudentCardController(studentCardService)

	accountRepository := repository.NewAccountsRepository()
	accountService := services.NewAccountServices(database, accountRepository, studentService)
	accountController := controllers.NewAccountController(accountService)

	return &App{
		BookController:        bookController,
		StudentController:     studentController,
		CardController:        cardController,
		BorrowController:      borrowController,
		BookCardController:    bookCardController,
		StudentCardController: studentCardController,
		AccountController:     accountController,
	}
}

func main() {
	database, _ := db.Connection()
	controller := NewApp(database)
	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "PID ${pid} | [${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	router.RegisterBookRoutes("books", app, controller.BookController, controller.BookCardController)
	router.RegisterCardRoutes("cards", app, controller.CardController)
	router.RegisterStudentRoutes("students", app, controller.StudentController, controller.StudentCardController)
	router.RegisterBorrowRoutes("borrows", app, controller.BorrowController)
	router.RegisterAccountRoutes("accounts", app, controller.AccountController)
	router.RegisterAuthRoutes("auth", app, controller.AccountController)
	log.Fatal(app.Listen(":3000"))
}
