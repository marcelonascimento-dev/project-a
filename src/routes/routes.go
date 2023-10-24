package routes

import (
	"project-a/handlers"
	"project-a/types"
)

func SetupRoutes(appcontext *types.ApplicationContext) {
	gauth := appcontext.FiberApp.Group("/auth")
	gapi := appcontext.FiberApp.Group("/api")
	gauth.Get("/callback", handlers.GoogleCallbackHandler)
	gauth.Get("/login", handlers.GoogleLoginHandler)
	gapi.Get("/product/", handlers.GetProductsBasicHandler)

}
