package gomap

import (
	"sync"
)

type syncMap[K comparable, V any] struct {
	store sync.Map
}

func NewSyncMap[K comparable, V any](opts ...Option) Map[K, V] {
	opt := option{}
	for _, o := range opts {
		o(&opt)
	}

	m := &syncMap[K, V]{
		store: sync.Map{},
	}
	return m
}

func (m *syncMap[K, V]) Store(key K, val V) {
	m.store.Store(key, val)
}

func (m *syncMap[K, V]) Load(key K) (V, bool) {
	var zero V
	if val, ok := m.store.Load(key); ok {
		v, ok := val.(V)
		return v, ok
	}
	return zero, false
}

func (m *syncMap[K, V]) LoadAndDelete(key K) (V, bool) {
	var zero V
	if val, ok := m.store.LoadAndDelete(key); ok {
		v, ok := val.(V)
		return v, ok
	}
	return zero, false
}

func (m *syncMap[K, V]) Delete(key K) {
	_, _ = m.store.LoadAndDelete(key)
}

func (m *syncMap[K, V]) Contain(key K) bool {
	_, ok := m.store.Load(key)
	return ok
}

func (m *syncMap[K, V]) Clear() {
	m.store.Range(func(key, value any) bool {
		m.store.Delete(key)
		return true
	})
}
