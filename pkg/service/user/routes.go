package user

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
	router.HandleFunc("/", h.handlerGet).Methods("GET")
	router.HandleFunc("/login", h.handlerLogin).Methods("POST")
	router.HandleFunc("/register", h.handlerRegister).Methods("POST")
}

func (h *Handler) handlerGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
