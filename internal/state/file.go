package state

import (
	"encoding/json"
	"os"
	"sync"
)

type FileStore struct {
	path string
	mu   sync.RWMutex
	val  int
}

func NewFileStore(path string) (*FileStore, error) {
	s := &FileStore{path: path}
	_ = s.load()
	return s, nil
}

func (s *FileStore) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	f, err := os.Open(s.path)
	if err != nil {
		return nil // файл может отсутствовать
	}
	defer f.Close()
	var v struct {
		Last int `json:"last_seen_id"`
	}
	if err := json.NewDecoder(f).Decode(&v); err == nil {
		s.val = v.Last
	}
	return nil
}

func (s *FileStore) Get() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.val
}

func (s *FileStore) Set(v int) {
	s.mu.Lock()
	s.val = v
	defer s.mu.Unlock()
	_ = s.persist()
}

func (s *FileStore) persist() error {
	tmp := s.path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(struct {
		Last int `json:"last_seen_id"`
	}{Last: s.val}); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}
