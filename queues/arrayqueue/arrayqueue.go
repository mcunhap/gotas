package arrayqueue

import (
	"errors"
)

type ArrayQueue[T any] struct {
	data     []T
	front    int
	rear     int
	capacity int
}

var ErrQueueFull = errors.New("queue is full")
var ErrQueueEmpty = errors.New("queue is empty")

func New[T any](cap int) *ArrayQueue[T] {
	return &ArrayQueue[T]{data: make([]T, cap), front: -1, rear: -1, capacity: cap}
}

func (q *ArrayQueue[T]) Enqueue(data T) error {
	if q.Full() {
		return ErrQueueFull
	}

	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = data

	if q.front == -1 {
		q.front = q.rear
	}

	return nil
}

func (q *ArrayQueue[T]) Dequeue() (T, error) {
	if q.Empty() {
		var t T
		return t, ErrQueueEmpty
	}

	data := q.data[q.front]

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front + 1) % q.capacity
	}

	return data, nil
}

func (q *ArrayQueue[T]) Front() (T, error) {
	if q.Empty() {
		var t T
		return t, ErrQueueEmpty
	}

	return q.data[q.front], nil
}

func (q *ArrayQueue[T]) Size() int {
	if q.Empty() {
		return 0
	}

	return (q.capacity-q.front+q.rear)%q.capacity + 1
}

func (q *ArrayQueue[T]) Empty() bool {
	return q.front == -1
}

func (q *ArrayQueue[T]) Full() bool {
	return (q.rear+1)%q.capacity == q.front
}

func (q *ArrayQueue[T]) Values() []T {
	return q.data
}
