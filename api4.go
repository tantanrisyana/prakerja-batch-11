package main

import (
	"fmt"
	"net/http"
)

func main() {
	var ID int16
	// ID postingan yang akan dihapus
	fmt.Print("Masukkan Posting ID: ")
	fmt.Scan(&ID)

	postID := ID // Ganti dengan ID postingan yang ingin Anda hapus

	// Membuat permintaan DELETE ke API
	fmt.Print(postID)
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending DELETE request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Error: %s\n", resp.Status)
		return
	}

	fmt.Printf("Post with ID %d has been deleted.\n", postID)
}
