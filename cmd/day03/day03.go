package main

import (
	"flag"
	"fmt"

	load "github.com/jambolo/advent-of-code-2015/internal/common"
)

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
	// move updates x and y based on the direction character and returns the new coordinates
	filePath := flag.String("file", "data/day03/day03-input.txt", "path to the data file")
	part := flag.Int("part", 1, "which part of the challenge (1 or 2)")

	// 2. Parse the flags
	flag.Parse()

	// Validate part
	if *part != 1 && *part != 2 {
		fmt.Println("Invalid part specified. Must be 1 or 2.")
		return
	}

	// Print a banner showing the current day and if it is part 1 or part 2
	fmt.Printf("=== Day 3 - Part %d ===\n", *part)

	// Load the data from the specified file. Abort on error.
	input, err := load.ReadAll(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Part 1
	if *part == 1 {
		x := 0
		y := 0
		visited := make(map[string]bool)
		visited["0,0"] = true

		for _, char := range input {
			x, y = move(x, y, char)
			visited[fmt.Sprintf("%d,%d", x, y)] = true
		}
		fmt.Printf("Houses visited: %d\n", len(visited))
	}

	if *part == 2 {
		santa_x := 0
		santa_y := 0
		robo_x := 0
		robo_y := 0
		visited := make(map[string]bool)
		visited["0,0"] = true

		for i, char := range input {
			if i%2 == 0 {
				santa_x, santa_y = move(santa_x, santa_y, char)
				visited[fmt.Sprintf("%d,%d", santa_x, santa_y)] = true
			} else {
				robo_x, robo_y = move(robo_x, robo_y, char)
				visited[fmt.Sprintf("%d,%d", robo_x, robo_y)] = true
			}
		}
		fmt.Printf("Houses visited: %d\n", len(visited))
	}

}
