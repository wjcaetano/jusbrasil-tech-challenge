package scrapperprocess

import "jusbrasil-tech-challenge/app/scrapper"

type ServiceScrapper interface {
	GetLegalCases(url string) ([]scrapper.LegalCase, error)
}
