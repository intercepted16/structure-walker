package Examples

import . "structure-walker/BreadthFirstSearch"

func removeOnes(maze [][]int) Maze {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == 1 {
				maze[i][j] = 0
			}
		}
	}
	return maze
}

// ExampleBreadthFirstSearch demonstrates a breadth-first search
// WIP: currently this is used to test the implementation
// The breadth first search demonstration is not implemented in C
func ExampleBreadthFirstSearch() {
	// Each row represents a row in the maze
	// Each column represents a column in the maze
	// So the first row with 0's means there is a wall at the top of the maze
	solvedMaze := Maze{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 2, 1, 0, 1, 3, 0},
		{0, 0, 1, 0, 1, 0, 0},
		{0, 1, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	// create the maze without the '1's so there is no path
	maze := removeOnes(solvedMaze)
	SolveMaze(maze)
}
