package doublylinkedlist

import (
	"slices"
	"testing"
)

func TestListNew(t *testing.T) {
	l1 := New[int]()

	if l1.size != 0 {
		t.Errorf("Expected size to be 0, got %d", l1.size)
	}

	l2 := New[int](1, 2, 3)

	if l2.size != 3 {
		t.Errorf("Expected size to be 3, got %d", l2.size)
	}
}

func TestGet(t *testing.T) {
	l1 := New[int]()

	if _, ok := l1.Get(0); ok {
		t.Errorf("Expected element not be retrieved, but got %t", ok)
	}

	l2 := New[int](1, 2, 3)

	if v, _ := l2.Get(0); v != 1 {
		t.Errorf("Expected element to be 1, got %d", v)
	}

	if v, _ := l2.Get(1); v != 2 {
		t.Errorf("Expected element to be 2, got %d", v)
	}

	if v, _ := l2.Get(2); v != 3 {
		t.Errorf("Expected element to be 3, got %d", v)
	}
}

func TestPrepend(t *testing.T) {
	l := New[int]()

	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)

	if l.size != 3 {
		t.Errorf("Expected size to be 3, got %d", l.size)
	}

	if v, _ := l.Get(0); v != 3 {
		t.Errorf("Expected element to be 3, got %d", v)
	}

	if v, _ := l.Get(1); v != 2 {
		t.Errorf("Expected element to be 2, got %d", v)
	}

	if v, _ := l.Get(2); v != 1 {
		t.Errorf("Expected element to be 1, got %d", v)
	}
}

func TestAppend(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	if l.size != 3 {
		t.Errorf("Expected size to be 3, got %d", l.size)
	}

	if v, _ := l.Get(0); v != 1 {
		t.Errorf("Expected element to be 1, got %d", v)
	}

	if v, _ := l.Get(1); v != 2 {
		t.Errorf("Expected element to be 2, got %d", v)
	}
}

func TestAdd(t *testing.T) {
	l := New[int]()

	ok := l.Add(1, 1)

	if ok {
		t.Errorf("Expected element not be added, but got %t", ok)
	}

	l.Add(0, 1)
	l.Add(1, 2)
	l.Add(2, 3)
	l.Add(1, 4)

	if l.size != 4 {
		t.Errorf("Expected size to be 4, got %d", l.size)
	}

	if v, _ := l.Get(0); v != 1 {
		t.Errorf("Expected element to be 1, got %d", v)
	}

	if v, _ := l.Get(1); v != 4 {
		t.Errorf("Expected element to be 4, got %d", v)
	}

	if v, _ := l.Get(2); v != 2 {
		t.Errorf("Expected element to be 2, got %d", v)
	}

	if v, _ := l.Get(3); v != 3 {
		t.Errorf("Expected element to be 3, got %d", v)
	}
}

func TestDelete(t *testing.T) {
	l := New[int]()

	ok := l.Delete(0)

	if ok {
		t.Errorf("Expected element not be deleted, but got %t", ok)
	}

	ok = l.Delete(1)

	if ok {
		t.Errorf("Expected element not be deleted, but got %t", ok)
	}

	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	l.Delete(0)
	l.Delete(2)
	l.Delete(1)

	if l.size != 1 {
		t.Errorf("Expected size to be 1, got %d", l.size)
	}

	if v, _ := l.Get(0); v != 2 {
		t.Errorf("Expected element to be 2, got %d", v)
	}
}

func TestValues(t *testing.T) {
	l := New[int]()

	l.Append(1)
	l.Append(2)
	l.Append(3)

	actualValues := l.Values()
	expectedValues := []int{1, 2, 3}

	if !slices.Equal(actualValues, expectedValues) {
		t.Errorf("Got %v, expected %v", actualValues, expectedValues)
	}
}

func TestEmpty(t *testing.T) {
	l := New[int]()

	if !l.Empty() {
		t.Errorf("Expected empty list")
	}

	l.Append(1)

	if l.Empty() {
		t.Errorf("Expected non empty list")
	}
}

func TestSize(t *testing.T) {
	l := New[int]()

	if l.size != 0 {
		t.Errorf("Expected size to be 0, got %d", l.size)
	}

	l.Append(1)
	l.Append(2)

	if l.size != 2 {
		t.Errorf("Expected size to be 2, got %d", l.size)
	}
}
