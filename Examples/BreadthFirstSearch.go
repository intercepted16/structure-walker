package Examples

import (
	. "structure-walker/BreadthFirstSearch"
	. "structure-walker/Maze"
)

// ExampleBreadthFirstSearch demonstrates a breadth-first search
// WIP: currently this is used to test the implementation
// The breadth first search demonstration is not implemented in C
func ExampleBreadthFirstSearch() {
	// Each row represents a row in the maze
	// Each column represents a column in the maze
	// So the first row with 0's means there is a wall at the top of the maze
	maze := Maze{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 2, 1, 0, 1, 3, 0},
		{0, 0, 1, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	path := SolveMaze(maze)
	DrawMaze(maze, path)
}
