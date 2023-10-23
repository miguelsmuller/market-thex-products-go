package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupMiddlewares(app *fiber.App) {
    app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max: 3,
	}))

	app.Use(logger.New())
}
