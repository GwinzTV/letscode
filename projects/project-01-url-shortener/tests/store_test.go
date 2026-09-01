package tests

import (
	"errors"
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

func TestFileStore_SaveOrGetCodeReusesExistingCode(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "urls_idempotency_test.json")
	defer os.Remove(tmp)
	fs := store.NewFileStore(tmp)

	code, err := fs.SaveOrGetCode("abc123", "https://example.com")
	if err != nil {
		t.Fatalf("save failed: %v", err)
	}
	if code != "abc123" {
		t.Fatalf("got code %q, want %q", code, "abc123")
	}

	code, err = fs.SaveOrGetCode("different", "https://example.com")
	if err != nil {
		t.Fatalf("reuse failed: %v", err)
	}
	if code != "abc123" {
		t.Fatalf("got code %q, want existing code %q", code, "abc123")
	}
}

func TestFileStore_SaveOrGetCodeRejectsCodeCollision(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "urls_collision_test.json")
	defer os.Remove(tmp)
	fs := store.NewFileStore(tmp)

	if _, err := fs.SaveOrGetCode("abc123", "https://example.com/one"); err != nil {
		t.Fatalf("initial save failed: %v", err)
	}
	if _, err := fs.SaveOrGetCode("abc123", "https://example.com/two"); !errors.Is(err, store.ErrCodeCollision) {
		t.Fatalf("error = %v, want code collision", err)
	}
}
