package main

import (
	"log"
	"net/http"
	"project-a/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/auth/google/login", handlers.GoogleLoginHandler)
	r.HandleFunc("/auth/google/callback", handlers.GoogleCallbackHandler)

	log.Println("Server is running at http://localhost:3333")
	log.Println("Login http://localhost:3333/auth/google/login")

	log.Fatal(http.ListenAndServe(":3333", r))
}
