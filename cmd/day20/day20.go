package main

import (
	"fmt"

	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

func main() {
	day := 20

	_, part := setup.Parameters(day)
	setup.Banner(day, part)

	maxPresents := 29000000

	if part == 1 {
		maxVisits := (maxPresents + 10 - 1) / 10

		sieve := make([]int, maxVisits+1)
		for i := 1; i <= maxVisits; i++ {
			for j := i; j <= maxVisits; j += i {
				sieve[j] += i
			}
		}

		// Find the first house with at least 290000 maxVisits visits.
		for i := 1; i <= maxVisits; i++ {
			if sieve[i] >= maxVisits {
				fmt.Printf("The first house to get at least %d visits is %d\n", maxVisits, i)
				break
			}
		}
	}

	if part == 2 {
		maxVisits := (maxPresents + 11 - 1) / 11

		sieve := make([]int, maxVisits+1)
		for i := 1; i <= maxVisits; i++ {
			for j := i; j <= min(maxVisits, i*50); j += i {
				sieve[j] += i
			}
		}

		// Find the first house with at least 290000 maxVisits visits.
		for i := 1; i <= maxVisits; i++ {
			if sieve[i] >= maxVisits {
				fmt.Printf("The first house to get at least %d visits is %d\n", maxVisits, i)
				break
			}
		}
	}
}
