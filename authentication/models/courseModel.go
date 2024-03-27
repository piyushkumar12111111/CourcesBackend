package models

import (
    "errors"
    "regexp"
)

type Course struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

// Mock database for courses
var Courses = []Course{
    {ID: "1", Title: "Go Fundamentals", Description: "Learn the basics of Go programming."},
    // Add more predefined courses here
}

// ValidateCourse checks if the course details meet the requirements.
func ValidateCourse(course Course) error {
    if course.ID == "" || course.Title == "" || course.Description == "" {
        return errors.New("all fields must be filled")
    }
    if !regexp.MustCompile(`^\w+$`).MatchString(course.ID) {
        return errors.New("ID can only contain alphanumeric characters and underscores")
    }
    return nil
}
