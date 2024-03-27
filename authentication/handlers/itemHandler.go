package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/piyushkumar/authentication/authentication/models"
)

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
    pageStr := r.URL.Query().Get("page")
    limitStr := r.URL.Query().Get("limit")

    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        page = 1
    }

    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit < 1 {
        limit = 10 // Default limit
    }

    startIndex := (page - 1) * limit
    endIndex := startIndex + limit

    if startIndex > len(models.Items) {
        json.NewEncoder(w).Encode([]models.Item{}) // Empty response
        return
    }

    if endIndex > len(models.Items) {
        endIndex = len(models.Items)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Items[startIndex:endIndex])
}
