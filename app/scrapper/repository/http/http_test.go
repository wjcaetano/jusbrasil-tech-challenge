package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gocolly/colly/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewScraperRepository(t *testing.T) {
	repo := NewScraperRepository()
	t.Run("should return not nil a scraper repository", func(t *testing.T) {
		assert.NotNil(t, repo)
	})

	t.Run("should return not nil a collector", func(t *testing.T) {
		assert.NotNil(t, repo.collector)
	})
}

func TestScraperRepository_FetchPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`<html><body>Test Page Content</body></html>`))
	}))
	defer server.Close()

	t.Run("should return a valid content from site", func(t *testing.T) {
		expectedContent := "<html><body>Test Page Content</body></html>"

		scraper := NewScraperRepository()

		scraper.collector.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Host", server.URL)
		})

		content, err := scraper.FetchPage(server.URL)

		require.NoError(t, err)
		assert.Equal(t, expectedContent, content)
	})

	t.Run("should return an error when failing to visit url", func(t *testing.T) {
		invalidURL := "http://invalid-url"
		expectedError := "failed to scrape page: Get \"http://invalid-url\": dial tcp: lookup invalid-url on 127.0.0.53:53: server misbehaving"

		scraper := NewScraperRepository()
		content, err := scraper.FetchPage(invalidURL)

		require.Error(t, err)
		assert.Empty(t, content)
		assert.Equal(t, expectedError, err.Error())

	})

}

func Test_getUserAgent(t *testing.T) {
	t.Run("should return user agent", func(t *testing.T) {
		expectedUserAgent := "MockUserAgent/1.0"

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("User-Agent", expectedUserAgent)
		}))
		defer server.Close()

		originalURL := userAgentURL
		userAgentURL = server.URL
		defer func() { userAgentURL = originalURL }()

		result := getUserAgent()

		assert.Equal(t, expectedUserAgent, result)
	})
}
