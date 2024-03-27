package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/piyushkumar/authentication/authentication/models"
)

func AddContactHandler(w http.ResponseWriter, r *http.Request) {
    var newContact models.Contact
    if err := json.NewDecoder(r.Body).Decode(&newContact); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    models.Contacts = append(models.Contacts, newContact)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(newContact)
}

func GetAllContactsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Contacts)
}
