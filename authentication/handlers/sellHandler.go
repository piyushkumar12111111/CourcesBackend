package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/piyushkumar/authentication/authentication/models"
)

// GetAllItemsHandler handles GET requests to fetch all items
func GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")  //!  


	json.NewEncoder(w).Encode(models.Items)
}

// AddItemHandler handles POST requests to add a new item
func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")  
	var newItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newItem.ID = uuid.New().String() // Generate a new ID
	models.Items = append(models.Items, newItem)
	json.NewEncoder(w).Encode(newItem)
}
