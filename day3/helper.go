package main

import (
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToi(s string) int {
	s2n := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
	if s2n[s] != 0 || s == "0" {
		return s2n[s]
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return num
}

func readFile(fname string) string {
	data, err := os.ReadFile(fname)
	check(err)
	return string(data)
}

func reMatch(pattern string, s string) ([]string, [][]int) {
	r, _ := regexp.Compile(pattern)
	return r.FindAllString(s, -1), r.FindAllStringIndex(s, -1)
}
