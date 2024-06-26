package linkedlistqueue

import "testing"

func TestNew(t *testing.T) {
	q := New[int]()

	if q.data.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.data.Size())
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int]()

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
	}

	if q.data.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", q.data.Size())
	}
}

func TestDequeue(t *testing.T) {
	q := New[int]()

	_, err := q.Dequeue()

	if err != ErrQueueEmpty {
		t.Errorf("Expected error to be ErrQueueEmpty, got %v", err)
	}

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
	}

	for i := 1; i <= 3; i++ {
		v, _ := q.Dequeue()

		if v != i {
			t.Errorf("Expected value to be %d, got %d", i, v)
		}
	}
}

func TestFront(t *testing.T) {
	q := New[int]()

	_, err := q.Front()

	if err != ErrQueueEmpty {
		t.Errorf("Expected error to be ErrQueueEmpty, got %v", err)
	}

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
	}

	v, _ := q.Front()

	if v != 1 {
		t.Errorf("Expected value to be 1, got %d", v)
	}

	v, _ = q.Dequeue()

	if v != 1 {
		t.Errorf("Expected value to be 1, got %d", v)
	}
}

func TestSize(t *testing.T) {
	q := New[int]()

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
	}

	if q.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", q.Size())
	}
}

func TestEmpty(t *testing.T) {
	q := New[int]()

	if !q.Empty() {
		t.Errorf("Expected queue to be empty, got not empty")
	}

	q.Enqueue(1)

	if q.Empty() {
		t.Errorf("Expected queue to be not empty, got empty")
	}
}

func TestFull(t *testing.T) {
	q := New[int]()

	if q.Full() {
		t.Errorf("Expected queue to be not full, got full")
	}
}

func TestValues(t *testing.T) {
	q := New[int]()

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
	}

	values := q.Values()

	for i, v := range values {
		if v != i+1 {
			t.Errorf("Expected value to be %d, got %d", i+1, v)
		}
	}
}
