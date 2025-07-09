package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello form dockerized go")
}
func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Server started at localhost:8000")
	http.ListenAndServe("localhost:8000", nil)
}
