package main

import (
	"fmt"
	"main/algo/maze"
)

func main() {

	building := [][]string{
		{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
		{"S", " ", " ", " ", "#", " ", " ", " ", " ", "#"},
		{"#", " ", "#", " ", "#", " ", "#", "#", " ", "#"},
		{"#", " ", "#", " ", " ", " ", "#", " ", " ", "#"},
		{"#", " ", "#", "#", "#", " ", "#", " ", "#", "#"},
		{"#", " ", " ", " ", "#", " ", " ", " ", "#", "#"},
		{"#", "#", "#", " ", "#", "#", "#", " ", " ", "#"},
		{"#", " ", " ", " ", " ", " ", "#", "#", " ", "#"},
		{"#", " ", "#", "#", "#", " ", " ", "#", " ", "E"},
		{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
	}

	startPoint := maze.Point{X: 1, Y: 0}
	endPoint := maze.Point{X: 8, Y: 9}

	var mazeSolve maze.Maze
	mazeSolve.New(startPoint, endPoint, building, "#")

	solution := maze.Solve(mazeSolve, startPoint, endPoint)

	for _, p := range solution {
		building[p.X][p.Y] = "."
	}

	for _, row := range building {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
