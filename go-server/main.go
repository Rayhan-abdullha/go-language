package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	fmt.Fprintln(w, "Welcome to the Go Server!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "This is a About page", http.StatusOK)
}

type IProduct struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []IProduct{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Smartphone", Price: 499.99},
	{ID: 3, Name: "Tablet", Price: 299.99},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(products)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
func createProduct(w http.ResponseWriter, r *http.Request) {
	// writ e your code here
	var newProduct IProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Product created successfully")
}
func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if r.Method == "GET" {
		fmt.Println(r.Header)
		getProducts(w, r)
		return
	}
	if r.Method == "POST" {
		createProduct(w, r)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", productHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("http://localhost:8080")
	}
}
