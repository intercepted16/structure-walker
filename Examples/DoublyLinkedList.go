package Examples

/*
#include "../DoublyLinkedList/DoublyLinkedList.c"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func PrintDoubleList(list *C.struct_DoubleNode) {
	// Print the doubly linked list
	node := list
	for node != nil {
		if node.val != nil {
			val := *(*C.int)(node.val)
			fmt.Print(val, "->")
		} else {
			fmt.Print("nil ")
		}
		node = node.next
	}
	println("nil")
}

func PrintDoubleListBackwards(list *C.struct_DoubleNode) {
	// Start from the end of the doubly linked list
	node := list
	for node.next != nil {
		node = node.next
	}
	// Print the doubly linked list backwards
	for node != nil {
		if node.val != nil {
			val := *(*C.int)(node.val)
			fmt.Print(val, "->")
		} else {
			fmt.Print("nil ")
		}
		node = node.prev
	}
	println("nil")
}

// ExampleDoublyLinkedList Demonstrates a doubly linked list
func ExampleDoublyLinkedList() {
	// Create a new doubly linked list
	list := C.createDoubleNode()
	if list == nil {
		return
	}
	defer C.freeDoubleNode(list)
	// Insert some values
	val1 := C.int(1)
	val2 := C.int(2)
	val3 := C.int(3)
	C.appendDoubleNode(&list, unsafe.Pointer(&val1))
	C.appendDoubleNode(&list, unsafe.Pointer(&val2))
	C.appendDoubleNode(&list, unsafe.Pointer(&val3))
	// Print the doubly linked list
	PrintDoubleList(list)
	// Print the doubly linked list backwards
	PrintDoubleListBackwards(list)
}
