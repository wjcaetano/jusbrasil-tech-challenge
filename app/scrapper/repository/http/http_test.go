package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	tableHTML string = `
			<html>
				<body>
					<div id="divDadosResultado-A">
						<table>
							<tbody>
								<tr>
									<td>Test Row 1</td>
								</tr>
								<tr>
									<td>Test Row 2</td>
								</tr>
							</tbody>
						</table>
					</div>
				</body>
			</html>
		`
)

func TestNewScrapperRepository(t *testing.T) {
	repo := NewScrapperRepository()
	t.Run("should return not nil a scraper repository", func(t *testing.T) {
		assert.NotNil(t, repo)
	})

	t.Run("should return not nil a collector", func(t *testing.T) {
		assert.NotNil(t, repo.collector)
	})
}

func TestScrapperRepository_FetchPage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(tableHTML))
	}))
	defer server.Close()

	t.Run("should return a valid content from site", func(t *testing.T) {
		expectedContent := `
								<tr>
									<td>Test Row 1</td>
								</tr>
								<tr>
									<td>Test Row 2</td>
								</tr>
							`

		scraper := NewScrapperRepository()

		content, err := scraper.FetchPage(server.URL)

		require.NoError(t, err)
		assert.Equal(t, expectedContent, content)
	})

	t.Run("should return an error when failing to visit url", func(t *testing.T) {
		invalidURL := "http://invalid-url"
		expectedError := "failed to fetch page: Get \"http://invalid-url\": dial tcp: lookup invalid-url on 127.0.0.53:53: server misbehaving"

		scraper := NewScrapperRepository()
		content, err := scraper.FetchPage(invalidURL)

		require.Error(t, err)
		assert.Empty(t, content)
		assert.Equal(t, expectedError, err.Error())

	})
}
