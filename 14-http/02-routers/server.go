package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float32 `json:"cost"`
	Category string  `json:"category"`
	// Category string `json:"-"`
}

var products = []Product{
	{100, "Pen", 10, "Stationary"},
	{101, "Pencil", 10, "Stationary"},
	{102, "Marker", 10, "Stationary"},
	{103, "Charger", 10, "Electronics"},
}

// application specific logic
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// product handlers
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
func saveNewProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

// middlewares
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("%s - %s\n", r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/products", getAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/products", saveNewProduct).Methods(http.MethodPost)
	r.HandleFunc("/customers", customersHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
