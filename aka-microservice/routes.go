package routes

import (
	"github.com/gofiber/fiber/v2"
	"microservice/internal/handlers"
)

func InitHandlers(app *fiber.App) {
	app.Get("/get-items", handlers.GetItems)
}
