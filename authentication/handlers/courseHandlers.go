package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/piyushkumar/authentication/authentication/models"
)

func GetAllCoursesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Courses)
}


func AddCourseHandler(w http.ResponseWriter, r *http.Request) {
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


func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
    var updatedCourse models.Course
    err := json.NewDecoder(r.Body).Decode(&updatedCourse)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate the updated course
    if err := models.ValidateCourse(updatedCourse); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Attempt to update the course
    if err := models.UpdateCourse(updatedCourse); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Course updated successfully")
}