package examples

import (
	"fmt"
	"gotas/stacks/linkedliststack"
)

func LinkedListStackExample() {
	// Create a new linked list stack
	stack := linkedliststack.New[int]()

	fmt.Println("Trying to pop from an empty stack...")
	_, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nPushing values 1, 2 and 3...")
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Values())

	fmt.Println("\nPeeking the last element...")
	value, _ := stack.Peek()
	fmt.Println(value)

	fmt.Println("\nPoping all elements...")
	for !stack.Empty() {
		value, _ := stack.Pop()
		fmt.Println(value)
	}

	fmt.Println("\nChecking if the stack is empty...")
	fmt.Println(stack.Empty())
}
