package main

import (
	"thex-products/configurations"
	"thex-products/initializers"
	"thex-products/middlewares"
	"thex-products/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	cfg, err := configurations.LoadEnvVars()
    if err != nil {
        log.Fatal("Failed to load environment variables: ", err)
    }

	app := fiber.New()

	routes.SetupRoutes(app)
	middlewares.SetupMiddlewares(app)

	log.Fatal(app.Listen(cfg.APIAddress))
}
