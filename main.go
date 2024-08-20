package main

import "C"
import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

import (
	. "data-structures/linkedListExamples"
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
		).
		Value(&demonstration).Run()
	if err != nil {
		return
	}
	println("You chose:", demonstration)
	switch demonstration {
	case "linked-list":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates a linked list.",
				"It is a simple implementation of a linked list in C.",
				"It works by creating a `Node` struct that contains a value and a pointer to the next node.",
				"A linked list is a data structure that consists of nodes, each node containing a value and a pointer to the next node.",
				"This is contrasted with an array, where you have to resize the array to add or remove elements.",
				"For example, this isn't well known, but, in Python, the 'arrays' are actually linked lists.",
				"They are especially useful when you don't know the size of the data you are working with.",
			}
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})
			printMessagesWithDelay(messages, 2*time.Second)
		}
		ExampleLinkedList()
		err := huh.NewSelect[string]().
			Title("Do you want to see the code?").
			Options(
				huh.NewOption("Yes", "yes"),
				huh.NewOption("No", "no"),
			).
			Value(&demonstration).Run()
		if err != nil {
			return
		}
		if demonstration == "yes" {
			err := readAndPrintCode("linkedList/ListNode.c")
			if err != nil {
				return
			}
			err = readAndPrintCode("linkedList/addTwoNums.c")
			if err != nil {
				return
			}
		}
	case "add-two-numbers":
		if !disableInteractivity {
			messages := []string{
				"This demonstrates adding two numbers, each number represented by a linked list of digits.",
				"It is a solution to the LeetCode problem for adding two numbers.",
				"It works by creating two linked lists, each representing a number.",
				"It then adds the two linked lists and returns the result as a linked list.",
				"it adds them manually in reverse order, carrying over the remainder to the next node.",
				"Almost as if you were adding two numbers on paper.",
			}

			// add '\n' to each message
			messages = Map(messages, func(message string) string {
				return message + "\n"
			})

			printMessagesWithDelay(messages, 2*time.Second)
		}

		ExampleAddTwoNumbers()

		// add a button to choose if you want to see the code
		err := huh.NewSelect[string]().
			Title("Do you want to see the code?").
			Options(
				huh.NewOption("Yes", "yes"),
				huh.NewOption("No", "no"),
			).
			Value(&demonstration).Run()
		if err != nil {
			return
		}
		if demonstration == "yes" {
			// Print the introductory message
			fmt.Println("Here is the code for adding two numbers using linked lists")
			err := readAndPrintCode("linkedList/addTwoNums.c")
			if err != nil {
				return
			}

		}
	}
}
