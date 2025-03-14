package routes

import (
	"evermos-backend/controllers"
	"evermos-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(app *fiber.App) {
	transaction := app.Group("/transactions")

	transaction.Post("/", middleware.AuthMiddleware, controllers.CreateTransaction)
	transaction.Get("/", middleware.AuthMiddleware, controllers.GetTransactions)
}
