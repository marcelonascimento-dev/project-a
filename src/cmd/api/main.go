package main

import (
	"log"
	"project-a/handlers"
	"project-a/routes"
	"project-a/storage"
	"project-a/types"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	store := session.New(session.Config{
		CookieSecure:   true,
		CookieHTTPOnly: true,
	})

	fiberapp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	fiberapp.Use(favicon.New())
	fiberapp.Use(recover.New())
	fiberapp.Use("/api", handlers.WithAuthenticatedUser)

	appcontext := types.ApplicationContext{
		FiberApp:     fiberapp,
		DB:           storage.NewDatabase(),
		SessionStore: store,
	}

	handlers.CreateProductHandler(&appcontext)
	handlers.CreateAuthHandler(&appcontext)

	routes.SetupRoutes(&appcontext)

	log.Fatal(appcontext.FiberApp.Listen(":3333"))
}
