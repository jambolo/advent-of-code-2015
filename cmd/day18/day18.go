package main

import (
	"fmt"
	"log"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

func neighborsCount(m [][]byte, x, y int) int {
	width := len(m[0])
	height := len(m)
	count := 0
	for i := -1; i <= 1; i++ {
		if y+i >= 0 && y+i < height {
			for j := -1; j <= 1; j++ {
				if j == 0 && i == 0 {
					continue
				}
				if x+j >= 0 && x+j < width {
					if m[y+i][x+j] == '#' {
						count++
					}
				}
			}
		}
	}
	return count
}

func step(m [][]byte) [][]byte {
	width := len(m[0])
	height := len(m)
	newMap := make([][]byte, height)
	for i := range m {
		newMap[i] = make([]byte, width)
		for j := range m[i] {
			count := neighborsCount(m, j, i)
			if m[i][j] == '#' {
				if count == 2 || count == 3 {
					newMap[i][j] = '#'
				} else {
					newMap[i][j] = '.'
				}
			} else {
				if count == 3 {
					newMap[i][j] = '#'
				} else {
					newMap[i][j] = '.'
				}
			}
		}
	}
	return newMap
}

func main() {
	day := 18

	path, part := setup.Parameters(day)
	setup.Banner(day, part)

	m, err := load.Map(path)
	if err != nil {
		log.Fatal(err)
	}

	if part == 2 {
		m[0][0] = '#'
		m[0][len(m[0])-1] = '#'
		m[len(m)-1][0] = '#'
		m[len(m)-1][len(m[0])-1] = '#'
	}

	for range 100 {
		m = step(m)
		if part == 2 {
			m[0][0] = '#'
			m[0][len(m[0])-1] = '#'
			m[len(m)-1][0] = '#'
			m[len(m)-1][len(m[0])-1] = '#'
		}
	}

	count := 0
	for i := range m {
		for j := range m[i] {
			if m[i][j] == '#' {
				count++
			}
		}
	}

	fmt.Printf("There are %d lights on.\n", count)
}
