package lists

// List is an interface that defines the methods that a list must implement
type List[T any] interface {
	// Retrieve the element at the given index
	Get(int) (T, bool)

	// Add an element to the beginning of the list
	Prepend(T)

	// Add an element to the end of the list
	Append(T)

	// Add an element at the given index
	Add(int, T) bool

	// Delete the element at the given index
	Delete(int) bool

	// Retrieve elements from the list
	Values() []T

	// Check if the list is empty
	Empty() bool

	// Check the size of the list
	Size() int
}
