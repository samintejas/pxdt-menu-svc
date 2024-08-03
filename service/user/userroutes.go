package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"projectx.io/drivethru/store"
	"projectx.io/drivethru/types"
	"projectx.io/drivethru/utils"
)

type Handler struct {
	store store.UserStore
}

func NewHandler(store store.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/view/{id}", h.handleView).Methods("GET")
}

func (h *Handler) handleView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	user, err := h.store.GetUserById(uint(id))

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}

	utils.WriteJson(w, http.StatusAccepted, user)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("login functionality is under development"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterUser
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	u, err := h.store.GetUserByEmail(payload.Email)

	if err != nil {
		log.Println(err)
	}

	if u == nil {
		newuser := types.User{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Password:  payload.Password,
		}
		userId, err := h.store.CreateUser(&newuser)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create new user"))
		}
		utils.WriteJson(w, http.StatusAccepted, types.RegistedUser{ID: userId})
	} else {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("email already registered"))
	}
}
