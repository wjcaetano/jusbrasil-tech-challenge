package scrapperprocess

import (
	"errors"
	entity "jusbrasil-tech-challenge/app/scrapper"
	"jusbrasil-tech-challenge/app/scrapper/service/scrapper/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessScrapper_FetchAndParseCases(t *testing.T) {

	t.Run("should fetch and parse cases successfully", func(t *testing.T) {
		mockService := mocks.ScrapperService{}
		scrapper := NewProcessScrapper(&mockService)
		expectedCases := []entity.LegalCase{
			{
				CaseNumber:      "12345-67.2023.8.26.0001",
				Summary:         "Test Summary for Legal Case.",
				Reporter:        "John Doe",
				Court:           "São Paulo",
				JudgingBody:     "7ª Câmara de Direito Criminal",
				JudgementDate:   "25/01/2023",
				CaseClass:       "Habeas Corpus / Test Case",
				PublicationDate: "26/01/2023",
			},
		}

		mockService.EXPECT().GetLegalCases("http://test-url.com").Return(expectedCases, nil)

		cases, err := scrapper.FetchAndParseCases("http://test-url.com")
		require.NoError(t, err)
		assert.Equal(t, expectedCases, cases)
		mockService.AssertExpectations(t)
	})

	t.Run("should return error when error to fetch cases", func(t *testing.T) {
		mockService := mocks.ScrapperService{}
		scrapper := NewProcessScrapper(&mockService)
		expectedError := errors.New("something went wrong")

		mockService.EXPECT().GetLegalCases("http://test-url.com").Return(nil, expectedError)
		cases, err := scrapper.FetchAndParseCases("http://test-url.com")
		require.Error(t, err)
		assert.Nil(t, cases)
	})
}
