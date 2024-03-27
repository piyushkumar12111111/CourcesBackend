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
