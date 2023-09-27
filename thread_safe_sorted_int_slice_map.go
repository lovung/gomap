package gomap

import (
	"sync"

	bf "github.com/lovung/bloomfilter"
	"golang.org/x/exp/constraints"
)

// Define the threadSafeIntSortedSliceMap struct
type threadSafeIntSortedSliceMap[K constraints.Integer, V any] struct {
	store       []intSliceItem[K, V]
	bloomFilter bf.BloomFilter[K]
	mu          sync.RWMutex // Mutex for thread-safety
}

// NewThreadSafeIntSortedSliceMap creates a new threadSafeIntSortedSliceMap instance
func NewThreadSafeIntSortedSliceMap[K constraints.Integer, V any](opts ...Option) Map[K, V] {
	opt := option{}
	for _, o := range opts {
		o(&opt)
	}

	m := &threadSafeIntSortedSliceMap[K, V]{
		store:       make([]intSliceItem[K, V], 0),
		bloomFilter: bf.BloomFilter[K](0),
	}
	return m
}

// binarySearch finds the index of the given key using binary search
func (m *threadSafeIntSortedSliceMap[K, V]) binarySearch(key K) (int, bool) {
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

func (m *threadSafeIntSortedSliceMap[K, V]) Store(key K, val V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	idx, exist := m.binarySearch(key)

	if !exist {
		// Key doesn't exist, insert it at the correct position.
		m.store = append(m.store, intSliceItem[K, V]{key, val})
		copy(m.store[idx+1:], m.store[idx:len(m.store)-1])
	}
	m.store[idx].k = key
	m.store[idx].v = val
	m.bloomFilter.Add(key)
}

func (m *threadSafeIntSortedSliceMap[K, V]) Load(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var zero V
	if !m.bloomFilter.MayExist(key) {
		return zero, false
	}
	idx, exist := m.binarySearch(key)
	if !exist {
		return zero, exist
	}
	return m.store[idx].v, exist
}

func (m *threadSafeIntSortedSliceMap[K, V]) LoadAndDelete(key K) (V, bool) {
	var zero V

	v, exist := m.Load(key)
	if !exist {
		return zero, exist
	}
	m.Delete(key)
	return v, exist
}

func (m *threadSafeIntSortedSliceMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	idx, found := m.binarySearch(key)
	if found {
		// Remove the key-value pair at the found index by slicing the store.
		m.store = m.store[:idx+copy(m.store[idx:], m.store[idx+1:])]
	}
}

func (m *threadSafeIntSortedSliceMap[K, V]) Contain(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.bloomFilter.MayExist(key) {
		return false
	}
	_, exist := m.binarySearch(key)
	return exist
}

func (m *threadSafeIntSortedSliceMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store = make([]intSliceItem[K, V], 0)
}
