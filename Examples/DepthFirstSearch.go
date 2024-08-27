package Examples

import . "structure-walker/DepthFirstSearch"
import . "structure-walker/Maze"

func ExampleDepthFirstSearch() {
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
	solvedPath := SolveMaze(maze)
	DrawMaze(maze, solvedPath)
}
