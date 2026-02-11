package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

func main() {
	day := 19

	path, part := setup.Parameters(day)
	setup.Banner(day, part)

	lines, err := load.Lines(path)
	if err != nil {
		log.Fatal(err)
	}

	// Get replacements.
	replacements := make(map[string][]string)
	for i := range lines {
		if lines[i] == "" {
			lines = lines[i+1:]
			break
		}
		var from, to string
		parts := strings.Fields(lines[i])
		if len(parts) != 3 || parts[1] != "=>" {
			log.Fatalf("invalid replacement: %s", lines[i])
		}
		from, to = parts[0], parts[2]
		replacements[from] = append(replacements[from], to)
	}

	// Get the molecule.
	molecule := lines[len(lines)-1]

	replaced := make(map[string]struct{})
	for i := range molecule {
		prefix := molecule[:i]
		remainder := molecule[i:]
		for from, to := range replacements {
			if strings.HasPrefix(remainder, from) {
				for _, t := range to {
					newMolecule := prefix + t + remainder[len(from):]
					replaced[newMolecule] = struct{}{}
				}
			}
		}
	}

	fmt.Printf("Results: %d.\n", len(replaced))
}
