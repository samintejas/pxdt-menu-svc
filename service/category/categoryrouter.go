package category

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

	router.HandleFunc("/categories", getCategories).Methods("GET")
	router.HandleFunc("/category", addCategory).Methods("POST")
	router.HandleFunc("/category/{id}", getCategory).Methods("GET")
	router.HandleFunc("/category/{id}", updateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", deleteCategory).Methods("DELETE")
}

func getCategories(w http.ResponseWriter, r *http.Request)  { /* implementation */ }
func addCategory(w http.ResponseWriter, r *http.Request)    { /* implementation */ }
func getCategory(w http.ResponseWriter, r *http.Request)    { /* implementation */ }
func updateCategory(w http.ResponseWriter, r *http.Request) { /* implementation */ }
func deleteCategory(w http.ResponseWriter, r *http.Request) { /* implementation */ }
