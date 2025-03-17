package routes

import (
	"evermos-backend/controllers"
	"evermos-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func UploadRoutes(app *fiber.App) {
	upload := app.Group("/upload")

	upload.Post("/products/:id/image", middleware.AuthMiddleware, controllers.UploadProductImage)
	upload.Post("/users/avatar", middleware.AuthMiddleware, controllers.UploadUserAvatar)
	upload.Post("/stores/:id/logo", middleware.AuthMiddleware, controllers.UploadStoreLogo)
}
