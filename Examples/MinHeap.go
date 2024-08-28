package Examples

/*
#include "../MinHeapTree/MinHeapTree.c"
*/
import "C"

// ExampleMinHeap demonstrates a minimum heap BST
func ExampleMinHeap() {
	minHeap := C.createMinHeapNode(10)
	C.insertMinHeapNode(&minHeap, 20)
	C.insertMinHeapNode(&minHeap, 30)
	// print the values in the min heap
	printTree(minHeap, "", true)
}
