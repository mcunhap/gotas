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

	fmt.Printf("\n==== Array Stack ====\n")
	examples.ArrayStackExample()
	fmt.Printf("\n=====================\n")

	fmt.Printf("\n==== Linked List Stack ====\n")
	examples.LinkedListStackExample()
	fmt.Printf("\n=====================\n")

	fmt.Printf("\n==== Array Queue ====\n")
	examples.ArrayQueueExample()
	fmt.Printf("\n=====================\n")

	fmt.Printf("\n==== Linked List Queue ====\n")
	examples.LinkedListQueueExample()
	fmt.Printf("\n=====================\n")
}
