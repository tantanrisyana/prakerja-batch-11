package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// Membuat permintaan GET ke API
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Membaca dan memparsing respons JSON
	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Menampilkan data
	for _, post := range posts {
		fmt.Printf("User ID: %d\n", post.UserID)
		fmt.Printf("ID: %d\n", post.ID)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Body: %s\n", post.Body)
		fmt.Println("------------")
	}
}
