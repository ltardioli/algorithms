package maze

import "fmt"

type Point struct {
	X, Y int
}

type Maze struct {
	building   [][]string
	Start, End Point
	Directions []Point
	wall       string
}

func (maze *Maze) New(start Point, end Point, building [][]string, wall string) {
	maze.Start = start
	maze.End = end
	maze.Directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	maze.building = building
	maze.wall = wall
}

func (maze *Maze) GetLocation(point Point) string {
	return maze.building[point.X][point.Y]
}

//This solution finds a path of a maze, not the best path of it.
func Walk(maze Maze, curr Point, seen *[][]bool, path *[]Point) bool {
	// Base cases
	// Off the grid
	if curr.X < 0 || curr.X > len(maze.building[0]) || curr.Y < 0 || curr.Y > len(maze.building[0]) {
		return false
	}

	// On a wall
	if maze.GetLocation(curr) == maze.wall {
		return false
	}

	// Got to the end
	if curr.X == maze.End.X && curr.Y == maze.End.Y {
		*path = append(*path, maze.End)
		fmt.Println(path)
		return true
	}

	// Seen
	if (*seen)[curr.X][curr.Y] {
		return false
	}

	// Recursion

	// Pre
	// Push the point
	(*seen)[curr.X][curr.Y] = true
	*path = append(*path, curr)

	// Recurse
	for _, p := range maze.Directions {
		newCurr := Point{curr.X + p.X, curr.Y + p.Y}
		if Walk(maze, newCurr, seen, path) {
			return true
		}
	}

	// Post
	// Pop the point
	*path = (*path)[:len(*path)-1]

	return false
}

func Solve(maze Maze, startPoint Point, endPoint Point) []Point {
	seen := make([][]bool, len(maze.building))
	for i := range seen {
		seen[i] = make([]bool, len(maze.building))
	}

	path := make([]Point, 0, len(maze.building))

	found := Walk(maze, startPoint, &seen, &path)
	fmt.Println(found)
	return path
}
