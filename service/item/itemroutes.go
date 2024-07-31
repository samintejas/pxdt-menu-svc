package item

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/item", addItem).Methods("POST")
	router.HandleFunc("/item/{id}", getItem).Methods("GET")
	router.HandleFunc("/item/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")
}

func getItems(w http.ResponseWriter, r *http.Request)   { /* implementation */ }
func addItem(w http.ResponseWriter, r *http.Request)    { /* implementation */ }
func getItem(w http.ResponseWriter, r *http.Request)    { /* implementation */ }
func updateItem(w http.ResponseWriter, r *http.Request) { /* implementation */ }
func deleteItem(w http.ResponseWriter, r *http.Request) { /* implementation */ }
