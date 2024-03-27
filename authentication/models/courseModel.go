package models

type Course struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

// Mock database for courses
var Courses = []Course{
    {ID: "1", Title: "Go Fundamentals", Description: "Learn the basics of Go programming."},
    {ID: "2", Title: "Advanced Go", Description: "Dive deeper into Go with advanced concepts and techniques."},
    // Add more courses as needed
}
