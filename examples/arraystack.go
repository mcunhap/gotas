package examples

import (
	"fmt"
	"gotas/stacks/arraystack"
)

func ArrayStackExample() {
	stack := arraystack.New[int](3)

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

	fmt.Println("\nTrying to push to a full stack...")
	err = stack.Push(4)
	if err != nil {
		fmt.Println(err)
	}

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
