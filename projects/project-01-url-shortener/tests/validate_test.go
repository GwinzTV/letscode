package tests

import (
	"testing"

	"letscode/project-01-url-shortener/internal/handlers"
)

func TestValidateURL(t *testing.T) {
	cases := []struct{
		name string
		raw string
		wantErr bool
	}{
		{"empty", "", true},
		{"no-scheme", "example.com", true},
		{"http-scheme", "http://example.com", true},
		{"https-scheme", "https://example.com/path", false},
		{"no-host", "https:///path", true},
		{"invalid-url", "://bad", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := handlers.ValidateURL(c.raw)
			if (err != nil) != c.wantErr {
				t.Fatalf("ValidateURL(%q) error = %v, wantErr=%v", c.raw, err, c.wantErr)
			}
		})
	}
}
