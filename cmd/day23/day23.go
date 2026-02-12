package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/jambolo/advent-of-code-2015/internal/load"
	"github.com/jambolo/advent-of-code-2015/internal/setup"
)

type instruction struct {
	opcode string
	reg    int
	offset int
}

type cpu struct {
	registers [2]int
	pc        int
}

func (c *cpu) hlf(r *int) {
	*r /= 2
	c.pc++
}

func (c *cpu) tpl(r *int) {
	*r *= 3
	c.pc++
}

func (c *cpu) inc(r *int) {
	*r++
	c.pc++
}

func (c *cpu) jmp(offset int) {
	c.pc += offset
}

func (c *cpu) jie(r *int, offset int) {
	if *r%2 == 0 {
		c.pc += offset
	} else {
		c.pc++
	}
}

func (c *cpu) jio(r *int, offset int) {
	if *r == 1 {
		c.pc += offset
	} else {
		c.pc++
	}
}

func (c *cpu) execute(inst []instruction) {
	for c.pc >= 0 && c.pc < len(inst) {
		i := inst[c.pc]
		switch i.opcode {
		case "hlf":
			c.hlf(&c.registers[i.reg])
		case "tpl":
			c.tpl(&c.registers[i.reg])
		case "inc":
			c.inc(&c.registers[i.reg])
		case "jmp":
			c.jmp(i.offset)
		case "jie":
			c.jie(&c.registers[i.reg], i.offset)
		case "jio":
			c.jio(&c.registers[i.reg], i.offset)
		}
	}
}

func main() {
	day := 23

	file, part := setup.Parameters(day)
	setup.Banner(day, part)

	lines, err := load.Lines(file)
	if err != nil {
		log.Fatal(err)
	}

	program := parseInstructions(lines)

	c := &cpu{}
	if part == 2 {
		c.registers[0] = 1
	}
	c.execute(program)
	println("Result:", c.registers[1])
}

func parseInstructions(lines []string) []instruction {
	var instructions []instruction

	for _, line := range lines {
		instructions = append(instructions, parseInstruction(line))
	}

	return instructions
}

func parseInstruction(line string) instruction {

	// hlf r sets register r to half its current value, then continues with the next instruction.
	// tpl r sets register r to triple its current value, then continues with the next instruction.
	// inc r increments register r, adding 1 to it, then continues with the next instruction.
	// jmp offset is a jump; it continues with the instruction offset away relative to itself.
	// jie r, offset is like jmp, but only jumps if register r is even ("jump if even").
	// jio r, offset is like jmp, but only jumps if register r is 1 ("jump if one", not odd).

	tokens := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ',' })

	var i instruction
	i.opcode = tokens[0]
	switch i.opcode {
	case "hlf", "tpl", "inc":
		// Extract the register (a or b)
		if tokens[1] == "a" {
			i.reg = 0
		} else {
			i.reg = 1
		}
	case "jmp":
		// Extract the offset
		i.offset, _ = strconv.Atoi(tokens[1])
	case "jie", "jio":
		// Extract the register (a or b)
		if tokens[1] == "a" {
			i.reg = 0
		} else {
			i.reg = 1
		}
		// Extract the offset
		i.offset, _ = strconv.Atoi(tokens[2])
	}

	return i
}
