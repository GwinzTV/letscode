package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"letscode/project-01-url-shortener/internal/store"
)

type Handler struct {
	store store.Store
	rnd   *rand.Rand
}

func NewHandler(s store.Store) *Handler {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Handler{store: s, rnd: r}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", h.shorten)
	mux.HandleFunc("/health", h.health)
	mux.HandleFunc("/", h.redirect)
	return mux
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Code string `json:"code"`
}

func (h *Handler) shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// validate URL: must parse, have host and https scheme
if err := ValidateURL(req.URL); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}
	code := h.generateCode(6)
	if err := h.store.Save(code, req.URL); err != nil {
		log.Printf("save error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortenResponse{Code: code})
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	code := r.URL.Path
	if len(code) > 0 && code[0] == '/' {
		code = code[1:]
	}
	if code == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("URL Shortener: POST /shorten with JSON {\"url\": \"https://...\"}"))
		return
	}
	if dest, ok := h.store.Lookup(code); ok {
		http.Redirect(w, r, dest, http.StatusFound)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (h *Handler) generateCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[h.rnd.Intn(len(letters))]
	}
	return string(b)
}
