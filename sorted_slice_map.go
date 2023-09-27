package gomap

import (
	"golang.org/x/exp/constraints"
)

type sliceItem[K constraints.Ordered, V any] struct {
	k K
	v V
}

type sortedSliceMap[K constraints.Ordered, V any] struct {
	store []sliceItem[K, V]
}

// NewSortedSliceMap create sorted slice map,
// using binary search to find the item
// non-thread-safe
func NewSortedSliceMap[K constraints.Ordered, V any](opts ...Option) Map[K, V] {
	opt := option{}
	for _, o := range opts {
		o(&opt)
	}

	m := &sortedSliceMap[K, V]{
		store: make([]sliceItem[K, V], 0),
	}
	return m
}

// return left_index and exist
// E.g.
// Current data: []int{1, 3},
// Input: 0 -> Output: 0, false
// Input: 1 -> Output: 0, true
// Input: 2 -> Output: 1, false
// Input: 3 -> Output: 1, true
// Input: 4 -> Output: 2, false
func (m *sortedSliceMap[K, V]) binarySearch(key K) (int, bool) {
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

func (m *sortedSliceMap[K, V]) Store(key K, val V) {
	idx, exist := m.binarySearch(key)

	if !exist {
		// Key doesn't exist, insert it at the correct position.
		m.store = append(m.store, sliceItem[K, V]{key, val})
		copy(m.store[idx+1:], m.store[idx:len(m.store)-1])
	}
	m.store[idx].k = key
	m.store[idx].v = val
}

func (m *sortedSliceMap[K, V]) Load(key K) (V, bool) {
	var zero V
	idx, exist := m.binarySearch(key)
	if !exist {
		return zero, exist
	}
	return m.store[idx].v, exist
}

func (m *sortedSliceMap[K, V]) LoadAndDelete(key K) (V, bool) {
	v, exist := m.Load(key)
	if !exist {
		return v, exist
	}
	m.Delete(key)
	return v, exist
}

func (m *sortedSliceMap[K, V]) Delete(key K) {
	idx, found := m.binarySearch(key)
	if found {
		// Remove the key-value pair at the found index by slicing the store.
		m.store = m.store[:idx+copy(m.store[idx:], m.store[idx+1:])]
	}
}

func (m *sortedSliceMap[K, V]) Contain(key K) bool {
	_, exist := m.binarySearch(key)
	return exist
}

func (m *sortedSliceMap[K, V]) Clear() {
	m.store = make([]sliceItem[K, V], 0)
}
