package model

import (
	"go-rest-todolist/database"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DBConn
	todo := new(Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err, "message": "Check your input", "data": err})
	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": err, "message": "Could not create todo", "data": err})
	}

	return c.JSON(&todo)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var todo Todo
	err := db.Find(&todo, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": err, "message": "Could not find todo", "data": err})
	}

	return c.JSON(&todo)
}
