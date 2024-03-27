package models

import "fmt"

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Initial dummy data for posts
var Posts = []Post{
	{
		ID:      "1",
		Title:   "Introduction to Go",
		Content: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
		Author:  "Alice",
	},
	{
		ID:      "2",
		Title:   "Understanding Goroutines",
		Content: "Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as lightweight threads.",
		Author:  "Bob",
	},
	{
		ID:      "3",
		Title:   "Channels in Go",
		Content: "Channels are a typed conduit through which you can send and receive values with the channel operator.",
		Author:  "Charlie",
	},
}

func DeletePost(postID string) error {
	for i, post := range Posts {
		if post.ID == postID {
			Posts = append(Posts[:i], Posts[i+1:]...)
			return nil // Post found and deleted successfully
		}
	}
	return fmt.Errorf("post not found") // Post not found
}
