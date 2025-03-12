package routes

import (
	"evermos-backend/controllers"
	"evermos-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddressRoutes(app *fiber.App) {
	addressRoutes := app.Group("/addresses")

	addressRoutes.Post("/", middleware.AuthMiddleware, controllers.CreateAddress)
	addressRoutes.Get("/", middleware.AuthMiddleware, controllers.GetAddressesByUser)
	addressRoutes.Put("/:id", middleware.AuthMiddleware, controllers.UpdateAddress)
	addressRoutes.Delete("/:id", middleware.AuthMiddleware, controllers.DeleteAddress)
}
