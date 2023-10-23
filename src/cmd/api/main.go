package main

import (
	"log"
	"project-a/types"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	appcontext := types.ApplicationContext{
		FiberApp: fiber.New(),
	}

	SetupRoutes(&appcontext)
	
	log.Fatal(appcontext.FiberApp.Listen(":3333"))
}


