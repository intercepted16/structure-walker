package DepthFirstSearch

import (
	"fmt"
	. "structure-walker/Maze"
)

// SolveMaze solves a maze represented as a 2D array of numbers.
// It uses Depth First Search (DFS).
// This does not return the most efficient method if there are multiple routes.
// It may be more or less efficient than Breadth First Search (BFS) depending on the maze.
func SolveMaze(maze Maze) [][]int {
	start, end := FindStartAndEnd(maze)
	stack := createStack([][]int{start})
	fmt.Printf("Stack is %d\n", stack.tail.value)

	// Directions for moving in the maze (Up, Down, Left, Right)
	directions := [][2]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}

	// Create a map to track visited nodes
	visited := make(map[[2]int]bool)
	visited[[2]int{start[0], start[1]}] = true

	for {
		maybePath := stack.pop()
		if maybePath == nil {
			break
		}
		path := *maybePath
		point := path[len(path)-1]
		x, y := point[0], point[1]
		println(x, y)

		// If the path is the end point, return the path it took to get here
		if x == end[0] && y == end[1] {
			return path
		}

		// Push the neighbors
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < len(maze[0]) && newY >= 0 && newY < len(maze) {
				if !visited[[2]int{newX, newY}] && maze[newY][newX] != 0 {
					newPath := append([][]int{}, path...)
					newPath = append(newPath, []int{newX, newY})
					println("pushing", newX, newY)
					println("current top of stack", stack.head)

					// Mark the node as visited before pushing it onto the stack
					visited[[2]int{newX, newY}] = true
					stack.push(newPath)
				}
			}
		}
	}
	return nil
}
