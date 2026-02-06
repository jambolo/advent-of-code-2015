package setup

import (
	"flag"
	"fmt"
	"log"
)

// Parameters parses command-line flags and returns the input file name and part number.
func Parameters(day int) (path string, part int) {
	defaultPath := fmt.Sprintf("data/day%02d/day%02d-input.txt", day, day)
	pathFlag := flag.String("file", defaultPath, "Path to the input file")
	partFlag := flag.Int("part", 1, "Part number (1 or 2)")
	flag.Parse()

	if *partFlag != 1 && *partFlag != 2 {
		log.Fatal("Invalid part specified. Must be 1 or 2.")
	}
	return *pathFlag, *partFlag
}

// Banner prints a banner showing the current day and part.
func Banner(day int, part int) {
	fmt.Printf("=== Day %d - Part %d ===\n", day, part)
}
