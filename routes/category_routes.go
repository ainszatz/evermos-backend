package routes

import (
	"evermos-backend/controllers"
	"evermos-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	category := app.Group("/categories")

	// ✅ Gunakan middleware `AuthMiddleware` sebelum `AdminOnly`
	category.Post("/", middleware.AuthMiddleware, middleware.AdminOnly(), controllers.CreateCategory)
	category.Get("/", controllers.GetCategories)
	category.Put("/:id", middleware.AuthMiddleware, middleware.AdminOnly(), controllers.UpdateCategory)
	category.Delete("/:id", middleware.AuthMiddleware, middleware.AdminOnly(), controllers.DeleteCategory)
}
