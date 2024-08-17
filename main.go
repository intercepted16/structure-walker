package main

/*
#include "linkedList/ListNode.c"
#include "linkedList/addTwoNums.c"
*/
import "C"
import (
	"github.com/charmbracelet/huh"
	"unsafe"
)

var (
	demonstration string
)

func main() {
	err := huh.NewSelect[string]().
		Title("Choose a demonstration").
		Options(
			huh.NewOption("Linked List", "linked-list"),
			huh.NewOption("Add two numbers", "add-two-numbers"),
		).
		Value(&demonstration).Run()
	if err != nil {
		return
	}
	println("You chose:", demonstration)
	switch demonstration {
	case "linked-list":
		C.exampleLinkedList()
	case "add-two-numbers":
		// Create two linked lists
		val1 := 1
		val2 := 2
		node1 := C.createNode(unsafe.Pointer(&val1), 1)
		node2 := C.createNode(unsafe.Pointer(&val2), 1)
		// Add the two linked lists
		newNode := C.addTwoNumbers(node1, node2)
		// check if the node is nil
		if newNode == nil {
			println("The node is nil")
			return
		}
		// Print the result
		for newNode != nil {
			if newNode.val != nil {
				println("Value: ", *(*C.int)(newNode.val))
			}
			newNode = newNode.next
		}
	}
}
