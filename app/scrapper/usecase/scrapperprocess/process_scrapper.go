package scrapperprocess

import (
	"jusbrasil-tech-challenge/app/scrapper"
	service "jusbrasil-tech-challenge/app/scrapper/service/scrapper"
)

//go:generate mockery --all --output=./mocks --outpkg=mocks --with-expecter

type (
	ProcessScrapper interface {
		FetchAndParseCases(url string) ([]scrapper.LegalCase, error)
	}

	processScrapper struct {
		service service.ScrapperService
	}
)

func NewProcessScrapper(service service.ScrapperService) ProcessScrapper {
	return &processScrapper{service: service}
}

func (p *processScrapper) FetchAndParseCases(url string) ([]scrapper.LegalCase, error) {
	cases, err := p.service.GetLegalCases(url)
	if err != nil {
		return nil, err
	}
	return cases, nil
}
