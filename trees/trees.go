package trees

type Tree[T any] interface {
	// Insert an element into the left position
	InsertLeft(T) error

	// Insert an element into the right position
	InsertRight(T) error

	// Delete an element
	// Delete() error

	// Search for an element in the tree
	Search(T) bool

	// Traverse the tree in pre-order
	PreOrder() []T

	// Traverse the tree in in-order
	InOrder() []T

	// Traverse the tree in post-order
	PostOrder() []T
}
