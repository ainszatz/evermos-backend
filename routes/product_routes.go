package routes

import (
	"evermos-backend/controllers"
	"evermos-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	product := app.Group("/products")

	product.Get("/", controllers.GetProducts)
	product.Post("/", middleware.AuthMiddleware, controllers.CreateProduct)
	product.Put("/:id", middleware.AuthMiddleware, controllers.UpdateProduct)
	product.Delete("/:id", middleware.AuthMiddleware, controllers.DeleteProduct)
}
