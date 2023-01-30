package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/store/auth"
	"github.com/store/database"
	"github.com/store/router"

	_ "github.com/lib/pq"
)

func main() {
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	authCtx := context.Background()
	authPort := "8081"
	authRouter := auth.SetupServer(authCtx, "http://localhost:"+authPort)
	authServer := &http.Server{
		Addr:    ":" + authPort,
		Handler: authRouter,
	}
	go func() {
		err := authServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	app.Listen(":8080")
	<-authCtx.Done()
}
