package minutil

import (
	"log"
	"sync"
)

// MapTokenStore 是基于内存的 Token 存储实现
type MapTokenStore struct {
	mu    sync.RWMutex
	store map[string]bool
}

// NewMapTokenStore 创建一个新的 MapTokenStore
func NewMapTokenStore() *MapTokenStore {
	log.Println("Creating new MapTokenStore")
	return &MapTokenStore{
		store: make(map[string]bool),
	}
}

// Get 获取 Token 的状态
func (m *MapTokenStore) Get(token string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	log.Printf("Getting token status for token: %s", token)
	val, ok := m.store[token]
	if !ok {
		log.Printf("Token %s not found", token)
		return false, nil
	}
	log.Printf("Token %s found, status: %v", token, val)
	return val, nil
}

// Set 设置 Token 的状态
func (m *MapTokenStore) Set(token string, invalid bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	log.Printf("Setting token status for token: %s, invalid: %v", token, invalid)
	m.store[token] = invalid
	log.Printf("Token %s status set to %v", token, invalid)
	return nil
}