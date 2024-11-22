package http

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly/v2"
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
	r.collector.OnResponse(func(resp *colly.Response) {
		result = string(resp.Body)
	})

	fmt.Println(result)

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
