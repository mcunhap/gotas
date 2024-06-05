package linkedliststack

import "testing"

func TestNew(t *testing.T) {
	s := New[int]()

	if s.data.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", s.data.Size())
	}
}

func TestPush(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.data.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", s.data.Size())
	}

	if v, _ := s.Peek(); v != 3 {
		t.Errorf("Expected element to be 3, got %d", v)
	}
}

func TestPop(t *testing.T) {
	s := New[int]()

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
	s := New[int]()

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
	s := New[int]()

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
	s := New[int]()

	if !s.Empty() {
		t.Errorf("Expected stack to be empty")
	}

	s.Push(1)

	if s.Empty() {
		t.Errorf("Expected stack to not be empty")
	}
}

func TestFull(t *testing.T) {
	s := New[int]()

	if s.Full() {
		t.Errorf("Expected stack to not be full")
	}

	for i := 0; i < 3; i++ {
		s.Push(i + 1)
	}

	if s.Full() {
		t.Errorf("Expected stack not to be full")
	}
}

func TestValues(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	values := s.Values()

	for i := 0; i < 3; i++ {
		if values[i] != i+1 {
			t.Errorf("Expected element to be %d, got %d", i+1, values[i])
		}
	}
}
