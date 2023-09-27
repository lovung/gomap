package gomap

import (
	"testing"
)

func TestThreadSafeIntSortedSliceMap(t *testing.T) {
	// Create a new threadSafeIntSortedSliceMap instance
	m := NewThreadSafeIntSortedSliceMap[int, string]()

	// Test Store and Load methods
	m.Store(3, "three")
	m.Store(1, "one")
	m.Store(2, "two")

	val, ok := m.Load(1)
	if !ok {
		t.Errorf("Load: Expected key 1 to exist, but it doesn't")
	}
	if val != "one" {
		t.Errorf("Load: Expected value 'one', but got '%s'", val)
	}

	val, ok = m.Load(2)
	if !ok {
		t.Errorf("Load: Expected key 2 to exist, but it doesn't")
	}
	if val != "two" {
		t.Errorf("Load: Expected value 'two', but got '%s'", val)
	}

	val, ok = m.Load(3)
	if !ok {
		t.Errorf("Load: Expected key 3 to exist, but it doesn't")
	}
	if val != "three" {
		t.Errorf("Load: Expected value 'three', but got '%s'", val)
	}

	// Test LoadAndDelete method
	val, ok = m.LoadAndDelete(2)
	if !ok {
		t.Errorf("LoadAndDelete: Expected key 2 to exist, but it doesn't")
	}
	if val != "two" {
		t.Errorf("LoadAndDelete: Expected value 'two', but got '%s'", val)
	}

	// Key 2 should have been deleted
	val, ok = m.Load(2)
	if ok {
		t.Errorf("LoadAndDelete: Expected key 2 to be deleted, but it still exists")
	}

	// Test Delete method
	m.Store(4, "four")
	m.Delete(4)
	_, ok = m.Load(4)
	if ok {
		t.Errorf("Delete: Expected key 4 to be deleted, but it still exists")
	}

	// Test Contain method
	m.Store(5, "five")
	if !m.Contain(5) {
		t.Errorf("Contain: Expected key 5 to exist, but it doesn't")
	}
	if m.Contain(6) {
		t.Errorf("Contain: Expected key 6 not to exist, but it does")
	}
	if m.Contain(1024) {
		t.Errorf("Contain: Expected key 1024 not to exist, but it does")
	}
	m.Store(1024, "1024")
	// Test LoadAndDelete method
	val, ok = m.LoadAndDelete(1024)
	if !ok {
		t.Errorf("LoadAndDelete: Expected key 1024 to exist, but it doesn't")
	}
	if val != "1024" {
		t.Errorf("LoadAndDelete: Expected value '1024', but got '%s'", val)
	}

	// Test Clear method
	m.Clear()
	_, ok = m.Load(3)
	if ok {
		t.Errorf("Clear: Expected map to be empty, but it still contains keys")
	}
}

func TestThreadSafeIntSortedSliceMap_Empty(t *testing.T) {
	// Create an empty threadSafeIntSortedSliceMap
	m := NewThreadSafeIntSortedSliceMap[int, string]()

	// Test Load method for non-existent key
	_, ok := m.Load(1)
	if ok {
		t.Errorf("Load: Expected key 1 not to exist in an empty map, but it does")
	}

	// Test LoadAndDelete method for non-existent key
	_, ok = m.LoadAndDelete(2)
	if ok {
		t.Errorf("LoadAndDelete: Expected key 2 not to exist in an empty map, but it does")
	}

	// Test Delete method for non-existent key
	m.Delete(3) // Deleting a non-existent key should not cause an error
}
