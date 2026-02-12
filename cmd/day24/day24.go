package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
	"github.com/jambolo/advent-of-code-2015/internal/utils"
)

func main() {
	day := 24

	file, part := setup.Parameters(day)
	setup.Banner(day, part)

	lines, err := load.Lines(file)
	if err != nil {
		log.Fatal(err)
	}

	packages := make([]int, len(lines))

	for i, line := range lines {
		packages[i], err = strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
	}

	var groupWeight int
	if part == 1 {
		groupWeight = utils.SliceSum(packages) / 3
	}
	if part == 2 {
		groupWeight = utils.SliceSum(packages) / 4
	}

	groups := groupRecursive(packages, groupWeight)

	minGroupSize := len(groups[0])
	minEntanglement := utils.SliceProduct(utils.Gather(groups[0], packages))

	for i := 1; i < len(groups); i++ {
		groupSize := len(groups[i])
		if groupSize <= minGroupSize {
			entanglement := utils.SliceProduct(utils.Gather(groups[i], packages))
			if groupSize < minGroupSize {
				minGroupSize = groupSize
				minEntanglement = entanglement
			} else if entanglement < minEntanglement {
				minEntanglement = entanglement
			}
		}
	}
	fmt.Printf("Result: %d\n", minEntanglement)
}

func groupRecursive(packages []int, weight int) [][]int {
	group := make([][]int, 0)
	for i := range packages {
		if packages[i] == weight {
			group = append(group, []int{i})
		} else if packages[i] < weight {
			subgroups := groupRecursive(packages[i+1:], weight-packages[i])
			for j := range subgroups {
				for k := range subgroups[j] {
					subgroups[j][k] += i + 1
				}
				subgroups[j] = append(subgroups[j], i)
				group = append(group, subgroups[j])
			}
		}
	}
	return group
}
