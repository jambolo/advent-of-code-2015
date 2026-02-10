package main

import (
	"fmt"
	"log"
	"math"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
	"github.com/jambolo/advent-of-code-2015/internal/utils"
)

type stringSet map[string]struct{}

func main() {
	day := 9

	filePath, part := setup.Parameters(day)
	setup.Banner(day, part)

	lines, err := load.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse distances and collect unique city names.
	distances := make(map[[2]string]int)
	citySet := make(stringSet)
	for _, line := range lines {
		var city1, city2 string
		var distance int
		fmt.Sscanf(line, "%s to %s = %d", &city1, &city2, &distance)
		distances[[2]string{city1, city2}] = distance
		distances[[2]string{city2, city1}] = distance
		citySet[city1] = struct{}{}
		citySet[city2] = struct{}{}
	}
	var cities []string
	for city := range citySet {
		cities = append(cities, city)
	}

	allRoutes := utils.Permutations(len(cities), len(cities))
	minDistance := math.MaxInt
	maxDistance := 0
	for _, route := range allRoutes {
		var totalDistance int
		for i := range route[:len(route)-1] {
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
