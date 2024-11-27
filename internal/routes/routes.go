package routes

import (
	"jusbrasil-tech-challenge/app/scrapper/entrypoint/rest/scrapper"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, scrapperHandler *scrapper.Handler) {
	router.Get("/", homeHandler)
	router.Get("/health", healthHandler)

	scrapperHandler.RegisterScrapperRoutes(router)
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Println(err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
