package main

import (
	"fmt"
	"strconv"
	"strings"

	setup "github.com/jambolo/advent-of-code-2015/internal/setup"
)

func lookAndSay(input string) string {
	var b strings.Builder
	for i := 0; i < len(input); {
		ch := input[i]
		count := 1
		for i+count < len(input) && input[i+count] == ch {
			count++
		}
		b.WriteString(strconv.Itoa(count))
		b.WriteByte(ch)
		i += count
	}
	return b.String()
}

func main() {
	day := 10

	// Grab the command line parameters (file path and part number)
	_, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	input := "3113322113"
	var iterations int
	if part == 1 {
		iterations = 40
	} else {
		iterations = 50
	}

	for i := 0; i < iterations; i++ {
		input = lookAndSay(input)
	}

	fmt.Printf("The length of the result after %d iterations is %d\n", iterations, len(input))
}
