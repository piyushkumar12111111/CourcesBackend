package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piyushkumar/authentication/authentication/middleware"
	"github.com/piyushkumar/authentication/authentication/handlers"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
    r.HandleFunc("/signin", handlers.SignInHandler).Methods("POST")
    protected := r.PathPrefix("/protected").Subrouter()
    protected.Use(middleware.AuthMiddleware)
    protected.HandleFunc("", handlers.ProtectedHandler).Methods("GET")

    log.Fatal(http.ListenAndServe(":9080", r))
}
