package models

import (
    "errors"
    "fmt"
)

type Item struct {
    ID    string  `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

// Mock database for items with some dummy data
var Items = []Item{
    {ID: "1", Name: "Item 1", Price: 99.99},
    {ID: "2", Name: "Item 2", Price: 199.99},
    {ID: "3", Name: "Item 3", Price: 299.99},
    // Add more items as needed
}

// ValidateItem checks that an item has all necessary fields and valid data.
func ValidateItem(item Item) error {
    if item.ID == "" {
        return errors.New("item ID cannot be empty")
    }
    if item.Name == "" {
        return errors.New("item name cannot be empty")
    }
    if item.Price < 0 {
        return fmt.Errorf("item price cannot be negative, got: %.2f", item.Price)
    }
    return nil
}
