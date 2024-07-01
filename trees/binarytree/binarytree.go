package binarytree

import (
	"errors"
	"gotas/stacks/linkedliststack"
)

type Node[T comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

type child string

const (
	left  child = "left"
	right child = "right"
)

var ErrInvalidInsert = errors.New("cannot insert into a node that already has a left/right child")
var ErrInvalidDelete = errors.New("cannot delete a non-leaf node")

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (n *Node[T]) InsertLeft(value T) (*Node[T], error) {
	if n.left != nil {
		return nil, ErrInvalidInsert
	}

	left := NewNode(value)
	n.left = left

	return left, nil
}

func (n *Node[T]) InsertRight(value T) (*Node[T], error) {
	if n.right != nil {
		return nil, ErrInvalidInsert
	}

	right := NewNode(value)
	n.right = right

	return right, nil
}

// func (n *Node[T]) Delete() error {}

func (n *Node[T]) Search(data T) bool {
	// using inorder traversal to search for the value
	// inorder -> left, root, right
	var current *Node[T]

	s := linkedliststack.New[*Node[T]]()
	current = n

	for current != nil || !s.Empty() {
		for current != nil {
			s.Push(current)
			current = current.left
		}

		current, _ = s.Pop()
		if current.value == data {
			return true
		}

		current = current.right
	}

	return false
}

func (n *Node[T]) PreOrder() []T {
	// preorder -> root, left, right
	var current *Node[T]
	var r []T

	s := linkedliststack.New[*Node[T]]()
	current = n

	s.Push(current)

	for !s.Empty() {
		current, _ = s.Pop()
		r = append(r, current.value)

		if current.right != nil {
			s.Push(current.right)
		}

		if current.left != nil {
			s.Push(current.left)
		}
	}

	return r
}

func (n *Node[T]) InOrder() []T {
	// inorder -> left, root, right
	var current *Node[T]
	var r []T

	s := linkedliststack.New[*Node[T]]()
	current = n

	for current != nil || !s.Empty() {
		for current != nil {
			s.Push(current)
			current = current.left
		}

		current, _ = s.Pop()
		r = append(r, current.value)
		current = current.right
	}

	return r
}

func (n *Node[T]) PostOrder() []T {
	// postorder -> left, right, root
	var current *Node[T]
	var r []T

	s := linkedliststack.New[*Node[T]]()
	current = n

	for current != nil || !s.Empty() {

		for current != nil {
			s.Push(current)
			current = current.left
		}

		current, _ = s.Pop()
		r = append(r, current.value)

		if current.right != nil {
			s.Push(current.right)
		}

		current = current.right
	}

	return r
}
