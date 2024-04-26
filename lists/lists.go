package lists

type List[T comparable] interface {
	// Retrieve the first element of the list
	GetBegin() (T, bool)

	// Retrieve the last element of the list
	GetEnd() (T, bool)

	// Retrieve the element at the given index
	Get(int) (T, bool)

	// Add an element to the beginning of the list
	Prepend(T)

	// Add an element to the end of the list
	Append(T)

	// Add an element at the given index
	Add(int, T) bool

	// Delete the first element of the list
	DeleteBegin() bool

	// Delete the last element of the list
	DeleteEnd() bool

	// Delete the element at the given index
	Delete(int) bool

	// Retrieve elements from the list
	Values() []T

	// Check if the list is empty
	Empty() bool
}
