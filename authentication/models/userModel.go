// package models

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"` // In a real app, don't store passwords as plain text.
// }

// // Mock database
// var Users = []User{
// 	{Username: "test", Password: "test"},
// 	{Username: "piyush", Password: "12345"},
// }


package models

import (
    "errors"
    "regexp"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"` // Remember, don't store passwords as plain text in a real app.
}

var Users = []User{
    {Username: "test", Password: "test"}, // Example users, replace or remove as needed.
    {Username: "piyush", Password: "12345"},
}

// ValidateUsername checks if the username meets the requirements.
func ValidateUsername(username string) error {
    if len(username) < 3 || len(username) > 20 {
        return errors.New("username must be between 3 and 20 characters")
    }
    if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(username) {
        return errors.New("username can only contain alphanumeric characters, '-' and '_'")
    }
    return nil
}

// ValidatePassword checks if the password meets the requirements.
func ValidatePassword(password string) error {
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters long")
    }
    var (
        hasUpper   = regexp.MustCompile(`[A-Z]`)
        hasLower   = regexp.MustCompile(`[a-z]`)
        hasNumber  = regexp.MustCompile(`[0-9]`)
        hasSpecial = regexp.MustCompile(`[!@#\$%^&*]`)
    )
    if !hasUpper.MatchString(password) {
        return errors.New("password must include at least one uppercase letter")
    }
    if !hasLower.MatchString(password) {
        return errors.New("password must include at least one lowercase letter")
    }
    if !hasNumber.MatchString(password) {
        return errors.New("password must include at least one number")
    }
    if !hasSpecial.MatchString(password) {
        return errors.New("password must include at least one special character (!@#$%^&*)")
    }
    return nil
}
