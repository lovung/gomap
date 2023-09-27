package gomap

import (
	"testing"
)

func TestPureMap(t *testing.T) {
	// Create a new pureMap instance
	pm := NewPureMap[int, string]()

	// Test Store and Load methods
	pm.Store(1, "one")
	val, ok := pm.Load(1)
	if !ok {
		t.Errorf("Load: Expected key 1 to exist, but it doesn't")
	}
	if val != "one" {
		t.Errorf("Load: Expected value 'one', but got '%s'", val)
	}

	// Test LoadAndDelete method
	val, ok = pm.LoadAndDelete(1)
	if !ok {
		t.Errorf("LoadAndDelete: Expected key 1 to exist, but it doesn't")
	}
	if val != "one" {
		t.Errorf("LoadAndDelete: Expected value 'one', but got '%s'", val)
	}
	val, ok = pm.Load(1)
	if ok {
		t.Errorf("LoadAndDelete: Expected key 1 to be deleted, but it still exists")
	}

	// Test Delete method
	pm.Store(2, "two")
	pm.Delete(2)
	_, ok = pm.Load(2)
	if ok {
		t.Errorf("Delete: Expected key 2 to be deleted, but it still exists")
	}

	// Test Contain method
	pm.Store(3, "three")
	if !pm.Contain(3) {
		t.Errorf("Contain: Expected key 3 to exist, but it doesn't")
	}
	if pm.Contain(4) {
		t.Errorf("Contain: Expected key 4 not to exist, but it does")
	}

	// Test Clear method
	pm.Clear()
	_, ok = pm.Load(3)
	if ok {
		t.Errorf("Clear: Expected map to be empty, but it still contains keys")
	}
}

func TestPureMap_Empty(t *testing.T) {
	// Create an empty pureMap
	pm := NewPureMap[int, string]()

	// Test Load method for non-existent key
	_, ok := pm.Load(1)
	if ok {
		t.Errorf("Load: Expected key 1 not to exist in an empty map, but it does")
	}

	// Test LoadAndDelete method for non-existent key
	_, ok = pm.LoadAndDelete(2)
	if ok {
		t.Errorf("LoadAndDelete: Expected key 2 not to exist in an empty map, but it does")
	}

	// Test Delete method for non-existent key
	pm.Delete(3) // Deleting a non-existent key should not cause an error
}
