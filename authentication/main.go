package main

import (
	"fmt"
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

	//! api for adding new course
	r.HandleFunc("/courses", handlers.AddCourseHandler).Methods("POST") // Add new course

	//! update course api
	r.HandleFunc("/courses/{id}", handlers.UpdateCourseHandler).Methods("PUT") // New line for updating a course

	//! Profile routes
	r.HandleFunc("/profile", handlers.AddProfileHandler).Methods("POST")
	r.HandleFunc("/profile/{userID}", handlers.GetProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{userID}", handlers.DeleteProfileHandler).Methods("DELETE")

	//! Contact routes
	// Contact routes
	r.HandleFunc("/contacts", handlers.AddContactHandler).Methods("POST")
	r.HandleFunc("/contacts", handlers.GetAllContactsHandler).Methods("GET") // Although you asked for POST, typically, retrieval uses GET

	//! Teacher routes
	r.HandleFunc("/courses/{courseID}/teachers", handlers.AddTeacherToCourseHandler).Methods("POST")
	r.HandleFunc("/courses/{courseID}/teachers", handlers.GetTeachersOfCourseHandler).Methods("GET")

	//! Post routes
	r.HandleFunc("/posts", handlers.AddPostHandler).Methods("POST")
	r.HandleFunc("/posts", handlers.GetAllPostsHandler).Methods("GET")

	r.HandleFunc("/posts/{id}", handlers.DeletePostHandler).Methods("DELETE")

	//! Register the route for getting items with pagination
	r.HandleFunc("/items", handlers.GetItemsHandler).Methods("GET")

	//!  Selling routhes
	r.HandleFunc("/selling", handlers.GetAllItemsHandler).Methods("GET")
	r.HandleFunc("/selling", handlers.AddItemHandler).Methods("POST")
	r.HandleFunc("/selling/{id}", handlers.DeleteItemHandler).Methods("DELETE")
	r.HandleFunc("/selling/{id}", handlers.UpdateItemHandler).Methods("PUT")

	//! Start the server

	fmt.Println("Server is running on port 9050")
	log.Fatal(http.ListenAndServe(":9050", r))

}
