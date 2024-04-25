package lists

type List[T comparable] interface {
	// Retrieve the first element of the list
	GetBegin() T

	// Retrieve the last element of the list
	GetEnd() T

	// Retrieve the element at the given index
	GetIndex(int) T

	// Add an element to the beginning of the list
	AddBegin(T)

	// Add an element to the end of the list
	AddEnd(T)

	// Add an element at the given index
	AddIndex(int, T)

	// Delete the first element of the list
	DeleteBegin()

	// Delete the last element of the list
	DeleteEnd()

	// Delete the element at the given index
	DeleteIndex(int)

	// Display the list
	Display()
}
