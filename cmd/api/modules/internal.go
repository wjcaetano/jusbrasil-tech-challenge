package modules

import (
	"jusbrasil-tech-challenge/internal/config"
	"jusbrasil-tech-challenge/internal/db"
	"jusbrasil-tech-challenge/internal/router"
	"jusbrasil-tech-challenge/internal/server"

	"go.uber.org/fx"
)

var internalFactory = fx.Provide(
	config.NewConfig,
	db.NewDatabase,
	router.NewRouter,
)

var InternalModule = fx.Options(
	internalFactory,
	fx.Invoke(server.StartHTTPServer),
)
