package examples

import (
	"fmt"
	"gotas/lists/doublylinkedlist"
)

func DoublyLinkedListExample() {
	list := doublylinkedlist.New[int]()

	fmt.Println("Appending values 1, 2 and 3...")
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Printf("%v", list.Values())

	fmt.Println("\n\nPrepending value 0...")
	list.Prepend(0)
	fmt.Println(list.Values())

	fmt.Println("\nAdding value 4 at index 2...")
	list.Add(2, 4)
	fmt.Println(list.Values())

	fmt.Println("\nDeleting element at index 3...")
	list.Delete(3)
	fmt.Println(list.Values())

	fmt.Println("\nDeleting first element...")
	list.Delete(0)
	fmt.Println(list.Values())

	fmt.Println("\nDeleting last element...")
	list.Delete(2)
	fmt.Println(list.Values())

	fmt.Println("\nChecking if the list is empty...")
	fmt.Println(list.Empty())
}
