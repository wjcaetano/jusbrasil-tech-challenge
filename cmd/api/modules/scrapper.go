package modules

import (
	"jusbrasil-tech-challenge/app/scrapper/repository/http"

	"go.uber.org/fx"
)

var scrapperFactory = fx.Provide(
	http.NewScraperRepository,
)

var ScrapperModule = fx.Options(
	scrapperFactory,
)
