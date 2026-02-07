package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

var (
	reAnd    = regexp.MustCompile(`^([a-z]+|\d+) AND ([a-z]+|\d+) -> ([a-z]+)$`)
	reOr     = regexp.MustCompile(`^([a-z]+|\d+) OR ([a-z]+|\d+) -> ([a-z]+)$`)
	reNot    = regexp.MustCompile(`^NOT ([a-z]+|\d+) -> ([a-z]+)$`)
	reAssign = regexp.MustCompile(`^([a-z]+|\d+) -> ([a-z]+)$`)
	reLshift = regexp.MustCompile(`^([a-z]+|\d+) LSHIFT ([a-z]+|\d+) -> ([a-z]+)$`)
	reRshift = regexp.MustCompile(`^([a-z]+|\d+) RSHIFT ([a-z]+|\d+) -> ([a-z]+)$`)
)

type op int

const (
	and op = iota
	or
	not
	assign
	lshift
	rshift
)

type gate struct {
	op     op
	lWire  string
	rWire  string
	lValue uint16
	rValue uint16
}

// parseOperand returns the wire name and literal value for an operand string.
func parseOperand(s string) (string, uint16) {
	if s == "" {
		return "", 0
	}
	if v, err := strconv.ParseUint(s, 10, 16); err == nil {
		return "", uint16(v)
	}
	return s, 0
}

func parseGate(op op, left, right string) gate {
	lWire, lValue := parseOperand(left)
	rWire, rValue := parseOperand(right)
	return gate{op: op, lWire: lWire, rWire: rWire, lValue: lValue, rValue: rValue}
}

func buildCircuit(lines []string) map[string]gate {
	gates := make(map[string]gate)
	for _, line := range lines {
		if m := reAnd.FindStringSubmatch(line); m != nil {
			gates[m[3]] = parseGate(and, m[1], m[2])
		} else if m := reOr.FindStringSubmatch(line); m != nil {
			gates[m[3]] = parseGate(or, m[1], m[2])
		} else if m := reNot.FindStringSubmatch(line); m != nil {
			gates[m[2]] = parseGate(not, m[1], "")
		} else if m := reAssign.FindStringSubmatch(line); m != nil {
			gates[m[2]] = parseGate(assign, m[1], "")
		} else if m := reLshift.FindStringSubmatch(line); m != nil {
			gates[m[3]] = parseGate(lshift, m[1], m[2])
		} else if m := reRshift.FindStringSubmatch(line); m != nil {
			gates[m[3]] = parseGate(rshift, m[1], m[2])
		} else {
			log.Fatalf("unrecognized instruction: %s", line)
		}
	}
	return gates
}

func evaluate(gates map[string]gate, cache map[string]uint16, wire string) uint16 {
	if value, ok := cache[wire]; ok {
		return value
	}
	gate, ok := gates[wire]
	if !ok {
		log.Fatalf("undefined wire: %s", wire)
	}
	resolve := func(wire string, literal uint16) uint16 {
		if wire != "" {
			return evaluate(gates, cache, wire)
		}
		return literal
	}
	lValue := resolve(gate.lWire, gate.lValue)
	rValue := resolve(gate.rWire, gate.rValue)

	var value uint16
	switch gate.op {
	case and:
		value = lValue & rValue
	case or:
		value = lValue | rValue
	case not:
		value = ^lValue
	case assign:
		value = lValue
	case lshift:
		value = lValue << rValue
	case rshift:
		value = lValue >> rValue
	default:
		log.Fatalf("invalid operation: %d", gate.op)
	}
	cache[wire] = value
	return value
}

func main() {
	day := 7

	// Grab the command line parameters (file path and part number)
	filePath, part := setup.Parameters(day)

	// Print a banner showing the current day and if it is part 1 or part 2
	setup.Banner(day, part)

	// Load the data from the specified file. Abort on error.
	lines, err := load.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// build the circuit
	circuit := buildCircuit(lines)

	// evaluate the circuit
	cache := make(map[string]uint16)
	value := evaluate(circuit, cache, "a")

	// Part 1
	if part == 1 {
		// Print the answer for part 1
		fmt.Printf("Value of wire a: %d\n", value)
		return
	}

	// Part 2
	if part == 2 {
		cache2 := make(map[string]uint16)
		cache2["b"] = value
		value2 := evaluate(circuit, cache2, "a")
		fmt.Printf("Value of wire a with b overridden: %d\n", value2)
	}
}
