package modules

import (
	"jusbrasil-tech-challenge/app/scrapper/entrypoint/rest/scrapper"
	"jusbrasil-tech-challenge/internal/config"
	"jusbrasil-tech-challenge/internal/db"
	"jusbrasil-tech-challenge/internal/router"
	"jusbrasil-tech-challenge/internal/routes"
	"jusbrasil-tech-challenge/internal/server"

	"github.com/go-chi/chi/v5"

	"go.uber.org/fx"
)

var internalFactory = fx.Provide(
	config.NewConfig,
	db.NewDatabase,
	router.NewRouter,
)

var InternalModule = fx.Options(
	internalFactory,
	fx.Invoke(
		server.StartHTTPServer,
		func(router *chi.Mux, scrapperHandler *scrapper.Handler) {
			routes.RegisterRoutes(router, scrapperHandler)
		},
	),
)
