package Mutexes

import "sync"

type mutex struct {
	m sync.Map
}

var ProdMutex mutex

func (m *mutex) Lock(key int) bool {
	_, ok := m.m.Load(key)
	if ok {
		return false
	}

	m.m.Store(key, true)
	return !ok
}

func (m *mutex) UnLock(id int) {
	m.m.Delete(id)
}
