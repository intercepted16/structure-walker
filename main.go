package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

import (
	. "structure-walker/Examples"
)

var (
	demonstration        string
	disableInteractivity bool
	verbose              bool
)

func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		if verbose {
			fmt.Println("Error loading .env file")
		}
	}

	// Function to get environment variable with a default value
	getEnv := func(key, defaultValue string) string {
		value, exists := os.LookupEnv(key)
		if !exists {
			return defaultValue
		}
		return value
	}

	verbose, err = strconv.ParseBool(getEnv("VERBOSE", "false"))
	if err != nil {
		if verbose {
			fmt.Println("Error parsing VERBOSE:", err)
		}
		verbose = false
	}

	// Example usage
	disableInteractivity, err = strconv.ParseBool(getEnv("DISABLE_INTERACTIVITY", "false"))
	if err != nil {
		if verbose {
			fmt.Println("Error parsing DISABLE_INTERACTIVITY:", err)
		}
		disableInteractivity = false
	}
	if verbose {
		fmt.Println("DISABLE_INTERACTIVITY:", disableInteractivity)
	}
}

func init() {
	LoadEnv()
}

func main() {
	err := huh.NewSelect[string]().
		Title("Choose a demonstration").
		Options(
			huh.NewOption("Linked List", "linked-list"),
			huh.NewOption("Add two numbers represented by linked lists", "add-two-numbers"),
			huh.NewOption("Hash Table", "hash-table"),
			huh.NewOption("Binary Search Tree", "binary-search-tree"),
			huh.NewOption("Queue", "queue"),
			huh.NewOption("Stack", "stack"),
			huh.NewOption("Doubly Linked List", "doubly-linked-list"),
			huh.NewOption("Breadth First Search", "bfs"),
			huh.NewOption("Depth First Search", "dfs"),
		).
		Value(&demonstration).Run()
	if err != nil {
		return
	}
	fmt.Println("You chose:", demonstration)
	switch demonstration {
	case "linked-list":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a linked list.",
				"It is a simple implementation of a linked list in C.",
				"It works by creating a `Node` struct that contains a value and a pointer to the next node.",
				"A linked list is a data structure that consists of nodes, each node containing a value and a pointer to the next node.",
				"This is contrasted with an array, where you have to resize the array to add or remove elements.",
				"For example, this isn't well known, but in Python, the 'arrays' are actually linked lists.",
				"They are especially useful when you don't know the size of the data you are working with.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}
		ExampleLinkedList()
		promptForCode("LinkedList/ListNode.c")
	case "add-two-numbers":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates adding two numbers, each number represented by a linked list of digits.",
				"It is a solution to the LeetCode problem for adding two numbers.",
				"It works by creating two linked lists, each representing a number.",
				"It then adds the two linked lists and returns the result as a linked list.",
				"It adds them manually in reverse order, carrying over the remainder to the next node.",
				"Almost as if you were adding two numbers on paper.",
			}

			// Add '\n' to each message
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})

			printMessagesWithDelay(messages, 2*time.Second)
		}

		ExampleAddTwoNumbers()
		promptForCode("LinkedList/ListNode.c", "LinkedList/addTwoNums.c")

	case "hash-table":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a hash table.",
				"It is a simple implementation of a hash table in C.",
				"It works by creating a `HashTable` struct that contains an array of linked lists.",
				"A hash table is a data structure that maps keys to values.",
				"It is useful for storing key-value pairs, and is often used in databases and caches.",
				"It is especially useful when you need to quickly look up a value based on a key.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}
		ExampleHashTable()
		promptForCode("HashTable/HashTable.h", "HashTable/HashTable.c")
	case "binary-search-tree":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a binary search tree.",
				"It is a simple implementation of a binary search tree in C.",
				"It works by creating a `Node` struct that contains a value and pointers to the left and right children.",
				"A binary search tree is a data structure that allows for efficient lookup, insertion, and deletion of values.",
				"It is especially useful when you need to maintain a sorted collection of values.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}

		ExampleBinaryTree()
		promptForCode("BinarySearchTree/BinaryTree.h", "BinarySearchTree/BinaryTree.c")
	case "queue":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a queue.",
				"It is relatively self explanatory; a grocery store has a queue, any waiting line does.",
				"It is a simple implementation of a queue in C.",
				"It works by creating a `Queue` struct that contains pointers to the head and tail of the queue.",
				"A queue is a data structure that allows for first-in, first-out (FIFO) access to elements.",
				"It is useful for implementing algorithms that require a first-in, first-out order, such as breadth-first search.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}
		ExampleQueue()
		promptForCode("Queue/Queue.h", "Queue/Queue.c")

	case "stack":
		// Print the introductory message
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a stack.",
				"It is a simple implementation of a stack in C.",
				"It works by creating a `Stack` struct that contains a pointer to the top of the stack.",
				"A stack is a data structure that allows for last-in, first-out (LIFO) access to elements.",
				"It is useful for implementing algorithms that require a last-in, first-out order, such as depth-first search.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}
		ExampleStack()
		promptForCode("Stack/Stack.h", "Stack/Stack.c")
	case "doubly-linked-list":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a doubly linked list.",
				"It is a simple implementation of a doubly linked list in C.",
				"It works by creating a `DoubleNode` struct that contains a value and pointers to the next and previous nodes.",
				"A doubly linked list is a data structure that allows for efficient insertion and deletion of elements.",
				"It is especially useful when you need to quickly insert or delete elements in the middle of a list.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}
		ExampleDoublyLinkedList()
		promptForCode("DoublyLinkedList/DoublyLinkedList.h", "DoublyLinkedList/DoublyLinkedList.c")
	case "bfs":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a breadth-first search.",
				"It is a simple implementation of a breadth-first search in C.",
				"It works by creating a `Queue` struct that contains a queue of paths.",
				"A breadth-first search is an algorithm that searches a tree or graph level by level.",
				"It is useful for finding the shortest path between two nodes.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}
		ExampleBreadthFirstSearch()
		promptForCode("BreadthFirstSearch/BreadthFirstSearch.go")
	case "dfs":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a depth-first search.",
				"It is a simple implementation of a depth-first search in C.",
				"It works by creating a `Stack` struct that contains a stack of paths.",
				"A depth-first search is an algorithm that searches a tree or graph by exploring as far as possible along each branch before backtracking.",
				"It is useful for finding a path between two nodes, but may not find the shortest path.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}
		ExampleDepthFirstSearch()
		promptForCode("DepthFirstSearch/Stack.go", "DepthFirstSearch/DepthFirstSearch.go")
	}
}
