package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GAZIMAGomeDDD/url-shortener/internal/storage"
	"github.com/GAZIMAGomeDDD/url-shortener/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4"
)

type Handler struct {
	mux   *chi.Mux
	store storage.StorageIface
}

func NewHandler(s storage.StorageIface) *Handler {
	return &Handler{
		mux:   chi.NewRouter(),
		store: s,
	}
}

func (h *Handler) createShortenedURL(w http.ResponseWriter, r *http.Request) {
	var body struct {
		URL string `json:"url"`
	}

	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err := utils.ValidateURL(body.URL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slug, err := h.store.CreateShortenedURL(r.Context(), body.URL)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	shortenedURL := fmt.Sprintf("http://localhost:8080/%s", slug)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"shortened_url": shortenedURL})
}

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	url, err := h.store.GetURL(r.Context(), slug)
	switch err {
	case nil:
		http.Redirect(w, r, url, http.StatusSeeOther)
	case pgx.ErrNoRows:
		http.NotFound(w, r)
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (h *Handler) Init() *chi.Mux {
	h.mux.Post("/", h.createShortenedURL)
	h.mux.Get("/{slug}", h.redirect)

	return h.mux
}
