package routes

import (
	"evermos-backend/controllers"
	"evermos-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func StoreRoutes(app *fiber.App) {
	store := app.Group("/stores", middleware.AuthMiddleware)

	store.Get("/me", controllers.GetStore)
	store.Put("/me", controllers.UpdateStore)
}
