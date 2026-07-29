package tests

import (
	"os"
	"path/filepath"
	"testing"

	"letscode/project-01-url-shortener/internal/store"
)

func TestFileStore_SaveLookup(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "urls_test.json")
	defer os.Remove(tmp)
	fs := store.NewFileStore(tmp)
	if err := fs.Save("abc123", "https://example.com"); err != nil {
		t.Fatalf("save failed: %v", err)
	}
	if u, ok := fs.Lookup("abc123"); !ok || u != "https://example.com" {
		t.Fatalf("lookup failed: got %v, ok=%v", u, ok)
	}
}
