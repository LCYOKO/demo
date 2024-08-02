package factor

import (
	"demo/internal/sample/store"
	"demo/internal/sample/store/memorystore"
	"fmt"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]store.BookStore)
)

func init() {
	Register("memo", memorystore.NewStore())
}
func Register(name string, p store.BookStore) {
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
func New(providerName string) (store.BookStore, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", providerName)
	}
	return p, nil
}
