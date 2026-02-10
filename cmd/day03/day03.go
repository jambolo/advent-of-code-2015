package main

import (
	"fmt"
	"log"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

type point struct{ x, y int }
type pointSet map[point]struct{}

// move updates x and y based on the direction character and returns the new coordinates
func move(x, y int, char rune) (int, int) {
	switch char {
	case '^':
		y++
	case 'v':
		y--
	case '>':
		x++
	case '<':
		x--
	}
	return x, y
}

func main() {
	day := 3

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	input, err := load.All(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	if part == 1 {
		x := 0
		y := 0
		visited := pointSet{{0, 0}: {}}

		for _, char := range input {
			x, y = move(x, y, char)
			visited[point{x, y}] = struct{}{}
		}
		fmt.Printf("Houses visited: %d\n", len(visited))
	}

	if part == 2 {
		santaX := 0
		santaY := 0
		roboX := 0
		roboY := 0
		visited := pointSet{{0, 0}: {}}

		for i, char := range input {
			if i%2 == 0 {
				santaX, santaY = move(santaX, santaY, char)
				visited[point{santaX, santaY}] = struct{}{}
			} else {
				roboX, roboY = move(roboX, roboY, char)
				visited[point{roboX, roboY}] = struct{}{}
			}
		}
		fmt.Printf("Houses visited: %d\n", len(visited))
	}

}
