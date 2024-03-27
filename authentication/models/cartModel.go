package models

type cart struct {
    ID    string  `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

// Mock database for items
var Cartitems = []Selling{

	{ID: "1", Name: "Laptop", Price: 1000},
	{ID: "2", Name: "Mobile", Price: 500},
	{ID: "3", Name: "Tablet", Price: 300},
	{ID: "4", Name: "Iphone", Price: 600},
}
