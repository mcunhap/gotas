package main

import (
	"fmt"
	"gotas/examples"
)

func main() {
	fmt.Printf("==== Singly Linked List ===\n")
	examples.SinglyLinkedListExample()
	fmt.Printf("\n=========================\n")

	fmt.Printf("\n==== Doubly Linked List ====\n")
	examples.DoublyLinkedListExample()
	fmt.Printf("\n============================\n")

	fmt.Printf("\n==== Circular Linked List ====\n")
	examples.CircularLinkedListExample()
	fmt.Printf("\n==============================\n")
}
