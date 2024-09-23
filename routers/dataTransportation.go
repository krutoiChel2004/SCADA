package routers

import (
	"test/handlers"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func RegisterDataTransportationRoutes(app *fiber.App) {
	api := app.Group("/dataTransportation")

	api.Get("/C2C/:name", websocket.New(handlers.C2C))
}
