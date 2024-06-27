package arraystack

import "errors"

type ArrayStack[T any] struct {
	data     []T
	top      int
	capacity int
}

var ErrStackFull = errors.New("stack is full")
var ErrStackEmpty = errors.New("stack is empty")

func New[T any](cap int) *ArrayStack[T] {
	return &ArrayStack[T]{data: make([]T, cap), top: -1, capacity: cap}
}

func (s *ArrayStack[T]) Push(data T) error {
	if s.Full() {
		return ErrStackFull
	}

	s.top++
	s.data[s.top] = data

	return nil
}

func (s *ArrayStack[T]) Pop() (T, error) {
	if s.Empty() {
		var t T
		return t, ErrStackEmpty
	}

	data := s.data[s.top]
	s.top--

	return data, nil
}

func (s *ArrayStack[T]) Peek() (T, error) {
	if s.Empty() {
		var t T
		return t, ErrStackEmpty
	}

	return s.data[s.top], nil
}

func (s *ArrayStack[T]) Size() int {
	return s.top + 1
}

func (s *ArrayStack[T]) Empty() bool {
	return s.top == -1
}

func (s *ArrayStack[T]) Full() bool {
	return s.top == s.capacity-1
}
func (s *ArrayStack[T]) Values() []T {
	return s.data
}
