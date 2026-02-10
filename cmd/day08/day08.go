package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

func processNext(line string, i int) (rune, int) {
	char := line[i]
	var value rune
	if char == '\\' {
		next := line[i+1]
		switch next {
		case '\\', '"':
			value = rune(next)
			i += 2
		case 'x':
			c, err := strconv.ParseUint(line[i+2:i+4], 16, 8)
			if err != nil {
				log.Fatalf("Invalid hex escape sequence: %s", line[i+2:i+4])
			}
			value = rune(c)
			i += 4
		default:
			log.Fatalf("Unexpected escape sequence: \\%c", next)
		}
	} else {
		value = rune(char)
		i++
	}
	return value, i
}

func encode(char rune) []rune {
	switch char {
	case '\\':
		return []rune{'\\', '\\'}
	case '"':
		return []rune{'\\', '"'}
	default:
		return []rune{char}
	}
}

func main() {
	day := 8

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.Lines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	if part == 1 {
		var characters []rune
		totalLineLength := 0
		for _, line := range lines {
			length := len(line)
			totalLineLength += length
			i := 1             // Skip the first character, which is a double quote.
			for i < length-1 { // Skip the last character, which is a double quote.
				value, nextIndex := processNext(line, i)
				characters = append(characters, value)
				i = nextIndex
			}
			//		fmt.Printf("Line: %s, length: %d, Characters: %q, total length: %d\n", line, length, characters, totalLineLength) // Print the line and the characters parsed from it.
		}

		fmt.Printf("Result: %d\n", totalLineLength-len(characters))
	}

	if part == 2 {
		totalLineLength := 0
		var encoded []rune
		for _, line := range lines {
			totalLineLength += len(line)
			encoded = append(encoded, '"') // Add the starting double quote for the encoded string.
			for _, char := range line {
				encoded = append(encoded, encode(char)...)
			}
			encoded = append(encoded, '"') // Add the ending double quote for the encoded string.
		}

		fmt.Printf("Result: %d\n", len(encoded)-totalLineLength)
	}
}
