// premium quality code
// do not steal
// \joke
// my code is actually pretty bad, do whatever the hell
// you want with it
package Examples

/*
#include "../Trie/Trie.c"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Function to print the Trie structure recursively
func printTrie(node *C.struct_TrieNode, indent string, last bool, word []rune) {
	if node == nil {
		return
	}

	// Skip printing for the root node
	if len(word) > 0 {
		// Print the current structure with indentation
		fmt.Print(indent)
		if last {
			fmt.Print("└── ")
			indent += "    "
		} else {
			fmt.Print("├── ")
			indent += "│   "
		}

		// If the current node is the end of a word, print the word
		if node.isEndOfWord {
			fmt.Printf("%s (end)\n", string(word))
		} else {
			fmt.Printf("%s\n", string(word))
		}
	}

	// Variable to check if there's a non-nil child (for proper last element handling)
	hasChildren := false
	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			hasChildren = true
			break
		}
	}

	// Recursively print the children nodes
	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			// Append the current character to the word
			char := 'a' + rune(i)
			newWord := append(word, char)
			// Print recursively with correct 'last' indicator
			printTrie(node.children[i], indent, i == 25 || !hasChildren, newWord)
		}
	}
}

// ExampleTrie demonstrates a trie data structure
func ExampleTrie() {
	// create a Trie
	trie := C.createTrieNode()
	if trie == nil {
		return
	}
	defer C.freeTrieNode(trie)

	val1 := C.CString("apple")
	val2 := C.CString("app")
	defer C.free(unsafe.Pointer(val1))
	defer C.free(unsafe.Pointer(val2))

	C.insertTrieNode(&trie, val1)
	C.insertTrieNode(&trie, val2)
	printTrie(trie, "", true, []rune{})
}
