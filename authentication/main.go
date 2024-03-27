package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piyushkumar/authentication/authentication/handlers"
	"github.com/piyushkumar/authentication/authentication/middleware"
)

func main() {
	r := mux.NewRouter()

	// Authentication routes
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
	r.HandleFunc("/signin", handlers.SignInHandler).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("", handlers.ProtectedHandler).Methods("GET")

	// Public routes
	r.HandleFunc("/courses", handlers.GetAllCoursesHandler).Methods("GET") // New line

	log.Fatal(http.ListenAndServe(":9080", r))
}
