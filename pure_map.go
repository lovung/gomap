package gomap

// Define the pureMap struct
type pureMap[K comparable, V any] struct {
	store map[K]V
}

// NewPureMap creates a new pureMap instance
func NewPureMap[K comparable, V any]() *pureMap[K, V] {
	return &pureMap[K, V]{
		store: make(map[K]V),
	}
}

// Store implements the Store method of the Map interface
func (pm *pureMap[K, V]) Store(key K, val V) {
	pm.store[key] = val
}

// Load implements the Load method of the Map interface
func (pm *pureMap[K, V]) Load(key K) (V, bool) {
	val, ok := pm.store[key]
	return val, ok
}

// LoadAndDelete implements the LoadAndDelete method of the Map interface
func (pm *pureMap[K, V]) LoadAndDelete(key K) (V, bool) {
	val, ok := pm.store[key]
	if ok {
		delete(pm.store, key)
	}
	return val, ok
}

// Delete implements the Delete method of the Map interface
func (pm *pureMap[K, V]) Delete(key K) {
	delete(pm.store, key)
}

// Contain implements the Contain method of the Map interface
func (pm *pureMap[K, V]) Contain(key K) bool {
	_, ok := pm.store[key]
	return ok
}

// Clear implements the Clear method of the Map interface
func (pm *pureMap[K, V]) Clear() {
	pm.store = make(map[K]V)
}
