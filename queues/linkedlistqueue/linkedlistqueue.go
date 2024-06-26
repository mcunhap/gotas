package linkedlistqueue

import (
	"errors"
	"gotas/lists/singlylinkedlist"
)

type LinkedListQueue[T any] struct {
	data *singlylinkedlist.List[T]
}

var ErrQueueEmpty = errors.New("queue is empty")

func New[T any]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{data: singlylinkedlist.New[T]()}
}

func (q *LinkedListQueue[T]) Enqueue(data T) error {
	q.data.Append(data)

	return nil
}

func (q *LinkedListQueue[T]) Dequeue() (T, error) {
	data, ok := q.data.Get(0)

	if !ok {
		var t T
		return t, ErrQueueEmpty
	}

	q.data.Delete(0)

	return data, nil
}

func (q *LinkedListQueue[T]) Front() (T, error) {
	data, ok := q.data.Get(0)

	if !ok {
		var t T
		return t, ErrQueueEmpty
	}

	return data, nil
}

func (q *LinkedListQueue[T]) Size() int {
	return q.data.Size()
}

func (q *LinkedListQueue[T]) Empty() bool {
	return q.data.Empty()
}

func (q *LinkedListQueue[T]) Full() bool {
	return false
}

func (q *LinkedListQueue[T]) Values() []T {
	return q.data.Values()
}
