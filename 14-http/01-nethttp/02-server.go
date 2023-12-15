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
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func (appServer *AppServer) Register(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	appServer.routes[pattern] = handler
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	if handler, exists := appServer.routes[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	http.Error(w, "Resource not found", http.StatusNotFound)
	return
}

// application specific logic
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

func main() {
	appServer := NewAppServer()
	appServer.Register("/", indexHandler)
	appServer.Register("/products", productsHandler)
	appServer.Register("/customers", customersHandler)
	http.ListenAndServe(":8080", appServer)
}
