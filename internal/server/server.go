package server

import (
	"context"
	"errors"
	"jusbrasil-tech-challenge/internal/config"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func StartHTTPServer(lx fx.Lifecycle, router *chi.Mux, cfg *config.Configuration) {
	server := &http.Server{
		Addr:              cfg.HTTPClient.Addr,
		Handler:           router,
		ReadHeaderTimeout: cfg.HTTPClient.ConnMaxLifetime,
	}

	lx.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				go func() {
					log.Printf("Starting server on %s", cfg.HTTPClient.Addr)
					if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
						log.Fatal("Error starting server: ", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Println("Shutting down server")
				return server.Shutdown(ctx)
			},
		},
	)
}
