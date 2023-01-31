package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

func main() {
	// database.Connect()

	// app := fiber.New()

	// app.Use(logger.New())

	// app.Use(cors.New())

	// router.SetupRoutes(app)

	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.SendStatus(404) // => 404 "Not Found"
	// })

	// authCtx := context.Background()
	// authPort := "8081"
	// authRouter := auth.SetupServer(authCtx, "http://localhost:"+authPort)
	// authServer := &http.Server{
	// 	Addr:    ":" + authPort,
	// 	Handler: authRouter,
	// }
	// go func() {
	// 	err := authServer.ListenAndServe()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// app.Listen(":8080")
	// <-authCtx.Done()
	app := fiber.New()

	goth.UseProviders(
		google.New(os.Getenv("OAUTH_KEY"), os.Getenv("OAUTH_SECRET"), "http://127.0.0.1:8088/auth/callback"),
	)

	app.Get("/login", goth_fiber.BeginAuthHandler)
	app.Use("/auth/callback", func(ctx *fiber.Ctx) {
		user, err := goth_fiber.CompleteUserAuth(ctx)
		if err != nil {
			log.Fatal(err)
		}

		ctx.Send(user)
	})

	app.Use("/logout", func(c *fiber.Ctx) error {
		return c.SendString("logout") // => 404 "Not Found"
	})

	if err := app.Listen(":8088"); err != nil {
		log.Fatal(err)
	}
}
