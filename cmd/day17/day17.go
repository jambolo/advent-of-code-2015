package main

import (
	"fmt"
	"log"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
	"github.com/jambolo/advent-of-code-2015/internal/utils"
)

func gather(x []int, containers []int) []int {
	y := make([]int, len(x))
	for i, v := range x {
		y[i] = containers[v]
	}
	return y
}

func main() {
	day := 17

	path, part := setup.Parameters(day)
	setup.Banner(day, part)

	lines, err := load.Lines(path)
	if err != nil {
		log.Fatal(err)
	}

	containers := make([]int, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "%d", &containers[i])
	}

	if part == 1 {
		count := 0
		for i := 1; i <= len(containers); i++ {
			combinations := utils.Combinations(len(containers), i)
			for _, p := range combinations {
				if utils.SliceSum(gather(p, containers)) == 150 {
					count++
				}
			}
		}
		fmt.Printf("There are %d combinations of containers that can hold 150 liters.\n", count)
	}

	if part == 2 {
		for i := 1; i <= len(containers); i++ {
			count := 0
			combinations := utils.Combinations(len(containers), i)
			for _, p := range combinations {
				if utils.SliceSum(gather(p, containers)) == 150 {
					count++
				}
			}
			if count > 0 {
				fmt.Printf("Count at minimum: %d.\n", count)
				break
			}
		}
	}
}
