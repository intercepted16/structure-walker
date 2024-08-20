package main

/*
#include "linkedList/ListNode.c"
#include "linkedList/addTwoNums.c"
*/
import "C"
import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"os"
	"os/exec"
	"time"
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
		// debug: check that the messages end with '\n'
		for _, message := range messages {
			if message[len(message)-1] != '\n' {
				println("The message does not end with a newline character")
			}
		}

		printMessagesWithDelay(messages, 2*time.Second)

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
				fmt.Printf("%d + %d = %d\n", val1, val2, int(*(*C.int)(newNode.val)))
			}
			newNode = newNode.next
		}

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

			// Open the file
			file, err := os.Open("linkedList/addTwoNums.c")
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					return
				}
			}(file)

			// Read the file content
			var fileContent bytes.Buffer
			buf := make([]byte, 1024)
			for {
				n, err := file.Read(buf)
				if err != nil {
					if err.Error() != "EOF" {
						fmt.Println("Error reading file:", err)
					}
					break
				}
				fileContent.Write(buf[:n])
			}

			// Set up the lexer, formatter, and style
			lexer := lexers.Get("c")
			formatter := formatters.Get("terminal256")
			style := styles.Get("monokai")

			// Tokenize the code
			iterator, err := lexer.Tokenise(nil, fileContent.String())
			if err != nil {
				fmt.Println("Error tokenizing code:", err)
				return
			}

			// Create a buffer to capture the formatted output
			var highlightedCode bytes.Buffer
			err = formatter.Format(&highlightedCode, style, iterator)
			if err != nil {
				fmt.Println("Error formatting code:", err)
				return
			}

			// Use lipgloss to style the entire block
			codeBlockStyle := lipgloss.NewStyle().
				Background(lipgloss.Color("black")).
				Foreground(lipgloss.Color("white")).
				Padding(1, 2).
				Margin(1, 2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("cyan"))

			// Render the styled code block
			styledCode := codeBlockStyle.Render(highlightedCode.String())
			// Set up the pager
			cmd := exec.Command("less", "-R") // Use -R to interpret raw control characters
			cmd.Stdin = bytes.NewReader([]byte(styledCode))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			// Run the pager
			err = cmd.Run()
			if err != nil {
				fmt.Println("Error running pager:", err)
			}
			// after the pager is closed, clear the output of this program
			cmd = exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println("Error running clear:", err)
			}
		}
	}
}
