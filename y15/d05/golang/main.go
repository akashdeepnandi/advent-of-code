package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	file, _ := os.Open(*f)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	solve(scanner)
}

func solve(scanner *bufio.Scanner) {
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sCount := 0
	for _, line := range lines {
		vowels, _, _ := reMatch("(a|e|i|o|u)", line)

		if len(vowels) < 3 { // check for vowels
			continue
		}
		doubles := false
		runes := []rune(line)
		check := runes[0]
		for i := 1; i < len(runes); i++ {
			if check == runes[i] {
				doubles = true
				break
			} else {
				check = runes[i]
			}
		}

		if !doubles { // check for doubles
			continue
		}

		bad, _, _ := reMatch("(ab|cd|pq|xy)", line)

		if len(bad) > 0 { // check for bad patterns
			continue
		}

		sCount++

	}

	fmt.Println(sCount)
}

func part2(lines []string) {
	sCount := 0
	for _, line := range lines {
		runes := []rune(line)

		pairs := false
		for i := 0; i < len(runes)-3; i++ {
			for j := i + 2; j < len(runes)-1; j++ {
				if runes[i] == runes[j] && runes[i+1] == runes[j+1] {
					pairs = true
					break

				}
			}
			if pairs {
				break
			}
		}

		if !pairs { // check for doubles
			continue
		}

		oneBetween := false
		for i := 0; i < len(runes)-2; i++ {
			if runes[i] == runes[i+2] {
				oneBetween = true
				break
			}
		}

		if !oneBetween { // check for bad patterns
			continue
		}

		sCount++

	}

	fmt.Println(sCount)
}
