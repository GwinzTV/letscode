package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"letscode/project-01-url-shortener/internal/handlers"
	"letscode/project-01-url-shortener/internal/store"
)

func TestShortenReusesExistingCode(t *testing.T) {
	st := store.NewFileStore(t.TempDir() + "/urls.json")
	h := handlers.NewHandler(st)
	routes := h.Routes()

	request := func() string {
		req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(`{"url":"https://example.com"}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		routes.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
		}
		var response struct {
			Code string `json:"code"`
		}
		if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		return response.Code
	}

	firstCode := request()
	if secondCode := request(); secondCode != firstCode {
		t.Fatalf("second code = %q, want existing code %q", secondCode, firstCode)
	}
}
