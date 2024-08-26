package Examples

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	. "structure-walker/BreadthFirstSearch"
)

func contains(path [][]int, point []int) bool {
	for _, p := range path {
		if p[0] == point[0] && p[1] == point[1] {
			return true
		}
	}
	return false
}

// drawMaze prints the maze to the console
func drawMaze(maze Maze, path [][]int) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if contains(path, []int{j, i}) {
				// if its the start point, print it red
				if maze[i][j] == 2 {
					fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(fmt.Sprintf("%d", maze[i][j])), " ")
					continue
				}
				// if its the end point, print it blue
				if maze[i][j] == 3 {
					fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#0000FF")).Render(fmt.Sprintf("%d", maze[i][j])), " ")
					continue
				}
				// print it green
				fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render(fmt.Sprintf("%d", maze[i][j])), " ")
				continue
			}
			fmt.Print(maze[i][j], " ")
		}
		fmt.Println()
	}
}

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
	drawMaze(maze, path)
}
