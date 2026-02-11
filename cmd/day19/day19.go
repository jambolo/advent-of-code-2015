package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
	"github.com/jambolo/advent-of-code-2015/internal/utils"
)

// Entry defines the unit stored in the queue.
type Entry struct {
	value string
	g     int // Cost so far
	f     int // g + heuristic (priority)
}

// PriorityQueue implements heap.Interface, ordered by lowest count.
type PriorityQueue []Entry

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].f == pq[j].f {
		return len(pq[i].value) < len(pq[j].value)
	}
	return pq[i].f < pq[j].f
}
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Entry)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// heuristic returns the difference in string lengths.
func heuristic(node, goal string) int {
	d := len(node) - len(goal)
	if d < 0 {
		return -d
	}
	if d == 0 && node != goal {
		return math.MaxInt
	}
	return d
}

func neighborsOf(m string, replacements map[string][]string) []string {
	var neighbors []string
	for i := range m {
		prefix := m[:i]
		remainder := m[i:]
		for from, tos := range replacements {
			if strings.HasPrefix(remainder, from) {
				for _, to := range tos {
					neighbors = append(neighbors, prefix+to+remainder[len(from):])
				}
			}
		}
	}
	return neighbors
}

// aStar finds the shortest path from start to goal using A*.
func aStar(start, goal string, neighborsOf func(string, map[string][]string) []string, replacements map[string][]string) int {
	queue := &PriorityQueue{{value: start, g: 0, f: heuristic(start, goal)}}
	heap.Init(queue)
	visited := make(map[string]struct{})

	for queue.Len() > 0 {
		item := heap.Pop(queue).(Entry)
		if item.value == goal {
			return item.g
		}
		if _, ok := visited[item.value]; ok {
			continue
		}
		visited[item.value] = struct{}{}

		for _, neighbor := range neighborsOf(item.value, replacements) {
			if _, ok := visited[neighbor]; !ok {
				g := item.g + 1
				f := g + heuristic(neighbor, goal)
				heap.Push(queue, Entry{value: neighbor, g: g, f: f})
			}
		}
	}
	return -1
}

func main() {
	day := 19

	path, part := setup.Parameters(day)
	setup.Banner(day, part)

	lines, err := load.Lines(path)
	if err != nil {
		log.Fatal(err)
	}

	// Get replacements.
	replacements := make(map[string][]string)
	for i := range lines {
		if lines[i] == "" {
			lines = lines[i+1:]
			break
		}
		var from, to string
		parts := strings.Fields(lines[i])
		if len(parts) != 3 || parts[1] != "=>" {
			log.Fatalf("invalid replacement: %s", lines[i])
		}
		from, to = parts[0], parts[2]
		replacements[from] = append(replacements[from], to)
	}

	// Get the molecule.
	molecule := lines[len(lines)-1]

	if part == 1 {
		replaced := make(map[string]struct{})
		for i := range molecule {
			prefix := molecule[:i]
			remainder := molecule[i:]
			for from, to := range replacements {
				if strings.HasPrefix(remainder, from) {
					for _, t := range to {
						newMolecule := prefix + t + remainder[len(from):]
						replaced[newMolecule] = struct{}{}
					}
				}
			}
		}

		fmt.Printf("Results: %d.\n", len(replaced))
	}

	if part == 2 {
		reversed := utils.InvertMap(replacements)
		goal := "e"

		result := aStar(molecule, goal, neighborsOf, reversed)
		fmt.Printf("Results: %d.\n", result)
	}
}
