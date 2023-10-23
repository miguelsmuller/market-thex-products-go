package routes

import (
	"thex-products/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/health", controllers.HandlerHealth)

    app.Route("/products", func(router fiber.Router) {
		router.Post("/", controllers.CreateProduct)
		router.Get("", controllers.ListProducts)
	})
	app.Route("/products/:productId", func(router fiber.Router) {
		router.Delete("", controllers.DeleteProduct)
		router.Get("", controllers.GetProduct)
		router.Put("", controllers.UpdateProduct)
	})
}
