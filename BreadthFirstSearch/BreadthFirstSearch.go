package BreadthFirstSearch

import "fmt"

// Maze represents a maze with a start and end point.
type Maze [][]int // 2D array of numbers

// FindStartAndEnd identifies the start and end points in the maze.
func FindStartAndEnd(maze Maze) ([]int, []int) {
	var start []int
	var end []int

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == 2 {
				start = []int{i, j}
			} else if maze[i][j] == 3 {
				end = []int{i, j}
			}
		}
	}

	return start, end
}

// Queue represents a queue data structure.
type Queue struct {
	items [][]int
}

// NewQueue creates a new Queue.
func NewQueue() *Queue {
	return &Queue{items: [][]int{}}
}

// Enqueue adds an item to the queue.
func (q *Queue) Enqueue(item []int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item from the queue.
func (q *Queue) Dequeue() []int {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// SolveMaze solves a maze through breadth-first search (BFS).
func SolveMaze(maze Maze) Maze {
	start, end := FindStartAndEnd(maze)
	if start == nil || end == nil {
		fmt.Println("Start or end point not found.")
		return maze
	}

	fmt.Println("Start:", start)
	fmt.Println("End:", end)

	// Create a queue
	queue := NewQueue() // Directly create a new Queue

	// Enqueue the start point with its path
	queue.Enqueue(start)

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
	maze[start[0]][start[1]] = 4 // Mark the start point as part of the path

	for {
		point := queue.Dequeue()
		if point == nil {
			break
		}

		x, y := point[0], point[1]
		println("Visiting", x, y)

		// Check if the point is the end point
		if x == end[0] && y == end[1] {
			fmt.Println("End found")
			maze[x][y] = 4 // Mark the endpoint as part of the path
			break
		}

		// Mark the point as visited
		visited[[2]int{x, y}] = true
		maze[x][y] = 4 // Mark as part of the path

		// Check if the point is a wall
		if maze[x][y] == 0 {
			continue
		}

		// If the neighbor is within the bounds of the maze and has not already been visited
		// Loop over each possible direction to visit
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			println("Checking neighbor", newX, newY)
			if newX >= 0 && newX < len(maze) && newY >= 0 && newY < len(maze[0]) {
				println("Neighbor is within bounds")
				if !visited[[2]int{newX, newY}] && maze[newX][newY] != 0 {
					// Enqueue the neighbor
					println("Enqueueing neighbor", newX, newY)
					queue.Enqueue([]int{newX, newY})
				}
			}
		}
	}

	return maze
}
