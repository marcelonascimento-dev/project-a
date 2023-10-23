package main

import (
	"project-a/handlers"
	"project-a/types"
)

func SetupRoutes(appcontext *types.ApplicationContext) {
	appcontext.FiberApp.Get("/auth/google/login", handlers.GoogleLoginHandler)
	appcontext.FiberApp.Get("/auth/google/callback", handlers.GoogleCallbackHandler)
}
