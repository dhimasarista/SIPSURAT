package routes

import (
	"sipsurat/controllers"

	"github.com/gofiber/fiber/v2"
)

func IndexRoutes(app *fiber.App) {
	app.Get("/", controllers.IndexRender)
}
