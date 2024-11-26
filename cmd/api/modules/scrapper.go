package modules

import (
	"jusbrasil-tech-challenge/app/scrapper/repository/http"
	service "jusbrasil-tech-challenge/app/scrapper/service/scrapper"

	"go.uber.org/fx"
)

var scrapperFactory = fx.Provide(
	http.NewScraperRepository,
	service.NewScrapperService,
)

var ScrapperModule = fx.Options(
	scrapperFactory,
)
