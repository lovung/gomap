package gomap

import (
	"testing"
)

func TestSortedSliceMap(t *testing.T) {
	// Create a new sortedSliceMap instance
	m := NewSortedSliceMap[int, string]()

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
		t.Errorf("Load: Expected key 2 to be deleted, but it still exists")
	}

	// Test Delete method
	m.Store(4, "four")
	m.Delete(4)
	_, ok = m.Load(4)
	if ok {
		t.Errorf("Delete: Expected key 4 to be deleted, but it still exists")
	}

	// Test Contain method
	if !m.Contain(1) {
		t.Errorf("Contain: Expected key 1 to exist, but it doesn't")
	}
	if m.Contain(5) {
		t.Errorf("Contain: Expected key 5 not to exist, but it does")
	}

	// Test LoadAndDelete method
	val, ok = m.LoadAndDelete(5)
	if ok {
		t.Errorf("LoadAndDelete: Expected key 5 not to exist, but it doesn't")
	}

	// Test Clear method
	m.Clear()
	_, ok = m.Load(1)
	if ok {
		t.Errorf("Clear: Expected map to be empty, but it still contains keys")
	}

	_, ok = m.Load(3)
	if ok {
		t.Errorf("Clear: Expected map to be empty, but it still contains keys")
	}
}
