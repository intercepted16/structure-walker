package main

/*
#include "linkedList/ListNode.c"
#include "linkedList/addTwoNums.c"
*/
import "C"
import (
	"github.com/charmbracelet/huh"
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
		node1 := C.createNode(1)
		node2 := C.createNode(2)
		C.addTwoNumbers(node1, node2)
	}
}
