package main

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/charmbracelet/lipgloss"
	"os"
	"os/exec"
	"time"
)

// Function to print each message with a delay
func printMessagesWithDelay(messages []string, delay time.Duration) {
	for _, message := range messages {
		fmt.Println(message)
		time.Sleep(delay)
	}
}

func Map(slice []string, fn func(string) string) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Function that takes a buffer, and automatically syntax highlights it
// It adds a border around it
func syntaxHighlight(buffer string) (string, error) {
	// Set up the lexer, formatter, and style
	lexer := lexers.Get("c")
	formatter := formatters.Get("terminal256")
	style := styles.Get("monokai")

	// Tokenize the code
	iterator, err := lexer.Tokenise(nil, buffer)
	if err != nil {
		fmt.Println("Error tokenizing code:", err)
		return "", err
	}

	// Create a buffer to capture the formatted output
	var highlightedCode bytes.Buffer
	err = formatter.Format(&highlightedCode, style, iterator)
	if err != nil {
		fmt.Println("Error formatting code:", err)
		return "", err
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
	return styledCode, nil
}

// Function to page/paginate a string
// It uses the `less` command to display the text
func paginateString(text string) error {
	// Set up the pager
	cmd := exec.Command("less", "-R") // Use -R to interpret raw control characters
	cmd.Stdin = bytes.NewReader([]byte(text))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the pager
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running pager:", err)
		return err
	}
	// after the pager is closed, clear the output of this program
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running clear:", err)
	}
	return nil
}

func readAndPrintCode(fileName string) error {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
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
	// syntax highlight the code
	styledCode, err := syntaxHighlight(fileContent.String())
	if err != nil {
		return err
	}
	// print paginated code
	err = paginateString(styledCode)
	if err != nil {
		return err
	}
	return nil
}
