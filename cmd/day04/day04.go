package main

import (
	"crypto/md5"
	"flag"
	"fmt"
)

func main() {
	part := flag.Int("part", 1, "which part of the challenge (1 or 2)")

	// 2. Parse the flags
	flag.Parse()

	// Validate part
	if *part != 1 && *part != 2 {
		fmt.Println("Invalid part specified. Must be 1 or 2.")
		return
	}

	// Print a banner showing the current day and if it is part 1 or part 2
	fmt.Printf("=== Day 4 - Part %d ===\n", *part)

	prefix := "iwrupvqb"

	// Part 1
	if *part == 1 {
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
	if *part == 2 {
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
