package main

import (
	"flag"
	"fmt"

	load "github.com/jambolo/advent-of-code-2015/internal/common"
)

func main() {
	filePath := flag.String("file", "data/day02/day02-input.txt", "path to the data file")
	part := flag.Int("part", 1, "which part of the challenge (1 or 2)")

	// 2. Parse the flags
	flag.Parse()

	// Validate part
	if *part != 1 && *part != 2 {
		fmt.Println("Invalid part specified. Must be 1 or 2.")
		return
	}

	// Print a banner showing the current day and if it is part 1 or part 2
	fmt.Printf("=== Day 2 - Part %d ===\n", *part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.ReadLines(*filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// For each line, parse the dimensions of the box into a vector of 3-element integer tuples
	type Box struct {
		length int
		width  int
		height int
	}

	var boxes []Box
	for _, line := range lines {
		var box Box
		_, err := fmt.Sscanf(line, "%dx%dx%d", &box.length, &box.width, &box.height)
		if err != nil {
			fmt.Printf("Error parsing line '%s': %v\n", line, err)
			return
		}
		boxes = append(boxes, box)
	}

	// Part 1
	if *part == 1 {
		total := 0
		for _, box := range boxes {
			// Calculate the surface area of the box and add the area of the smallest side as extra
			lw := box.length * box.width
			wh := box.width * box.height
			hl := box.height * box.length
			surfaceArea := 2*lw + 2*wh + 2*hl
			smallestSide := lw
			if wh < smallestSide {
				smallestSide = wh
			}
			if hl < smallestSide {
				smallestSide = hl
			}
			total += surfaceArea + smallestSide
		}
		fmt.Printf("Total : %d\n", total)
	}

	// Part 2
	if *part == 2 {
		total := 0
		for _, box := range boxes {
			// Calculate the ribbon needed to wrap the box, which is the smallest perimeter of any one face
			lw_perimeter := 2 * (box.length + box.width)
			wh_perimeter := 2 * (box.width + box.height)
			hl_perimeter := 2 * (box.height + box.length)
			smallestPerimeter := lw_perimeter
			if wh_perimeter < smallestPerimeter {
				smallestPerimeter = wh_perimeter
			}
			if hl_perimeter < smallestPerimeter {
				smallestPerimeter = hl_perimeter
			}
			// Add the volume of the box for the bow
			volume := box.length * box.width * box.height
			total += smallestPerimeter + volume
		}
		fmt.Printf("Total : %d\n", total)
	}
}
