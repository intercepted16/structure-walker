package Examples

/*
#include "../BinarySearchTree/BinaryTree.c"
*/
import "C"

// Function to print the binary search tree
func printTree(bst *C.TreeNode) {
	if bst == nil {
		return
	}
	// Print the left subtree
	printTree(bst.left)
	// Print the current node value
	print(int(bst.val), " ")
	// Print the right subtree
	printTree(bst.right)
}

func ExampleBinaryTree() {
	// Create a new binary search tree with the root value 1
	bst := C.createNode(C.int(1))
	if bst == nil {
		return
	}
	defer C.freeNode(bst)

	// Insert some values into the binary search tree
	C.insertNode(bst, C.int(5))
	C.insertNode(bst, C.int(3))
	C.insertNode(bst, C.int(7))
	C.insertNode(bst, C.int(2))
	C.insertNode(bst, C.int(4))
	C.insertNode(bst, C.int(6))
	C.insertNode(bst, C.int(8))

	// Print the binary search tree
	printTree(bst)
}
