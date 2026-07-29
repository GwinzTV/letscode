package store

import (
	"encoding/json"
	"os"
	"sync"

	"letscode/project-01-url-shortener/internal/model"
)

// Store is a minimal persistence interface for URL mappings.
type Store interface {
	Save(code, url string) error
	Lookup(code string) (string, bool)
	All() ([]model.URLMapping, error)
}

// FileStore implements Store using a single JSON file.
type FileStore struct {
	path string
	mu   sync.RWMutex
	m    map[string]string
}

// NewFileStore creates or loads a file-backed store.
func NewFileStore(path string) *FileStore {
	fs := &FileStore{path: path, m: make(map[string]string)}
	fs.load()
	return fs
}

func (f *FileStore) load() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.m = make(map[string]string)
	file, err := os.Open(f.path)
	if err != nil {
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&f.m)
}

// persistSnapshot writes the provided snapshot to disk without acquiring
// any locks. Callers must ensure they have created a snapshot while holding
// the appropriate lock (or otherwise ensure snapshot safety).
func (f *FileStore) persistSnapshot(snapshot map[string]string) error {
	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(snapshot)
}

func (f *FileStore) Save(code, url string) error {
	// Update map under write lock and create a snapshot while still holding
	// the lock. Release the lock before doing file I/O to avoid deadlocks and
	// keeping I/O off the critical path.
	f.mu.Lock()
	f.m[code] = url
	snapshot := make(map[string]string, len(f.m))
	for k, v := range f.m {
		snapshot[k] = v
	}
	f.mu.Unlock()

	return f.persistSnapshot(snapshot)
}

func (f *FileStore) Lookup(code string) (string, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	u, ok := f.m[code]
	return u, ok
}

func (f *FileStore) All() ([]model.URLMapping, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	out := make([]model.URLMapping, 0, len(f.m))
	for k, v := range f.m {
		out = append(out, model.URLMapping{Code: k, URL: v})
	}
	return out, nil
}
