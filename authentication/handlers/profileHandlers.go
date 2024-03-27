package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piyushkumar/authentication/authentication/models"
)

func AddProfileHandler(w http.ResponseWriter, r *http.Request) {
    var profile models.Profile
    if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    models.AddProfile(profile)
    json.NewEncoder(w).Encode(profile)
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["userID"]
    profile, err := models.GetProfile(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(profile)
}

func DeleteProfileHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["userID"]
    if err := models.DeleteProfile(userID); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Profile deleted successfully")
}
