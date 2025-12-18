package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)
	r.Post("/", h.Create)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)

	return r
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("GetAll failed")
		http.Error(w, "internal error", 500)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "internal error", 500)
		return
	}

	if user == nil {
		http.Error(w, "not found", 404)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)

	err := h.service.Create(r.Context(), &u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(u)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	u.ID = id

	err := h.service.Update(r.Context(), &u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(204)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	h.service.Delete(r.Context(), id)
	w.WriteHeader(204)
}
