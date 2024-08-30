package Examples

/*
#include "../BTree/BTree.c"
*/
import "C"
import "fmt"

// PrintBTree prints the BTree structure recursively
func PrintBTree(node *C.BNode, indent string, last bool) {
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

	// Print all keys in the node
	for i := 0; i < int(node.numKeys); i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d", node.keys[i])
	}
	fmt.Println()

	// Recursively print the children
	for i := 0; i <= int(node.numKeys); i++ {
		if node.children[i] != nil {
			PrintBTree(node.children[i], indent, i == int(node.numKeys))
		}
	}
}

func ExampleBTree() {
	// Create a new BTree
	btree := C.createBNode()
	// Insert some values into the BTree
	btree = C.insertBNode(&btree, C.int(5))
	btree = C.insertBNode(&btree, C.int(3))
	btree = C.insertBNode(&btree, C.int(7))
	btree = C.insertBNode(&btree, C.int(1))
	btree = C.insertBNode(&btree, C.int(9))

	// Print the BTree
	PrintBTree(btree, "", true)

	// Free the memory
	C.freeBNode(btree)
}
