package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/piyushkumar/authentication/authentication/models"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate username and password
	if err := models.ValidateUsername(newUser.Username); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := models.ValidatePassword(newUser.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add user to the database (mocked for this example)
	models.Users = append(models.Users, newUser)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User created")
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate user (mocked)
	for _, u := range models.Users {
		if u.Username == user.Username && u.Password == user.Password {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"exp":      time.Now().Add(time.Hour * 72).Unix(),
			})

			tokenString, error := token.SignedString([]byte("your_secret_key"))
			if error != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tokenString)
			return
		}
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("Invalid credentials")
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Access to protected resource granted")
}
