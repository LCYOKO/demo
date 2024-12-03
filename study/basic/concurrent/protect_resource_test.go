package concurrent

import "sync"

// safeResource 很棒，所有的期望对资源的操作都只能通过定义在上 safeResource 上的方法来进行
type safeResource struct {
	resource interface{}
	lock     sync.Mutex
}

func (s *safeResource) DoSomethingToResource() {
	s.lock.Lock()
	defer s.lock.Unlock()
}

// Registry 没有用锁，并不安全
type Registry struct {
	resources map[string]interface{}
}

func (r *Registry) Register(name string, resource interface{}) {
	r.resources[name] = resource
}

func (r *Registry) Get(name string) (interface{}, error) {
	return nil, nil
}
