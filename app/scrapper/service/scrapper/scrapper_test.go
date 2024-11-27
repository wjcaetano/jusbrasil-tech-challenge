package scrapper

import (
	"errors"
	"jusbrasil-tech-challenge/app/scrapper"
	"jusbrasil-tech-challenge/app/scrapper/service/scrapper/mocks"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

const (
	testCaseHTMLPath = "./testdata/test_case.html"
)

func TestScrapperService_GetLegalCases(t *testing.T) {
	t.Run("should return valid cases from valid HTML", func(t *testing.T) {
		mockRepo := mocks.ScrapperRepository{}
		htmlContent := readHTMLFromFile(t, testCaseHTMLPath)
		expectedCase := scrapper.LegalCase{
			CaseNumber:      "1001143-04.2021.8.26.0541",
			Summary:         "Sample summary",
			Reporter:        "John Doe",
			Court:           "São Paulo",
			JudgingBody:     "7ª Câmara de Direito Criminal",
			JudgementDate:   "25/01/2022",
			CaseClass:       "Apelação Cível / Bancários",
			PublicationDate: "25/01/2022",
		}
		mockRepo.EXPECT().FetchPage("http://test-url.com").Return(htmlContent, nil)

		service := NewScrapperService(&mockRepo)

		cases, err := service.GetLegalCases("http://test-url.com")
		require.NoError(t, err)
		assertCases(t, cases, expectedCase)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when occurs error to fetch page", func(t *testing.T) {
		mockRepo := mocks.ScrapperRepository{}
		service := NewScrapperService(&mockRepo)
		expectedError := errors.New("something went wrong")
		mockRepo.EXPECT().FetchPage("http://test-url.com").Return("", expectedError)

		_, err := service.GetLegalCases("http://test-url.com")
		require.Error(t, err)
		mockRepo.AssertExpectations(t)

	})
}

func TestScrapperService_ProcessCaseRow(t *testing.T) {
	t.Run("should append current case to list when all fields are complete and new case starts", func(t *testing.T) {
		service := &serviceScrapper{}

		currentCase := scrapper.LegalCase{
			CaseNumber:      "1001143-04.2021.8.26.0541",
			Summary:         "Existing summary",
			Reporter:        "John Doe",
			Court:           "São Paulo",
			JudgingBody:     "7ª Câmara de Direito Criminal",
			JudgementDate:   "25/01/2022",
			CaseClass:       "Apelação Cível / Bancários",
			PublicationDate: "25/01/2022",
		}
		expectedCaseNumber := "2001234-56.2023.8.26.0001"

		var cases []scrapper.LegalCase

		html := `
<table>
	<tr>
		<td><a>2001234-56.2023.8.26.0001</a></td>
	</tr>
</table>
`
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		require.NoError(t, err)

		row := doc.Find("tr")

		resultCase := service.processCaseRow(row, currentCase, &cases)

		assert.Len(t, cases, 1)
		assert.Equal(t, currentCase.CaseNumber, cases[0].CaseNumber)
		assert.Equal(t, expectedCaseNumber, resultCase.CaseNumber)
	})
}

func readHTMLFromFile(t *testing.T, fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	return string(content)
}

func assertCases(t *testing.T, cases []scrapper.LegalCase, expectedCases scrapper.LegalCase) {
	assert.Len(t, cases, 1)
	assert.Equal(t, expectedCases.CaseNumber, cases[0].CaseNumber)
	assert.Equal(t, expectedCases.Summary, cases[0].Summary)
	assert.Equal(t, expectedCases.Reporter, cases[0].Reporter)
	assert.Equal(t, expectedCases.Court, cases[0].Court)
	assert.Equal(t, expectedCases.JudgingBody, cases[0].JudgingBody)
	assert.Equal(t, expectedCases.JudgementDate, cases[0].JudgementDate)
	assert.Equal(t, expectedCases.CaseClass, cases[0].CaseClass)
	assert.Equal(t, expectedCases.PublicationDate, cases[0].PublicationDate)
}
