package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

type AppServer struct {
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		// fmt.Fprintln(w, "All the products will be served")
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "unsupported request method", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "All the customers will be served")
	default:
		http.Error(w, "Resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
