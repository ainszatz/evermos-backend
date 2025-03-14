package routes

import (
	"evermos-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductLogRoutes(app *fiber.App) {
	productLog := app.Group("/product-logs")
	productLog.Get("/", controllers.GetProductLogs)
}
