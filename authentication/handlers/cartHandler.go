package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/piyushkumar/authentication/authentication/models"
)

func AddItemCArt(w http.ResponseWriter, r *http.Request) {
   


	var newCourse models.Course
	err := json.NewDecoder(r.Body).Decode(&newCourse)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the new course
	if err := models.ValidateCourse(newCourse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add the new course to the database (mocked for this example)
	models.Courses = append(models.Courses, newCourse)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("New course added")


}