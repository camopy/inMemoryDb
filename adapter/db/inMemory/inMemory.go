package inmemory

import (
	"errors"
	"sync"
)

type InMemory struct {
	mu       sync.Mutex
	messages map[uint64]string
}

func New() *InMemory {
	return &InMemory{
		messages: make(map[uint64]string),
	}
}

func (im *InMemory) Get(id uint64) (string, error) {
	im.mu.Lock()
	defer im.mu.Unlock()

	if msg, ok := im.messages[id]; ok {
		return msg, nil
	}
	return "", errors.New("not found")
}

func (im *InMemory) Set(id uint64, message string) {
	im.mu.Lock()
	defer im.mu.Unlock()

	im.messages[id] = message
}
