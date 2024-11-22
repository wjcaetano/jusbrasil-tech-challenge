package modules

import (
	"go.uber.org/fx"
	"jusbrasil-tech-challenge/app/scrapper/repository/http"
)

var scrapperFactory = fx.Provide(
	http.NewScraperRepository,
)

var ScrapperModule = fx.Options(
	scrapperFactory,
)
