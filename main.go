package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Подключение к базе данных
	if err := ConnectDB(); err != nil {
		panic("Failed to connect to database")
	}

	app := fiber.New()

	// Роуты
	app.Post("/tasks", CreateTaskHandler)
	app.Get("/tasks", GetTasksHandler)
	app.Put("/tasks/:id", UpdateTaskHandler)
	app.Delete("/tasks/:id", DeleteTaskHandler)

	app.Listen(":3000")
}
