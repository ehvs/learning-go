package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer handles HTTP requests
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players")
	fmt.Println("Extracted player:", player) // Print player to console
	fmt.Fprintf(w, "Player: %s", player)     // Respond with player
}

func main() {
	http.HandleFunc("/", PlayerServer) // Register the handler function
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil) // Start the server
}
