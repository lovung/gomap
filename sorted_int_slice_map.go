package gomap

import (
	bf "github.com/lovung/bloomfilter"
	"golang.org/x/exp/constraints"
)

type intSliceItem[K constraints.Integer, V any] struct {
	k K
	v V
}

type intSortedSliceMap[K constraints.Integer, V any] struct {
	store []intSliceItem[K, V]

	bloomFilter bf.BloomFilter[K]
}

// NewIntSortedSliceMap create sorted slice map,
// using binary search to find the item
// using bloom filter to predict if the item doesn't exist
// non-thread-safe
func NewIntSortedSliceMap[K constraints.Integer, V any](opts ...Option) Map[K, V] {
	opt := option{}
	for _, o := range opts {
		o(&opt)
	}

	m := &intSortedSliceMap[K, V]{
		store:       make([]intSliceItem[K, V], 0),
		bloomFilter: bf.BloomFilter[K](0),
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
func (m *intSortedSliceMap[K, V]) binarySearch(key K) (int, bool) {
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

func (m *intSortedSliceMap[K, V]) Store(key K, val V) {
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

func (m *intSortedSliceMap[K, V]) Load(key K) (V, bool) {
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

func (m *intSortedSliceMap[K, V]) LoadAndDelete(key K) (V, bool) {
	var zero V
	v, exist := m.Load(key)
	if !exist {
		return zero, exist
	}
	m.Delete(key)
	return v, exist
}

func (m *intSortedSliceMap[K, V]) Delete(key K) {
	idx, found := m.binarySearch(key)
	if found {
		// Remove the key-value pair at the found index by slicing the store.
		m.store = m.store[:idx+copy(m.store[idx:], m.store[idx+1:])]
	}
}

func (m *intSortedSliceMap[K, V]) Contain(key K) bool {
	if !m.bloomFilter.MayExist(key) {
		return false
	}
	_, exist := m.binarySearch(key)
	return exist
}

func (m *intSortedSliceMap[K, V]) Clear() {
	m.store = make([]intSliceItem[K, V], 0)
}
