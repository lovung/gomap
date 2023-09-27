package gomap

import (
	"sync"

	"golang.org/x/exp/constraints"
)

// Define the threadSafeSortedSliceMap struct
type threadSafeSortedSliceMap[K constraints.Ordered, V any] struct {
	store []sliceItem[K, V]
	mu    sync.RWMutex // Mutex for thread-safety
}

// NewThreadSafeSortedSliceMap creates a new threadSafeSortedSliceMap instance
// If your key is Integer, please consider to use IntSortedSliceMap to have bloom filter feature
func NewThreadSafeSortedSliceMap[K constraints.Ordered, V any](opts ...Option) Map[K, V] {
	opt := option{}
	for _, o := range opts {
		o(&opt)
	}

	m := &threadSafeSortedSliceMap[K, V]{
		store: make([]sliceItem[K, V], 0),
	}
	return m
}

// binarySearch finds the index of the given key using binary search
func (m *threadSafeSortedSliceMap[K, V]) binarySearch(key K) (int, bool) {
	left := 0
	right := len(m.store) - 1

	for left <= right {
		mid := left + (right-left)/2

		if m.store[mid].k == key {
			return mid, true
		} else if m.store[mid].k < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// If the key is not found, return the index where it should be inserted and false.
	return left, false
}

func (m *threadSafeSortedSliceMap[K, V]) Store(key K, val V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	idx, exist := m.binarySearch(key)

	if !exist {
		// Key doesn't exist, insert it at the correct position.
		m.store = append(m.store, sliceItem[K, V]{key, val})
		copy(m.store[idx+1:], m.store[idx:len(m.store)-1])
	}
	m.store[idx].k = key
	m.store[idx].v = val
}

func (m *threadSafeSortedSliceMap[K, V]) Load(key K) (V, bool) {
	var zero V
	m.mu.RLock()
	defer m.mu.RUnlock()

	idx, exist := m.binarySearch(key)
	if !exist {
		return zero, exist
	}
	return m.store[idx].v, exist
}

func (m *threadSafeSortedSliceMap[K, V]) LoadAndDelete(key K) (V, bool) {
	v, exist := m.Load(key)
	if !exist {
		return v, exist
	}
	m.Delete(key)
	return v, exist
}

func (m *threadSafeSortedSliceMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	idx, found := m.binarySearch(key)
	if found {
		// Remove the key-value pair at the found index by slicing the store.
		m.store = m.store[:idx+copy(m.store[idx:], m.store[idx+1:])]
	}
}

func (m *threadSafeSortedSliceMap[K, V]) Contain(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, exist := m.binarySearch(key)
	return exist
}

func (m *threadSafeSortedSliceMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store = make([]sliceItem[K, V], 0)
}
