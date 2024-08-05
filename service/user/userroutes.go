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
	router.HandleFunc("/users/{id}", h.handleView).Methods("GET")
	router.HandleFunc("/users/{id}", h.handleUpdate).Methods("PUT")
}

func (h *Handler) handleUpdate(w http.ResponseWriter, r *http.Request) {

	var payload types.UpdateUser
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	u, err := h.store.GetUserById(uint(id))
	if u == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	} else {

		updateuser := new(types.User)
		updateuser.ID = u.ID
		updateuser.FirstName = payload.FirstName
		updateuser.LastName = payload.LastName
		updateuser.UserName = payload.UserName
		updateuser.Email = payload.Email
		updateuser.Status = payload.Status
		updateuser.Password = payload.Password

		latestuser, err := h.store.UpdateUser(updateuser)

		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		utils.WriteJson(w, http.StatusOK, latestuser)
	}

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
	} else {
		utils.WriteJson(w, http.StatusOK, user)
	}

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("login functionality is under development"))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterUser
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	ex, err := h.store.ExcistsByUsernameAndEmail(payload.UserName, payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	if !ex {
		newuser := types.User{
			UserName:  payload.UserName,
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Password:  payload.Password,
			Status:    payload.Status,
		}
		userId, err := h.store.CreateUser(&newuser)
		if err != nil {
			log.Print(err)
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create new user"))
		} else {
			utils.WriteJson(w, http.StatusCreated, types.RegistedUser{ID: userId})
		}
	} else {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user/email already registered"))
	}
}
