package queues

// Queue is an interface that defines the methods that a queue must implement
type Queue[T any] interface {

	// Enqueue adds a new element at the end of the queue
	Enqueue(T) error

	// Dequeue removes the element from the front of the queue and returns it
	Dequeue() (T, error)

	// Front returns the element from the front of the queue without removing it
	Front() (T, error)

	// Check the size of the stack
	Size() int

	// Check if the stack is empty
	Empty() bool

	// Check if the stack is full
	Full() bool

	// Retrieve elements from the stack
	Values() []T
}
