package main

import (
	"fmt"
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
