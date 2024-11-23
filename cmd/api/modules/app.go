package modules

import (
	"jusbrasil-tech-challenge/internal/config"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	options := []fx.Option{
		InternalModule,
		ScrapperModule,
	}

	if !config.IsLocalScope() {
		options = append(options, fx.NopLogger)
	}

	return fx.New(
		fx.Options(options...),
	)
}
