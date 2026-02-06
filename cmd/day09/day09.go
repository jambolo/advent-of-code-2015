package main

import (
	"fmt"
	"log"
	"math"

	load "github.com/jambolo/advent-of-code-2015/internal/load"
	setup "github.com/jambolo/advent-of-code-2015/internal/setup"
)

// Returns all permutations of the integers 0 to n-1.
func permutations(n int) [][]int {
	if n == 0 {
		return [][]int{{}}
	}
	var result [][]int
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	var generate func(int)
	generate = func(k int) {
		if k == 1 {
			tmp := make([]int, n)
			copy(tmp, perm)
			result = append(result, tmp)
			return
		}
		for i := 0; i < k; i++ {
			generate(k - 1)
			if k%2 == 0 {
				perm[i], perm[k-1] = perm[k-1], perm[i]
			} else {
				perm[0], perm[k-1] = perm[k-1], perm[0]
			}
		}
	}
	generate(n)
	return result
}

func appendCity(cities []string, city string) []string {
	for _, v := range cities {
		if v == city {
			return cities
		}
	}
	return append(cities, city)
}

func main() {
	day := 9

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a map of distances and a list of city names.
	distances := make(map[[2]string]int)
	var cities []string
	for _, line := range lines {
		var city1, city2 string
		var distance int
		_, err := fmt.Sscanf(line, "%s to %s = %d", &city1, &city2, &distance)
		if err != nil {
			log.Fatal(err)
		}
		distances[[2]string{city1, city2}] = distance
		distances[[2]string{city2, city1}] = distance
		cities = appendCity(cities, city1)
		cities = appendCity(cities, city2)
	}

	allRoutes := permutations(len(cities))
	minDistance := math.MaxInt
	maxDistance := 0
	for _, route := range allRoutes {
		var totalDistance int
		for i := 0; i < len(route)-1; i++ {
			city0 := cities[route[i]]
			city1 := cities[route[i+1]]
			totalDistance += distances[[2]string{city0, city1}]
		}
		minDistance = min(minDistance, totalDistance)
		maxDistance = max(maxDistance, totalDistance)
	}

	if part == 1 {
		fmt.Printf("The shortest route has distance %d.\n", minDistance)
	} else {
		fmt.Printf("The longest route has distance %d.\n", maxDistance)
	}
}
