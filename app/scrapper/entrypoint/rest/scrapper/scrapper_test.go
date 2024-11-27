package scrapper

import (
	"encoding/json"
	"errors"
	"jusbrasil-tech-challenge/app/scrapper"
	"jusbrasil-tech-challenge/app/scrapper/usecase/scrapperprocess/mocks"
	"jusbrasil-tech-challenge/tests/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testURL              = "https://test.com"
	testScrapperRouteURI = "/scrapper?url=" + testURL
)

func TestFetchScrapperHandler(t *testing.T) {
	mockUsecase := new(mocks.ProcessScrapper)

	handler := NewHandler(mockUsecase)

	router := chi.NewRouter()
	handler.RegisterScrapperRoutes(router)

	t.Run("should return bad request when URL parameter is missing", func(t *testing.T) {
		expectedError := "URL parameter is required\n"
		req := httptest.NewRequest(http.MethodGet, "/scrapper", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Equal(t, expectedError, resp.Body.String())
	})

	t.Run("should return internal server error on usecase failure", func(t *testing.T) {
		expectedError := "Failed to fetch and parse cases\n"
		mockUsecase.EXPECT().FetchAndParseCases(testURL).Return(nil, errors.New(expectedError)).Once()

		req := httptest.NewRequest(http.MethodGet, testScrapperRouteURI, nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Equal(t, expectedError, resp.Body.String())
		mockUsecase.AssertExpectations(t)
	})

	t.Run("should return status code 200 with parsed cases", func(t *testing.T) {
		mockCases := []scrapper.LegalCase{
			{
				CaseNumber: mock.RandomString(), Summary: mock.RandomString(),
			},
			{
				CaseNumber: mock.RandomString(), Summary: mock.RandomString(),
			},
		}

		mockUsecase.EXPECT().FetchAndParseCases(testURL).Return(mockCases, nil).Once()

		req := httptest.NewRequest(http.MethodGet, testScrapperRouteURI, nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)

		var responseCases []scrapper.LegalCase
		err := json.NewDecoder(resp.Body).Decode(&responseCases)
		require.NoError(t, err)
		assert.Equal(t, mockCases, responseCases)
		mockUsecase.AssertExpectations(t)
	})
}
