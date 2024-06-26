package examples

import (
	"fmt"
	"gotas/queues/linkedlistqueue"
)

func LinkedListQueueExample() {
	queue := linkedlistqueue.New[int]()

	fmt.Println("Trying to dequeue from an empty queue...")
	_, err := queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nEnqueuing values 1, 2 and 3...")
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Values())

	fmt.Println("\nPeeking the front element...")
	value, _ := queue.Front()
	fmt.Println(value)

	fmt.Println("\nDequeing all elements...")
	for !queue.Empty() {
		value, _ := queue.Dequeue()
		fmt.Println(value)
	}

	fmt.Println("\nChecking if the queue is empty...")
	fmt.Println(queue.Empty())
}
