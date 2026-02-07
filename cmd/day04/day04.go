package main

import (
	"crypto/md5"
	"fmt"

	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

func main() {
	day := 4

	// Grab the command line parameters (file path and part number)
	_, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	prefix := "iwrupvqb"

	// Part 1
	if part == 1 {
		for i := 0; ; i++ {
			input := fmt.Sprintf("%s%d", prefix, i)
			hash := md5.Sum([]byte(input))
			// Each byte is two hex characters, so we check the first three bytes for 5 leading zeroes (00000)
			if hash[0] == 0 && hash[1] == 0 && hash[2] < 16 {
				fmt.Printf("Answer: %d\n", i)
				break
			}
		}
	}
	// Part 2
	if part == 2 {
		for i := 0; ; i++ {
			input := fmt.Sprintf("%s%d", prefix, i)
			hash := md5.Sum([]byte(input))
			// Each byte is two hex characters, so we check the first three bytes for 6 leading zeroes (000000)
			if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
				fmt.Printf("Answer: %d\n", i)
				break
			}
		}
	}
}
