package Examples

/*
#include "../MaxHeapTree/MaxHeapTree.c"
*/
import "C"
import "fmt"

func printMaxHeap(node *C.struct_MaxHeapNode, indent string, last bool) {
	if node == nil {
		return
	}

	// Print the current node
	fmt.Print(indent)
	if last {
		fmt.Print("└── ")
		indent += "    "
	} else {
		fmt.Print("├── ")
		indent += "│   "
	}
	fmt.Println(node.val)

	// Recursively print the left and right subtrees
	if node.left != nil || node.right != nil {
		if node.left != nil {
			printMaxHeap(node.left, indent, false)
		}
		if node.right != nil {
			printMaxHeap(node.right, indent, true)
		}
	}
}

// ExampleMaxHeap demonstrates a maximum heap BST
func ExampleMaxHeap() {
	maxHeap := C.createMaxHeapNode(10)
	C.insertMaxHeapNode(&maxHeap, 20)
	C.insertMaxHeapNode(&maxHeap, 30)
	// print the values in the max heap
	printMaxHeap(maxHeap, "", true)
}
