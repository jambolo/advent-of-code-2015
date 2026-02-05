package main

import (
	"flag"
	"fmt"

	load "github.com/jambolo/advent-of-code-2015/internal/common"
)

func main() {
	filePath := flag.String("file", "data/default.dat", "path to the unique data file")
	part := flag.Int("part", 1, "which part of the challenge (1 or 2)")

	// 2. Parse the flags
	flag.Parse()

	// Validate part
	if *part != 1 && *part != 2 {
		fmt.Println("Invalid part specified. Must be 1 or 2.")
		return
	}

	// Print a banner showing the current day and if it is part 1 or part 2
	fmt.Printf("=== Day 1 - Part %d ===\n", *part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.ReadLines(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Part 1
	if *part == 1 {
		for _, line := range lines {
			// Count the difference between '(' and ')' and print the result
			floor := 0
			for _, char := range line {
				switch char {
				case '(':
					floor++
				case ')':
					floor--
				}
			}
			fmt.Printf("Final floor: %d\n", floor)
		}
	}

	// Part 2
	if *part == 2 {
		for _, line := range lines {
			floor := 0
			// Find the position of the first character that causes the floor to go below 0
			position := 1
			for _, char := range line {
				switch char {
				case '(':
					floor++
				case ')':
					floor--
				}
				if floor < 0 {
					fmt.Printf("Position of first basement entry: %d\n", position)
					break
				}
				position++
			}
		}
	}
}
