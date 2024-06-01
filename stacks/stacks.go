package stacks

// Stack is an interface that defines the methods that a stack must implement
type Stack[T any] interface {
	// Push an element onto the stack
	Push(T) error

	// Pop an element from the stack
	Pop() (T, error)

	// Peek at the top element of the stack
	Peek() (T, error)

	// Check the size of the stack
	Size() int

	// Check if the stack is empty
	Empty() bool

	// Check if the stack is full
	Full() bool

	// Retrieve elements from the stack
	Values() []T
}
