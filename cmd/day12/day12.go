package main

import (
	"fmt"
	"log"

	load "github.com/jambolo/advent-of-code-2015/internal/load"
	setup "github.com/jambolo/advent-of-code-2015/internal/setup"
)

func sumAllNumbers(data any) int {
	sum := 0

	switch v := data.(type) {
	case float64:
		sum += int(v)
	case []any:
		for _, item := range v {
			sum += sumAllNumbers(item)
		}
	case map[string]any:
		for _, value := range v {
			sum += sumAllNumbers(value)
		}
	}

	return sum
}

// containsItemWithRed checks if the provided map contains a "red" value.
func containsItemWithRed(data map[string]any) bool {
	for _, item := range data {
		if str, ok := item.(string); ok && str == "red" {
			return true
		}
	}
	return false
}

func sumAllNumbersWithoutRed(data any) int {
	sum := 0

	switch v := data.(type) {
	case float64:
		sum += int(v)
	case []any:
		for _, item := range v {
			sum += sumAllNumbersWithoutRed(item)
		}
	case map[string]any:
		// Recurse only if the map does not contain a "red" value
		if !containsItemWithRed(v) {
			for _, value := range v {
				sum += sumAllNumbersWithoutRed(value)
			}
		}
	}
	return sum
}

func main() {
	day := 12

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	var input any
	err := load.Json(filePath, &input)
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	if part == 1 {
		sum := sumAllNumbers(input)
		fmt.Printf("Sum of all numbers: %d\n", sum)
	}

	// Part 2
	if part == 2 {
		sum := sumAllNumbersWithoutRed(input)
		fmt.Printf("Sum of all numbers without red: %d\n", sum)
	}
}
