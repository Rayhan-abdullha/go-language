package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
func main() { // Entry point for the API application
	mux := http.NewServeMux()
	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))
	mux.Handle("GET /", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Welcome to the API Updated!")
		},
	))
	http.ListenAndServe(":3000", mux)
	fmt.Println("https://localhost:3000")
}
