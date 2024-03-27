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
	w.Header().Set("Access-Control-Allow-Origin", "*") //!

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


func UpdateItemHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var updatedItem models.Item

	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, item := range models.Items {
		if item.ID == updatedItem.ID {
			models.Items[i] = updatedItem
			json.NewEncoder(w).Encode("Item updated successfully")
			return
		}
	}

	json.NewEncoder(w).Encode("Item not found")


}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := r.URL.Query()

	id := params.Get("id")

	for i, item := range models.Items {

		if item.ID == id {
			models.Items = append(models.Items[:i], models.Items[i+1:]...)

			json.NewEncoder(w).Encode("Item deleted Succesfully")
			break
		}
	}

	json.NewEncoder(w).Encode("Item not found")

}
