package http

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type ScrapperRepository struct {
	collector *colly.Collector
}

func NewScrapperRepository() *ScrapperRepository {
	c := colly.NewCollector()

	return &ScrapperRepository{collector: c}
}

func (r *ScrapperRepository) FetchPage(url string) (string, error) {
	var content string

	r.collector.OnResponse(func(r *colly.Response) {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			fmt.Printf("Error parsing HTML document. %s\n", err)
			return
		}

		table := doc.Find("#divDadosResultado-A table tbody")

		html, err := table.Html()
		if err != nil {
			fmt.Printf("Error parsing HTML from table. %s\n", err)
			return
		}

		content = html
	})

	err := r.collector.Visit(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch page: %w", err)
	}

	return content, nil
}
