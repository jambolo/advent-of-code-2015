package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
	"github.com/jambolo/advent-of-code-2015/internal/utils"
)

type relationshipMap map[string]map[string]int

func computeGroupHappiness(p []int, people []string, relationships relationshipMap) int {
	happiness := 0
	n := len(p)
	for i := range n {
		person1 := people[p[i]]
		person2 := people[p[(i+1)%n]]
		h0 := relationships[person1][person2]
		h1 := relationships[person2][person1]
		happiness += h0 + h1
	}
	return happiness
}

func main() {
	day := 13

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Build relationship map
	relationships := make(relationshipMap)
	for _, line := range lines {
		var person1, person2 string
		var happiness int
		var gainOrLose string

		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s.", &person1, &gainOrLose, &happiness, &person2)
		person2 = strings.TrimSuffix(person2, ".") // Annoying

		if gainOrLose == "lose" {
			happiness = -happiness
		}

		if relationships[person1] == nil {
			relationships[person1] = make(map[string]int)
		}
		relationships[person1][person2] = happiness
	}

	// Create a list of people
	var people []string
	for person := range relationships {
		people = append(people, person)
	}

	if part == 2 {
		// Add "me" to the list of people and relationships
		people = append(people, "me")
		for _, person := range people {
			if person != "me" {
				relationships[person]["me"] = 0
				if relationships["me"] == nil {
					relationships["me"] = make(map[string]int)
				}
				relationships["me"][person] = 0
			}
		}
	}

	// Generate all permutations of people
	numberOfPeople := len(relationships)
	permutations := utils.Permutations(numberOfPeople)

	// Score each permutation and find the maximum
	maxHappiness := math.MinInt
	for _, p := range permutations {
		happiness := computeGroupHappiness(p, people, relationships)
		if happiness > maxHappiness {
			maxHappiness = happiness
		}
	}

	fmt.Printf("Maximum happiness is %d\n", maxHappiness)
}
