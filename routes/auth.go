package routes

import (
	"sipsurat/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Get("/login", controllers.LoginRender)
}
