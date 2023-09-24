package main

import (
	"bytes"
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
	// Data postingan yang akan disimpan
	newPost := Post{
		UserID: 12,
		ID:     212,
		Title:  "Test API",
		Body:   "Ini test API dengan menggunakan POST",
	}

	// Mengkonversi data postingan menjadi format JSON
	postData, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Membuat permintaan POST ke API
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("HTTP Error: %s\n", resp.Status)
		return
	}

	// Menampilkan respons dari server (biasanya berisi data postingan yang telah disimpan)
	var createdPost Post
	err = json.NewDecoder(resp.Body).Decode(&createdPost)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Postingan baru telah disimpan:")
	fmt.Printf("User ID: %d\n", createdPost.UserID)
	fmt.Printf("ID: %d\n", createdPost.ID)
	fmt.Printf("Title: %s\n", createdPost.Title)
	fmt.Printf("Body: %s\n", createdPost.Body)
}
