package modules

import (
	"jusbrasil-tech-challenge/app/scrapper/entrypoint/rest/scrapper"
	repository "jusbrasil-tech-challenge/app/scrapper/repository/http"
	service "jusbrasil-tech-challenge/app/scrapper/service/scrapper"
	usecase "jusbrasil-tech-challenge/app/scrapper/usecase/scrapperprocess"

	"go.uber.org/fx"
)

var scrapperFactory = fx.Options(
	fx.Provide(
		repository.NewScrapperRepository,
		fx.Annotate(
			repository.NewScrapperRepository,
			fx.As(new(service.RepositoryScrapper)),
		),
	),
	fx.Provide(
		fx.Annotate(
			service.NewScrapperService,
			fx.As(new(usecase.ServiceScrapper)),
		),
	),
	fx.Provide(
		usecase.NewProcessScrapper,
		scrapper.NewHandler,
	),
)

var ScrapperModule = fx.Options(
	scrapperFactory,
)
