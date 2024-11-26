package modules

import (
	"jusbrasil-tech-challenge/app/scrapper/repository/http"
	service "jusbrasil-tech-challenge/app/scrapper/service/scrapper"
	usecase "jusbrasil-tech-challenge/app/scrapper/usecase/scrapperprocess"

	"go.uber.org/fx"
)

var scrapperFactory = fx.Provide(
	http.NewScraperRepository,
	service.NewScrapperService,
	usecase.NewProcessScrapper,
)

var ScrapperModule = fx.Options(
	scrapperFactory,
)
