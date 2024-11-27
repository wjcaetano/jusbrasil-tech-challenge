package router

import (
	"jusbrasil-tech-challenge/app/scrapper/entrypoint/rest/scrapper"
	"jusbrasil-tech-challenge/internal/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(scrapperHandler *scrapper.Handler) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middlewareLogger)
	router.Use(middlewareRecoverer)

	routes.RegisterRoutes(router, scrapperHandler)

	return router
}

func middlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func middlewareRecoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
