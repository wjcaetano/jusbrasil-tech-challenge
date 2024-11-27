package scrapper

import (
	"encoding/json"
	"jusbrasil-tech-challenge/app/scrapper/usecase/scrapperprocess"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	scrapperUsecase scrapperprocess.ProcessScrapper
}

func NewHandler(usecase scrapperprocess.ProcessScrapper) *Handler {
	return &Handler{scrapperUsecase: usecase}
}

func (h *Handler) RegisterScrapperRoutes(router *chi.Mux) {
	router.Get("/scrapper", h.fetchScrapperHandler)
}

func (h *Handler) fetchScrapperHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	cases, err := h.scrapperUsecase.FetchAndParseCases(url)
	if err != nil {
		http.Error(w, "Failed to fetch and parse cases", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cases); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
