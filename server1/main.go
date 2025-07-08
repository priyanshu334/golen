package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Simple response
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Register the route
	http.HandleFunc("/", homeHandler)

	// Start the server on port 8080
	fmt.Println("🚀 Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("❌ Error starting server:", err)
	}
}
