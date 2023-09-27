package gomap

import (
	"testing"
)

func TestThreadSafePureMap(t *testing.T) {
	// Create a new threadSafePureMap instance
	m := NewThreadSafePureMap[int, string]()

	// Test Store and Load methods
	m.Store(1, "one")
	val, ok := m.Load(1)
	if !ok {
		t.Errorf("Load: Expected key 1 to exist, but it doesn't")
	}
	if val != "one" {
		t.Errorf("Load: Expected value 'one', but got '%s'", val)
	}

	// Test LoadAndDelete method
	val, ok = m.LoadAndDelete(1)
	if !ok {
		t.Errorf("LoadAndDelete: Expected key 1 to exist, but it doesn't")
	}
	if val != "one" {
		t.Errorf("LoadAndDelete: Expected value 'one', but got '%s'", val)
	}
	val, ok = m.Load(1)
	if ok {
		t.Errorf("LoadAndDelete: Expected key 1 to be deleted, but it still exists")
	}

	// Test Delete method
	m.Store(2, "two")
	m.Delete(2)
	_, ok = m.Load(2)
	if ok {
		t.Errorf("Delete: Expected key 2 to be deleted, but it still exists")
	}

	// Test Contain method
	m.Store(3, "three")
	if !m.Contain(3) {
		t.Errorf("Contain: Expected key 3 to exist, but it doesn't")
	}
	if m.Contain(4) {
		t.Errorf("Contain: Expected key 4 not to exist, but it does")
	}

	// Test Clear method
	m.Clear()
	_, ok = m.Load(3)
	if ok {
		t.Errorf("Clear: Expected map to be empty, but it still contains keys")
	}
}

func TestThreadSafePureMap_Empty(t *testing.T) {
	// Create an empty threadSafePureMap
	m := NewThreadSafePureMap[int, string]()

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
