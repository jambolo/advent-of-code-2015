package main

import (
	"fmt"
	"log"

	load "github.com/jambolo/advent-of-code-2015/internal/load"
	setup "github.com/jambolo/advent-of-code-2015/internal/setup"
)

func main() {
	day := 1

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	input, err := load.ReadAll(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	if part == 1 {
		// Count the difference between '(' and ')' and print the result
		floor := 0
		for _, char := range input {
			switch char {
			case '(':
				floor++
			case ')':
				floor--
			}
		}
		fmt.Printf("Final floor: %d\n", floor)
	}

	// Part 2
	if part == 2 {
		floor := 0
		// Find the position of the first character that causes the floor to go below 0
		for position, char := range input {
			switch char {
			case '(':
				floor++
			case ')':
				floor--
			}
			if floor < 0 {
				fmt.Printf("Position of first basement entry: %d\n", position+1)
				break
			}
		}
	}
}
