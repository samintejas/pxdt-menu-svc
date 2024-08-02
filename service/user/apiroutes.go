package user

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"projectx.io/drivethru/types"
	"projectx.io/drivethru/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("TODO : Login")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	u, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		log.Println(err)
	}

	if u == nil {
		log.Println("creating new user")
		newuser := types.User{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Password:  payload.Password,
		}
		h.store.CreateUser(&newuser)
	}
	// if it doesnt create new user
}
