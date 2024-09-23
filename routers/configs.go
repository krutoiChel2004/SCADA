package routers

import (
	"test/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterConfigRoutes(app *fiber.App) {
	api := app.Group("/config")

	api.Post("/", handlers.CreateConv)
}
