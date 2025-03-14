package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Создать задачу
func CreateTaskHandler(c *fiber.Ctx) error {
	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if task.Status == "" {
		task.Status = "new" // Значение по умолчанию
	}

	if err := CreateTask(task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create task"})
	}
	return c.Status(201).JSON(task)
}

// Получить все задачи
func GetTasksHandler(c *fiber.Ctx) error {
	tasks, err := GetTasks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}
	return c.JSON(tasks)
}

// Обновить задачу
func UpdateTaskHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	task := new(Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := UpdateTask(id, task); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update task"})
	}
	return c.JSON(task)
}

// Удалить задачу
func DeleteTaskHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	if err := DeleteTask(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete task"})
	}
	return c.SendStatus(204)
}
