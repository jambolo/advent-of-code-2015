package main

import (
	"fmt"
	"log"
	"strings"

	load "github.com/jambolo/advent-of-code-2015/internal/load"
	setup "github.com/jambolo/advent-of-code-2015/internal/setup"
	utils "github.com/jambolo/advent-of-code-2015/internal/utils"
)

const size = 1000

type action int

const (
	on action = iota
	off
	toggle
)

type rect struct {
	x1, y1, x2, y2 int
}

// instruction data structure
type instruction struct {
	operation action
	extents   rect
}

// parseInstructions takes a slice of strings representing the instructions and returns a slice of parsed instructions.
func parseInstructions(lines []string) []instruction {
	var instructions []instruction
	for _, line := range lines {
		var instr instruction
		switch {
		case strings.HasPrefix(line, "turn on "):
			instr.operation = on
			line = line[8:]
		case strings.HasPrefix(line, "turn off "):
			instr.operation = off
			line = line[9:]
		case strings.HasPrefix(line, "toggle "):
			instr.operation = toggle
			line = line[7:]
		default:
			log.Fatalf("Invalid instruction: %s", line)
		}
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d through %d,%d", &x1, &y1, &x2, &y2)
		instr.extents = rect{x1, y1, x2, y2}
		instructions = append(instructions, instr)
	}
	return instructions
}

func main() {
	day := 6

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	instructions := parseInstructions(lines)

	// Part 1
	if part == 1 {
		// Create a 1000x1000 grid of booleans to represent the lights
		grid := make([]bool, size*size)

		// Apply each instruction to the grid
		for _, instr := range instructions {
			for x := instr.extents.x1; x <= instr.extents.x2; x++ {
				for y := instr.extents.y1; y <= instr.extents.y2; y++ {
					switch instr.operation {
					case on:
						grid[y*size+x] = true
					case off:
						grid[y*size+x] = false
					case toggle:
						grid[y*size+x] = !grid[y*size+x]
					}
				}
			}
		}

		// Count the number of lights that are on
		count := 0
		for i := range grid {
			if grid[i] {
				count++
			}
		}
		fmt.Printf("Number of lights on: %d\n", count)
	}

	// Part 2
	if part == 2 {
		// Create a 1000x1000 grid of booleans to represent the lights
		grid := make([]int, size*size)

		// Apply each instruction to the grid
		for _, instr := range instructions {
			for x := instr.extents.x1; x <= instr.extents.x2; x++ {
				for y := instr.extents.y1; y <= instr.extents.y2; y++ {
					switch instr.operation {
					case on:
						grid[y*size+x]++
					case off:
						if grid[y*size+x] > 0 {
							grid[y*size+x]--
						}
					case toggle:
						grid[y*size+x] += 2
					}
				}
			}
		}

		// Count the number of lights that are on
		brightness := utils.Sum(grid)
		fmt.Printf("Total brightness: %d\n", brightness)
	}
}
