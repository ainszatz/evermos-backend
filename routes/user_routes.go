package routes

import (
	"evermos-backend/controllers"
	"evermos-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/users", middleware.AuthMiddleware)

	user.Get("/me", controllers.GetUserProfile)
	user.Put("/me", controllers.UpdateUserProfile)
	user.Delete("/me", controllers.DeleteUser)
}
