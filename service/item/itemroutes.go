package item

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"projectx.io/drivethru/store"
	"projectx.io/drivethru/utils"
)

type Handler struct {
	store store.ItemStore
}

func NewHandler(store store.ItemStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/items", h.getItems).Methods("GET")
	router.HandleFunc("/item", h.addItem).Methods("POST")
	router.HandleFunc("/item/{id}", h.getItem).Methods("GET")
	router.HandleFunc("/item/{id}", h.updateItem).Methods("PUT")
	router.HandleFunc("/item/{id}", h.deleteItem).Methods("DELETE")
}

func (h *Handler) getItems(w http.ResponseWriter, r *http.Request) {
	item, err := h.store.GetAllItems()

	if item == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	} else {
		utils.WriteJson(w, http.StatusOK, item)
	}

}

func (h *Handler) addItem(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getItem(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {

}
