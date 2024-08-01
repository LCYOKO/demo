package store

import (
	"fmt"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]BookStore)
)

func Register(name string, p BookStore) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if p == nil {
		panic("store: Register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider " + name)
	}
	providers[name] = p
}
func New(providerName string) (BookStore, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", providerName)
	}
	return p, nil
}
