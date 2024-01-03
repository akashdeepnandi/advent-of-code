package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	lines := readLines(*f)
	solve(lines)
}

type Instruction struct {
	gate  string   // NOT | SIGNAL | AND | OR | LSHIFT | RSHIFT | WIRE
	wires []string // Source wires | value
	tid   string   // Target wire identifier
}

func solve(lines []string) {
	instructions := []Instruction{}
	for _, line := range lines {
		opr, id := stringTwoSplit(line, " -> ")
		wires := strings.Split(opr, " ")
		instr := Instruction{tid: id}
		if len(wires) == 2 { // NOT
			instr.gate = "NOT"
			instr.wires = wires[1:]
		} else if len(wires) == 1 { // WIRE | SIGNAL
			instr.gate = "WIRE"
			m, _, _ := reMatch("\\d+", wires[0])
			if len(m) > 0 {
				instr.gate = "SIGNAL"
			}
			instr.wires = wires
		} else { // AND | OR | LSHIFT | RSHIFT
			instr.gate = wires[1]
			instr.wires = append(wires[0:1], wires[2])
		}

		instructions = append(instructions, instr)
	}

	// Part 1
	val := connectWires(instructions, "a")
	fmt.Println("Part 1", val)

	// For Part 2, update signal of b by value of a
	bIndex := 0
	for i := range instructions {
		if instructions[i].tid == "b" {
			bIndex = i
		}
	}

	instructions[bIndex].wires = []string{strconv.Itoa(int(val))}
	val = connectWires(instructions, "a")
	fmt.Println("Part 2", val)
}

func connectWires(instructions []Instruction, wire string) uint16 {
	values := make(map[string]uint16)
	visited := make(map[string]bool)
	getWireVal := func(wire string) (bool, uint16) {
		val := convertToInt(wire)
		if val != -1 {
			return true, uint16(val)
		}
		return visited[wire], values[wire]
	}

	for {
		resolved := true
		for _, instr := range instructions {
			tid, gate, wires := instr.tid, instr.gate, instr.wires
			if len(wires) == 1 { // single operand - SIGNAL | WIRE | NOT
				vis, val := getWireVal(wires[0]) // check visit wire 1
				if !vis {
					resolved = false
					continue
				}
				if gate == "NOT" { // check if NOT
					val = ^val
				}
				values[tid] = val
				visited[tid] = true
			} else { // two operands
				vis1, val1 := getWireVal(wires[0]) // check visit wire 1
				vis2, val2 := getWireVal(wires[1]) // check visit wire 2
				if !vis1 || !vis2 {                // did not visit one of them
					resolved = false
					continue
				}
				var val uint16

				switch gate {
				case "AND":
					val = val1 & val2
				case "OR":
					val = val1 | val2
				case "LSHIFT":
					val = val1 << val2
				case "RSHIFT":
					val = val1 >> val2
				}

				values[tid] = val
				visited[tid] = true
			}
		}

		if resolved {
			break
		}
	}
	return values[wire]
}
