package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

type connection struct {
}

func (c *connection) send() error {
	fmt.Println("send...")
	return nil
}

func (c *connection) Close() error {
	fmt.Println("closing...")
	return nil
}

type SafeMap struct {
	m     map[string]interface{}
	mutex sync.RWMutex
}

// LoadOrStore loaded 代表是返回老的对象，还是返回了新的对象
func (s *SafeMap) LoadOrStore(key string, newVale interface{}) (val interface{}, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}
	s.m[key] = newVale
	return newVale, false
}

type valProvider func() interface{}

func (s *SafeMap) LoadOrStoreHeavy(key string, p valProvider) (val interface{}, loaded bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}
	newVale := p()
	s.m[key] = newVale
	return newVale, false
}

func TestSafeMap_LoadOrStore(t *testing.T) {
	m := &SafeMap{
		m: map[string]interface{}{},
	}

	for i := 0; i < 10; i++ {
		go func() {
			con := &connection{}
			nc, loaded := m.LoadOrStore("hello", con)
			if loaded {
				_ = con.Close()
			}
			_ = nc.(*connection).send()
		}()
	}
}

func TestSafeMap_LoadOrStoreHeavy(t *testing.T) {
	m := &SafeMap{
		m: map[string]interface{}{},
	}

	for i := 0; i < 10; i++ {
		go func() {
			nc, _ := m.LoadOrStoreHeavy("hello", func() interface{} {
				return &connection{}
			})
			_ = nc.(*connection).send()
		}()
	}
}
