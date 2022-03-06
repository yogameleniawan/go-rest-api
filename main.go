package main

import (
	"fmt"
	"go-rest-todolist/database"
	model "go-rest-todolist/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func initDatabase() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/go_todo?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&model.Todo{})
	fmt.Println("Migrated DB")
}

func setupRoute(app *fiber.App) {
	app.Get("/todos", model.GetTodos)
	app.Get("/todos/:id", model.GetTodoById)
	app.Post("/todos", model.CreateTodo)
	app.Put("/todos/:id", model.UpdateTodo)
	app.Delete("/todos/:id", model.DeleteTodo)
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", helloWorld)
	setupRoute(app)
	app.Listen(":8000")
}
