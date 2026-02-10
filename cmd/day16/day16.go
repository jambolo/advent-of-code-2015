package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

var mfcsam = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

type sue map[string]int

func matchesSuePart1(sue sue) bool {
	for name, value := range sue {
		if mfcsam[name] != value {
			return false
		}
	}
	return true
}

func matchesSuePart2(sue sue) bool {
	for name, value := range sue {
		switch name {
		case "cats", "trees":
			if mfcsam[name] >= value {
				return false
			}
		case "pomeranians", "goldfish":
			if mfcsam[name] <= value {
				return false
			}
		default:
			if mfcsam[name] != value {
				return false
			}
		}
	}
	return true
}

func main() {
	day := 16

	path, part := setup.Parameters(day)
	setup.Banner(day, part)

	lines, err := load.Lines(path)
	if err != nil {
		log.Fatal(err)
	}

	sues := make([]sue, 500)
	re := regexp.MustCompile(`(\w+): (\d+)`)

	for _, line := range lines {
		var s int
		fmt.Sscanf(line, "Sue %d:", &s)

		if s < 1 || s > 500 {
			log.Fatalf("invalid sue number: %d", s)
		}

		properties := re.FindAllStringSubmatch(line, -1)

		i := s - 1
		if sues[i] == nil {
			sues[i] = make(sue)
		}
		for _, property := range properties {
			name := property[1]
			value, _ := strconv.Atoi(property[2])
			sues[i][name] = value
		}
	}

	if part == 1 {
		for i, sue := range sues {
			if matchesSuePart1(sue) {
				fmt.Printf("Sue %d matches\n", i+1)
				break
			}
		}
	}

	if part == 2 {
		for i, sue := range sues {
			if matchesSuePart2(sue) {
				fmt.Printf("Sue %d matches\n", i+1)
				break
			}
		}
	}
}
