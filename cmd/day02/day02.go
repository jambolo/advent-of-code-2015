package main

import (
	"fmt"
	"log"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

func main() {
	day := 2

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.Lines(filePath)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}
		boxes = append(boxes, box)
	}

	// Part 1
	if part == 1 {
		total := 0
		for _, box := range boxes {
			// Calculate the surface area of the box and add the area of the smallest side as extra
			lw := box.length * box.width
			wh := box.width * box.height
			hl := box.height * box.length
			total += 2*(lw+wh+hl) + min(lw, wh, hl)
		}
		fmt.Printf("Total : %d\n", total)
	}

	// Part 2
	if part == 2 {
		total := 0
		for _, box := range boxes {
			// Smallest perimeter + volume for the bow
			perimeter := 2 * min(box.length+box.width, box.width+box.height, box.height+box.length)
			volume := box.length * box.width * box.height
			total += perimeter + volume
		}
		fmt.Printf("Total : %d\n", total)
	}
}
