package main

import (
	"fmt"
	"strings"

	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

// hasStraight returns true if the password includes an increasing straight of at least three letters.
func hasStraight(password string) bool {
	for i := range len(password)-2 {
		if password[i]+1 == password[i+1] && password[i]+2 == password[i+2] {
			return true
		}
	}
	return false
}

// hasNoInvalidCharacters returns true if the password does not contain the letters i, o, or l.
func hasNoInvalidCharacters(password string) bool {
	return !strings.ContainsAny(password, "iol")
}

// hasTwoPairs returns true if the password contains at least two different, non-overlapping pairs of letters.
func hasTwoPairs(password string) bool {
	for i := range len(password)-3 {
		if password[i] == password[i+1] {
			// Found a pair, now look for another pair
			for j := i + 2; j < len(password)-1; j++ {
				if password[j] == password[j+1] && password[j] != password[i] {
					return true
				}
			}
		}
	}
	return false
}

func isValid(password string) bool {
	return hasStraight(password) && hasNoInvalidCharacters(password) && hasTwoPairs(password)
}

func incrementPassword(password string) string {
	runes := []rune(password)
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == 'z' {
			runes[i] = 'a'
		} else {
			runes[i]++
			break
		}
	}
	return string(runes)
}

func main() {
	day := 11

	// Grab the command line parameters (file path and part number)
	_, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	password := "cqjxjnds"

	password = incrementPassword(password)
	for !isValid(password) {
		password = incrementPassword(password)
	}

	if part == 1 {
		fmt.Printf("Next password: %s\n", password)
	}

	if part == 2 {
		password = incrementPassword(password)
		for !isValid(password) {
			password = incrementPassword(password)
		}
		fmt.Printf("New password: %s\n", password)
	}
}
