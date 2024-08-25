package Examples

/*
#include "../Queue/Queue.c"
*/
import "C"

import "fmt"

// PrintQueue prints all elements in the queue from head to tail
func PrintQueue(q *C.Queue) {
	current := q.head

	if current == nil {
		fmt.Println("Queue is empty.")
		return
	}

	for current != nil {
		fmt.Printf("%d", current.val)
		if current.next != nil {
			fmt.Printf(" -> ")
		}
		current = current.next
	}
	fmt.Print(" -> NULL\n")
}

// ExampleQueue demonstrates how to use the Queue data structure in C
func ExampleQueue() {
	// Create a new queue
	queue := C.createQueue()
	if queue == nil {
		fmt.Println("Failed to create queue")
		return
	}
	defer C.freeQueue(queue) // Ensure the queue is freed when done
	// Enqueue some values
	C.enqueue(queue, C.int(1))
	C.enqueue(queue, C.int(2))
	C.enqueue(queue, C.int(3))
	// Print the queue
	PrintQueue(queue)
}
