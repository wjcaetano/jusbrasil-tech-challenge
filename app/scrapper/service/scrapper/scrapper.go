package scrapper

import (
	"fmt"
	entity "jusbrasil-tech-challenge/app/scrapper"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//go:generate mockery --all --output=./mocks --outpkg=mocks --with-expecter

type (
	RepositoryScrapper interface {
		FetchPage(url string) (string, error)
	}

	ServiceScrapper interface {
		GetLegalCases(url string) ([]entity.LegalCase, error)
	}

	serviceScrapper struct {
		repo RepositoryScrapper
	}
)

func NewScrapperService(repo RepositoryScrapper) ServiceScrapper {
	return &serviceScrapper{repo: repo}
}

func (s *serviceScrapper) GetLegalCases(url string) ([]entity.LegalCase, error) {
	content, err := s.repo.FetchPage(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch page: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse document: %w", err)
	}

	return s.parseLegalCases(doc), nil
}

func (s *serviceScrapper) parseLegalCases(doc *goquery.Document) []entity.LegalCase {
	var cases []entity.LegalCase
	var currentCase entity.LegalCase

	currentCase, cases = casesExtract(doc, s)

	if isCaseFound(currentCase) {
		cases = append(cases, currentCase)
	}
	return cases
}

func casesExtract(doc *goquery.Document, s *serviceScrapper) (currentCase entity.LegalCase, cases []entity.LegalCase) {
	doc.Find("tr").Each(func(_ int, selection *goquery.Selection) {
		currentCase = s.processCaseRow(selection, currentCase, &cases)
	})
	return currentCase, cases
}

func (s *serviceScrapper) extractCaseNumber(row *goquery.Selection) string {
	caseNumber := strings.TrimSpace(row.Find("td a:first-of-type").Text())
	matched, _ := regexp.MatchString(`^\d{7}-\d{2}\.\d{4}\.\d\.\d{2}\.\d{4}$`, caseNumber)
	if matched {
		return caseNumber
	}
	return ""
}

func (s *serviceScrapper) extractSummary(row *goquery.Selection) string {
	var summary string

	row.Find("td div").EachWithBreak(func(_ int, selection *goquery.Selection) bool {
		strongText := strings.TrimSpace(selection.Find("strong").Text())
		if strings.Contains(strongText, "Ementa:") {
			selection.Find("strong").Remove()
			summary = strings.TrimSpace(selection.Text())
			summary = cleanText(summary)
			return false
		}
		return true
	})
	return summary
}

func (s *serviceScrapper) extractReporter(row *goquery.Selection) string {
	var reporter string
	row.Find("td").EachWithBreak(func(_ int, selection *goquery.Selection) bool {
		strongText := strings.TrimSpace(selection.Find("strong").Text())
		if strings.Contains(strongText, "Relator(a):") {
			selection.Find("strong").Remove()
			reporter = strings.TrimSpace(selection.Text())
			return false
		}
		return true
	})
	return reporter
}

func (s *serviceScrapper) extractCourt(row *goquery.Selection) string {
	var reporter string
	row.Find("td").EachWithBreak(func(_ int, selection *goquery.Selection) bool {
		strongText := strings.TrimSpace(selection.Find("strong").Text())
		if strings.Contains(strongText, "Comarca:") {
			selection.Find("strong").Remove()
			reporter = strings.TrimSpace(selection.Text())
			return false
		}
		return true
	})
	return reporter
}

func (s *serviceScrapper) extractJudgingBody(row *goquery.Selection) string {
	var reporter string
	row.Find("td").EachWithBreak(func(_ int, selection *goquery.Selection) bool {
		strongText := strings.TrimSpace(selection.Find("strong").Text())
		if strings.Contains(strongText, "Órgão julgador:") {
			selection.Find("strong").Remove()
			reporter = strings.TrimSpace(selection.Text())
			return false
		}
		return true
	})
	return reporter
}

func (s *serviceScrapper) extractJudgementDate(row *goquery.Selection) string {
	var reporter string
	row.Find("td").EachWithBreak(func(_ int, selection *goquery.Selection) bool {
		strongText := strings.TrimSpace(selection.Find("strong").Text())
		if strings.Contains(strongText, "Data do julgamento:") {
			selection.Find("strong").Remove()
			reporter = strings.TrimSpace(selection.Text())
			return false
		}
		return true
	})
	return reporter
}

func (s *serviceScrapper) extractCaseClass(row *goquery.Selection) string {
	var reporter string
	row.Find("td").EachWithBreak(func(_ int, selection *goquery.Selection) bool {
		strongText := strings.TrimSpace(selection.Find("strong").Text())
		if strings.Contains(strongText, "Classe/Assunto:") {
			selection.Find("strong").Remove()
			reporter = strings.TrimSpace(selection.Text())
			return false
		}
		return true
	})
	return reporter
}

func (s *serviceScrapper) extractPublicationDate(row *goquery.Selection) string {
	var reporter string
	row.Find("td").EachWithBreak(func(_ int, selection *goquery.Selection) bool {
		strongText := strings.TrimSpace(selection.Find("strong").Text())
		if strings.Contains(strongText, "Data de publicação:") {
			selection.Find("strong").Remove()
			reporter = strings.TrimSpace(selection.Text())
			return false
		}
		return true
	})
	return reporter
}

func isCaseFound(currentCase entity.LegalCase) bool {
	requiredFields := []string{
		currentCase.CaseNumber, currentCase.Summary, currentCase.Reporter, currentCase.JudgingBody,
		currentCase.PublicationDate, currentCase.JudgementDate, currentCase.CaseClass,
	}

	for _, field := range requiredFields {
		if field == "" {
			return false
		}
	}
	return true
}

func (s *serviceScrapper) processCaseRow(r *goquery.Selection, currentCase entity.LegalCase, cases *[]entity.LegalCase) entity.LegalCase {
	caseNumber := s.extractCaseNumber(r)
	if caseNumber != "" {
		if isCaseFound(currentCase) {
			*cases = append(*cases, currentCase)
		}
		currentCase = entity.LegalCase{CaseNumber: caseNumber}
	}

	currentCase = s.updateFields(r, currentCase)
	return currentCase
}

func (s *serviceScrapper) updateFields(row *goquery.Selection, currentCase entity.LegalCase) entity.LegalCase {
	if summary := s.extractSummary(row); summary != "" {
		currentCase.Summary = summary
	}
	if reporter := s.extractReporter(row); reporter != "" {
		currentCase.Reporter = reporter
	}
	if court := s.extractCourt(row); court != "" {
		currentCase.Court = court
	}
	if judgingBody := s.extractJudgingBody(row); judgingBody != "" {
		currentCase.JudgingBody = judgingBody
	}
	if judgementDate := s.extractJudgementDate(row); judgementDate != "" {
		currentCase.JudgementDate = judgementDate
	}
	if caseClass := s.extractCaseClass(row); caseClass != "" {
		currentCase.CaseClass = caseClass
	}
	if publicationDate := s.extractPublicationDate(row); publicationDate != "" {
		currentCase.PublicationDate = publicationDate
	}
	return currentCase
}

func cleanText(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", " ")

	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}
