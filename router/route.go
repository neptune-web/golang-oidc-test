package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/store/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/stores")
	v2 := api.Group("/testdata")

	// routes
	v1.Post("/", handler.CreateStore)
	v1.Get("/", handler.GetAllStores)
	v1.Get("/:store_id", handler.GetSingleStore)

	v2.Get("/", handler.GetSingleTestdatas)
	v2.Get("/:request_id", handler.GetSingleTestdata)
	v2.Post("/", handler.CreateTestdata)
}
