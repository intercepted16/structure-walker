package Examples

/*
#include "../BinarySearchTree/BinaryTree.c"
*/
import "C"
import (
	"fmt"
)

func printTree(node *C.struct_TreeNode, indent string, last bool) {
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
			printTree(node.left, indent, false)
		}
		if node.right != nil {
			printTree(node.right, indent, true)
		}
	}
}

func ExampleBinaryTree() {
	// Create a new binary search tree with the root value 1

	bst := C.createTreeNode(C.int(1))
	if bst == nil {
		return
	}
	defer C.freeTreeNode(bst)
	// Insert some values into the binary search tree
	C.insertTreeNode(&bst, C.int(5))
	C.insertTreeNode(&bst, C.int(3))
	println("left subtree", bst.left)

	// Print the binary search tree
	printTree(bst, "", true)
}
