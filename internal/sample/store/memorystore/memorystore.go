package memorystore

import "demo/internal/sample/store"

func init() {
	store.Register("mem", newStore())
}

type MemoryStore struct {
	table map[int64]*store.Book
}

func (m *MemoryStore) Save(book *store.Book) {
	if nil == book {
		return
	}
	m.table[book.Id] = book
}

func (m *MemoryStore) GetById(id int64) *store.Book {
	return m.table[id]
}

func newStore() *MemoryStore {
	return &MemoryStore{
		table: make(map[int64]*store.Book, 64),
	}
}
