package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
	"github.com/jambolo/advent-of-code-2015/internal/utils"
)

// isFlying returns true if the reindeer is flying at the given time, false if it is resting.
func isFlying(time int, flyTime int, cycleTime int) bool {
	return (time-1)%cycleTime < flyTime
}

type reindeer struct {
	speed, flyTime, restTime, cycleTime, cycleDistance int
}

func main() {
	day := 14

	// Grab the command line parameters (file path and part number)
	pathName, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	lines, err := load.ReadLines(pathName)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

	reindeers := make(map[string]reindeer)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 5 {
			log.Fatalf("unexpected line format: %s", line)
		}

		name := matches[1]
		speed, err := strconv.Atoi(matches[2])
		if err != nil {
			log.Fatalf("invalid speed: %s", matches[2])
		}
		flyTime, err := strconv.Atoi(matches[3])
		if err != nil {
			log.Fatalf("invalid fly time: %s", matches[3])
		}
		restTime, err := strconv.Atoi(matches[4])
		if err != nil {
			log.Fatalf("invalid rest time: %s", matches[4])
		}

		cycleTime := flyTime + restTime
		cycleDistance := speed * flyTime

		reindeers[name] = reindeer{
			speed:         speed,
			flyTime:       flyTime,
			restTime:      restTime,
			cycleTime:     cycleTime,
			cycleDistance: cycleDistance,
		}
	}

	totalTime := 2503

	if part == 1 {
		var distances []int
		for _, r := range reindeers {
			cycles := totalTime / r.cycleTime
			timeInLastCycle := totalTime % r.cycleTime
			distanceInLastCycle := min(r.flyTime, timeInLastCycle) * r.speed // After last full cycle
			distances = append(distances, cycles*r.cycleDistance+distanceInLastCycle)
		}

		fmt.Printf("Winner distance: %d\n", utils.SliceMax(distances))
	}

	if part == 2 {
		distances := make(map[string]int)
		for name := range reindeers {
			distances[name] = 0
		}

		points := make(map[string]int)
		for name := range reindeers {
			points[name] = 0
		}

		for t := 1; t <= totalTime; t++ {
			// Update the distances for each reindeer and track the leading distance
			leadingDistance := 0
			for name, r := range reindeers {
				if isFlying(t, r.flyTime, r.cycleTime) {
					distances[name] = distances[name] + r.speed
				}
				leadingDistance = max(leadingDistance, distances[name])
			}

			// Award points to the reindeer(s) in the lead
			for name, distance := range distances {
				if distance >= leadingDistance {
					points[name] = points[name] + 1
				}
			}
		}

		// Find the reindeer with the most points
		mostPoints := 0
		for _, p := range points {
			mostPoints = max(mostPoints, p)
		}

		fmt.Printf("Winner points: %d\n", mostPoints)
	}
}
