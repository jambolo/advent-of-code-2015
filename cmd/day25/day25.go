package main

import (
	"fmt"

	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

var row = 3010
var column = 3019

func main() {
	day := 25

	_, part := setup.Parameters(day)
	setup.Banner(day, part)

	// The index of the code is T_n + c, where T_n is the nth triangular number and n is (r + c - 2)
	n := row + column - 2
	t := n * (n + 1) / 2
	index := t + column

	x := 20151125
	a := 252533
	m := 33554393
	for i := 1; i < index; i++ {
		x = (x * a) % m
	}
	fmt.Printf("Code: %d\n", x)
}
