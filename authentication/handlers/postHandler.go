package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/piyushkumar/authentication/authentication/models"
)

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
    var newPost models.Post
    err := json.NewDecoder(r.Body).Decode(&newPost)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Generate a new UUID for the post. In a real application, you might want to check for collisions.
    newPost.ID = uuid.New().String()

    models.Posts = append(models.Posts, newPost)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newPost)
}


func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Posts)
}


func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postID := vars["id"]

    err := models.DeletePost(postID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Post deleted successfully"))
}