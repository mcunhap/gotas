package arrayqueue

import "testing"

func TestNew(t *testing.T) {
	q := New[int](3)

	if q.capacity != 3 {
		t.Errorf("Expected capacity to be 3, got %d", q.capacity)
	}

	if q.front != -1 {
		t.Errorf("Expected front to be -1, got %d", q.front)
	}

	if q.rear != -1 {
		t.Errorf("Expected rear to be -1, got %d", q.rear)
	}

	if len(q.data) != 3 {
		t.Errorf("Expected length of data to be 3, got %d", len(q.data))
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int](3)

	for i := 0; i < 3; i++ {
		err := q.Enqueue(i + 1)

		if err != nil {
			t.Errorf("Expected error to be nil, got %v", err)
		}

		if q.rear != i {
			t.Errorf("Expected rear to be %d, got %d", i, q.rear)
		}
	}

	err := q.Enqueue(4)

	if err != ErrQueueFull {
		t.Errorf("Expected error to be ErrQueueFull, got %v", err)
	}
}

func TestDequeue(t *testing.T) {
	q := New[int](3)

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
	q := New[int](3)

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
	q := New[int](3)

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
	}

	if q.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", q.Size())
	}

	for i := 0; i < 3; i++ {
		q.Dequeue()
	}

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}
}

func TestEmpty(t *testing.T) {
	q := New[int](3)

	if !q.Empty() {
		t.Errorf("Expected queue to be empty, got %v", q.Empty())
	}

	q.Enqueue(1)

	if q.Empty() {
		t.Errorf("Expected queue to not be empty, got %v", q.Empty())
	}

	q.Enqueue(2)
	q.Enqueue(3)

	if q.Empty() {
		t.Errorf("Expected queue to not be empty, got %v", q.Empty())
	}

	for i := 0; i < 3; i++ {
		q.Dequeue()
	}

	if !q.Empty() {
		t.Errorf("Expected queue to be empty, got %v", q.Empty())
	}
}

func TestFull(t *testing.T) {
	q := New[int](3)

	if q.Full() {
		t.Errorf("Expected queue to not be full, got %v", q.Full())
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if !q.Full() {
		t.Errorf("Expected queue to be full, got %v", q.Full())
	}

	q.Dequeue()

	if q.Full() {
		t.Errorf("Expected queue to not be full, got %v", q.Full())
	}
}

func TestValues(t *testing.T) {
	q := New[int](3)

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
	}

	values := q.Values()

	for i := 0; i < 3; i++ {
		if values[i] != i+1 {
			t.Errorf("Expected value to be %d, got %d", i+1, values[i])
		}
	}
}
