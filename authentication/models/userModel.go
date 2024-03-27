package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"` // In a real app, don't store passwords as plain text.
}

// Mock database
var Users = []User{
	{Username: "test", Password: "test"},
	{Username: "piyush", Password: "12345"},
}
