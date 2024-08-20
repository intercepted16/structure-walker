package Examples

import "C"
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

/*
#include "../linkedList/ListNode.c"
#include "../linkedList/addTwoNums.c"
*/
import "C"

func ExampleLinkedList() {
	l1 := C.createNode(nil, 0)
	l2 := C.createNode(nil, 0)
	current := l1

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter input: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Strip the newline character
		inputStr := strings.TrimSpace(input)

		fmt.Printf("You entered: %s\n", inputStr)

		if inputStr == "q" {
			break
		}

		// Create a new node
		data := C.CString(inputStr)
		size := C.size_t(len(inputStr) + 1)
		newNode := C.createNode(unsafe.Pointer(data), size)
		C.free(unsafe.Pointer(data))

		current.next = newNode
		current = newNode
	}

	C.freeList(l1)
	C.freeList(l2)
}

func ExampleAddTwoNumbers() {
	val1 := 1
	val2 := 2
	node1 := C.createNode(unsafe.Pointer(&val1), 1)
	node2 := C.createNode(unsafe.Pointer(&val2), 1)

	newNode := C.addTwoNumbers(node1, node2)

	if newNode == nil {
		fmt.Println("The node is nil")
		return
	}

	for newNode != nil {
		if newNode.val != nil {
			fmt.Printf("%d + %d = %d\n", val1, val2, int(*(*C.int)(newNode.val)))
		}
		newNode = newNode.next
	}
}
