package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	lines := readLines(*f)
	solve(lines)
}

// Why bytes? String is immutable, tried and failed.
// slices were easier to handle
func lookAndSay(sequence []byte) []byte {
	result := make([]byte, 0)
	currentDigit := sequence[0]
	count := 0
	for _, d := range sequence {
		if currentDigit == d {
			count++
		} else {
			result = append(result, strconv.Itoa(count)...)
			result = append(result, currentDigit)

			currentDigit = d
			count = 1
		}
	}
	result = append(result, strconv.Itoa(count)...)
	result = append(result, currentDigit)

	return result
}

func evolveSequence(sequence []byte, generation int) []byte {
	for range generation {
		sequence = lookAndSay(sequence)
	}
	return sequence
}

func solve(lines []string) {
	sequence := []byte(lines[0])

	part1 := evolveSequence(sequence, 40)
	fmt.Println("Part 1:", len(part1))

	part2 := evolveSequence(sequence, 50)
	fmt.Println("Part 2:", len(part2))
}
