package http

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type ScraperRepository struct {
	collector *colly.Collector
}

func NewScraperRepository() *ScraperRepository {
	c := colly.NewCollector()

	return &ScraperRepository{collector: c}
}

func (r *ScraperRepository) FetchPage(url string) (string, error) {
	var result string
	r.collector.OnResponse(func(resp *colly.Response) {
		result = string(resp.Body)
	})

	err := r.collector.Visit(url)
	if err != nil {
		return "", fmt.Errorf("failed to scrape page: %w", err)
	}
	return result, nil
}
