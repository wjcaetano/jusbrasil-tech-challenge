package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gocolly/colly/v2"
)

var (
	userAgentURL = "https://httpbin.org/user-agent"
	allowedURLs  = map[string]struct{}{
		userAgentURL: {},
	}
)

type ScraperRepository struct {
	collector *colly.Collector
}

func NewScraperRepository() *ScraperRepository {
	userAgent, _ := getUserAgent()
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	return &ScraperRepository{collector: c}
}

func (r *ScraperRepository) FetchPage(url string) (string, error) {
	if ok := isValidURL(url); !ok {
		return "", errors.New("url is not allowed")
	}
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

func getUserAgent() (string, error) {
	if ok := isValidURL(userAgentURL); !ok {
		return "", errors.New("user agent url is not allowed")
	}
	resp, err := http.Get(userAgentURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.Header.Get("User-Agent"), nil
}

func isValidURL(url string) bool {
	if _, ok := allowedURLs[url]; !ok {
		return false
	}
	return true
}
