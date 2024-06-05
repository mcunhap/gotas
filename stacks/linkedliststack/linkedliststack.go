package linkedliststack

import (
	"errors"
	"gotas/lists/singlylinkedlist"
)

type LinkedListStack[T any] struct {
	data *singlylinkedlist.List[T]
}

var ErrStackEmpty = errors.New("stack is empty")

func New[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{data: singlylinkedlist.New[T]()}
}

func (s *LinkedListStack[T]) Push(data T) error {
	s.data.Append(data)

	return nil
}

func (s *LinkedListStack[T]) Pop() (T, error) {
	if s.data.Empty() {
		var t T
		return t, ErrStackEmpty
	}

	v, _ := s.data.Get(s.data.Size() - 1)
	s.data.Delete(s.data.Size() - 1)

	return v, nil
}

func (s *LinkedListStack[T]) Peek() (T, error) {
	if s.Empty() {
		var t T
		return t, ErrStackEmpty
	}

	v, _ := s.data.Get(s.data.Size() - 1)

	return v, nil
}

func (s *LinkedListStack[T]) Size() int {
	return s.data.Size()
}

func (s *LinkedListStack[T]) Empty() bool {
	return s.data.Empty()
}

func (s *LinkedListStack[T]) Full() bool {
	return false
}

func (s *LinkedListStack[T]) Values() []T {
	values := []T{}

	for i := 0; i < s.data.Size(); i++ {
		v, _ := s.data.Get(i)
		values = append(values, v)
	}

	return values
}
