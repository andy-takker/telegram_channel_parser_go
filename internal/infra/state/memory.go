package state

import "sync"

type Memory struct {
	mu   sync.RWMutex
	last int
}

func NewMemory() *Memory { return &Memory{} }

func (m *Memory) Get() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.last
}

func (m *Memory) Set(v int) {
	m.mu.Lock()
	m.last = v
	m.mu.Unlock()
}
