package main

import (
	"fmt"
	"log"
	"strings"

	load "github.com/jambolo/advent-of-code-2015/internal/load"
	setup "github.com/jambolo/advent-of-code-2015/internal/setup"
)

// hasThreeVowels returns true if the string has at least three vowels (aeiou), false otherwise
func hasThreeVowels(s string) bool {
	count := 0
	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			count++
			if count >= 3 {
				return true
			}
		}
	}
	return false
}

// hasDoubleLetter returns true if the string contains at least one letter that appears twice in a row, false otherwise
func hasDoubleLetter(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}

// hasNoBadWords returns true if the string does not contain any of the bad words (ab, cd, pq, xy), false otherwise
func hasNoBadWords(s string) bool {
	badWords := []string{"ab", "cd", "pq", "xy"}
	for _, bad := range badWords {
		if strings.Contains(s, bad) {
			return false
		}
	}
	return true
}

// hasRepeatedPair returns true if the string contains a pair of any two letters that appears at least twice in the string without overlapping.
func hasRepeatedPair(s string) bool {
	for i := range len(s)-1 {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}
	return false
}

// hasSplitPair returns true if the string contains a pair of any two letters with a letter in between them.
func hasSplitPair(s string) bool {
	for i := range len(s)-2 {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func main() {
	day := 5

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	if part == 1 {
		niceCount := 0
		for _, line := range lines {
			if hasThreeVowels(line) && hasDoubleLetter(line) && hasNoBadWords(line) {
				niceCount++
			}
		}
		fmt.Printf("Nice strings: %d\n", niceCount)
	}

	if part == 2 {
		niceCount := 0
		for _, line := range lines {
			if hasRepeatedPair(line) && hasSplitPair(line) {
				niceCount++
			}
		}
		fmt.Printf("Nice strings: %d\n", niceCount)
	}
}
