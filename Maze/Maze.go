// Package Maze includes useful utilities when solving a maze, typically through algorithms such as a BFS or DFS.
// It includes types and functions to manipulate a maze.
package Maze

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

// Maze represents a maze with a start and end point.
type Maze [][]int // 2D array of numbers

func contains(path [][]int, point []int) bool {
	for _, p := range path {
		if p[0] == point[0] && p[1] == point[1] {
			return true
		}
	}
	return false
}

// DrawMaze prints the maze to the console
func DrawMaze(maze Maze, path [][]int) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if contains(path, []int{j, i}) {
				// if it's the start point, print it red
				if maze[i][j] == 2 {
					fmt.Print(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(fmt.Sprintf("%d", maze[i][j])), " ")
					continue
				}
				// if it's the end point, print it blue
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

// FindStartAndEnd identifies the start and end points in the maze.
func FindStartAndEnd(maze Maze) ([]int, []int) {
	var start []int
	var end []int

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == 2 {
				start = []int{x, y}
			} else if maze[y][x] == 3 {
				end = []int{x, y}
			}
		}
	}

	return start, end
}
