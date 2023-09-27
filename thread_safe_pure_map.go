package gomap

import "sync"

// Define the threadSafePureMap struct
type threadSafePureMap[K comparable, V any] struct {
	store map[K]V
	mu    sync.RWMutex // Mutex for thread-safety
}

// NewThreadSafePureMap creates a new threadSafePureMap instance
func NewThreadSafePureMap[K comparable, V any]() Map[K, V] {
	return &threadSafePureMap[K, V]{
		store: make(map[K]V),
	}
}

// Store implements the Store method of the Map interface
func (pm *threadSafePureMap[K, V]) Store(key K, val V) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.store[key] = val
}

// Load implements the Load method of the Map interface
func (pm *threadSafePureMap[K, V]) Load(key K) (V, bool) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	val, ok := pm.store[key]
	return val, ok
}

// LoadAndDelete implements the LoadAndDelete method of the Map interface
func (pm *threadSafePureMap[K, V]) LoadAndDelete(key K) (V, bool) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	val, ok := pm.store[key]
	if ok {
		delete(pm.store, key)
	}
	return val, ok
}

// Delete implements the Delete method of the Map interface
func (pm *threadSafePureMap[K, V]) Delete(key K) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	delete(pm.store, key)
}

// Contain implements the Contain method of the Map interface
func (pm *threadSafePureMap[K, V]) Contain(key K) bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	_, ok := pm.store[key]
	return ok
}

// Clear implements the Clear method of the Map interface
func (pm *threadSafePureMap[K, V]) Clear() {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.store = make(map[K]V)
}
