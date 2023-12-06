package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	INPUT_FILE := "./input1"
	// INPUT_FILE := "./sample_input"
	file, err := os.Open(INPUT_FILE)
	check(err)
	scanner := bufio.NewScanner(file)
	sum1 := 0
	sum2 := 0

	// regex to only look digits in a line as per part 1
	r1, _ := regexp.Compile("\\d")
	// regex to only look digits and digit word in a line as per part 2
	r2, _ := regexp.Compile("(\\d|one|two|three|four|five|six|seven|eight|nine)")
	for scanner.Scan() {
		line := scanner.Text()
		// get the rowDigit as per the regex for each
		rd1 := extractRowDigit(line, r1)
		rd2 := extractRowDigit(line, r2)
		// add to sum
		sum1 += rd1
		sum2 += rd2

	}
	fmt.Println("Part 1: Sum", sum1)
	fmt.Println("Part 2: Sum", sum2)
}

func extractRowDigit(s string, r *regexp.Regexp) int {
	var matches []string
	for i := 0; i < len(s)-1; i++ {
		m := (r.FindAllString(s[i:], -1))
		matches = append(matches, m...)
	}
	mlen := len(matches)
	fnum := 0
	lnum := 0
	if mlen > 0 {
		fm := matches[0]
		fnum = convertToi(fm)
		lm := matches[mlen-1]
		lnum = convertToi(lm)
		return fnum*10 + lnum
	}
	return 0
}
