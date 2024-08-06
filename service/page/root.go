package page

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	RootTemplate  *template.Template
	ErrorTempalte *template.Template
}

func NewHandler(rtpl *template.Template, etpl *template.Template) *Handler {
	return &Handler{
		RootTemplate:  rtpl,
		ErrorTempalte: etpl,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.handleRoot)
	router.HandleFunc("/error", h.handleError)
}

func (h *Handler) handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	h.RootTemplate.Execute(w, nil)
}

func (h *Handler) handleError(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	h.ErrorTempalte.Execute(w, nil)
}
