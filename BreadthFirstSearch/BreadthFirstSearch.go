package BreadthFirstSearch

import (
	"fmt"
	. "structure-walker/Maze"
)

// SolveMaze solves a maze through breadth-first search (BFS) and returns the path taken.
func SolveMaze(maze Maze) [][]int {
	start, end := FindStartAndEnd(maze)
	if start == nil || end == nil {
		fmt.Println("Start or end point not found.")
		return nil
	}

	// Create a queue
	queue := NewQueue()

	// Enqueue the start point with its path
	queue.Enqueue([][]int{start})

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
		path := queue.Dequeue()
		if path == nil {
			break
		}

		point := path[len(path)-1]
		x, y := point[0], point[1]

		// Check if the point is the end point
		if x == end[0] && y == end[1] {
			return path
		}

		// Mark the point as visited
		visited[[2]int{x, y}] = true

		// Enqueue the neighbors
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < len(maze[0]) && newY >= 0 && newY < len(maze) {
				if !visited[[2]int{newX, newY}] && maze[newY][newX] != 0 {
					newPath := append([][]int{}, path...)
					newPath = append(newPath, []int{newX, newY})
					queue.Enqueue(newPath)
				}
			}
		}
	}

	return nil
}
