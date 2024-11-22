package http

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/http"
)

var userAgentURL = "https://httpbin.org/user-agent"

type ScraperRepository struct {
	collector *colly.Collector
}

func NewScraperRepository() *ScraperRepository {
	c := colly.NewCollector(
		colly.UserAgent(getUserAgent()),
	)

	return &ScraperRepository{collector: c}
}

func (r *ScraperRepository) FetchPage(url string) (string, error) {
	var result string
	r.collector.OnHTML("body", func(e *colly.HTMLElement) {
		result = e.Text
	})

	err := r.collector.Visit(url)
	if err != nil {
		return "", fmt.Errorf("failed to scrape page: %w", err)
	}
	return result, nil
}

func getUserAgent() string {
	resp, err := http.Get(userAgentURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.Header.Get("User-Agent")
}
