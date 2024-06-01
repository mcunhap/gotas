package arraystack

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int](3)

	if s.capacity != 3 {
		t.Errorf("Expected capacity to be 3, got %d", s.capacity)
	}

	if s.top != -1 {
		t.Errorf("Expected top to be -1, got %d", s.top)
	}

	if len(s.data) != 0 {
		t.Errorf("Expected length of data to be 0, got %d", len(s.data))
	}
}

func TestPush(t *testing.T) {
	s := New[int](3)

	for i := 0; i < 3; i++ {
		err := s.Push(i + 1)

		if err != nil {
			t.Errorf("Expected error to be nil, got %v", err)
		}

		if s.top != i {
			t.Errorf("Expected top to be %d, got %d", i, s.top)
		}
	}

	err := s.Push(4)

	if err != ErrStackFull {
		t.Errorf("Expected error to be ErrStackFull, got %v", err)
	}
}

func TestPop(t *testing.T) {
	s := New[int](3)

	_, err := s.Pop()

	if err != ErrStackEmpty {
		t.Errorf("Expected error to be ErrStackEmpty, got %v", err)
	}

	for i := 0; i < 3; i++ {
		s.Push(i + 1)
	}

	for i := 3; i > 0; i-- {
		v, _ := s.Pop()

		if v != i {
			t.Errorf("Expected element to be %d, got %d", i, v)
		}
	}
}

func TestPeek(t *testing.T) {
	s := New[int](3)

	_, err := s.Peek()

	if err != ErrStackEmpty {
		t.Errorf("Expected error to be ErrStackEmpty, got %v", err)
	}

	for i := 0; i < 3; i++ {
		s.Push(i + 1)
	}

	v, _ := s.Peek()

	if v != 3 {
		t.Errorf("Expected element to be 3, got %d", v)
	}

	v, _ = s.Pop()

	if v != 3 {
		t.Errorf("Expected element to be 3, got %d", v)
	}
}

func TestSize(t *testing.T) {
	s := New[int](3)

	if s.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", s.Size())
	}

	for i := 0; i < 3; i++ {
		s.Push(i + 1)
	}

	if s.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", s.Size())
	}
}

func TestEmpty(t *testing.T) {
	s := New[int](3)

	if !s.Empty() {
		t.Errorf("Expected stack to be empty")
	}

	s.Push(1)

	if s.Empty() {
		t.Errorf("Expected stack to not be empty")
	}
}

func TestFull(t *testing.T) {
	s := New[int](3)

	if s.Full() {
		t.Errorf("Expected stack to not be full")
	}

	for i := 0; i < 3; i++ {
		s.Push(i + 1)
	}

	if !s.Full() {
		t.Errorf("Expected stack to be full")
	}
}

func TestValues(t *testing.T) {
	s := New[int](3)

	for i := 0; i < 3; i++ {
		s.Push(i + 1)
	}

	values := s.Values()

	for i := 0; i < 3; i++ {
		if values[i] != i+1 {
			t.Errorf("Expected element to be %d, got %d", i+1, values[i])
		}
	}
}
