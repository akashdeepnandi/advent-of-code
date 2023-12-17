package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	file, err := os.Open(*f)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	solution(scanner)
}

var cache = make(map[string]int)

func solution(scanner *bufio.Scanner) {
	sum1, sum2 := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		springs, g := stringTwoSplit(line, " ")
		groups := stringToIntSlice(strings.Split(g, ","))
		sum1 += getArrangements(springs, groups)

		// FOR PART 2 --
		springs = strings.Join(multiply(strings.Split(springs, ""), 5, "?"), "")
		g = strings.Join(multiply(strings.Split(g, ""), 5, ","), "")
		groups = stringToIntSlice(strings.Split(g, ","))
		sum2 += getArrangements(springs, groups)
	}

	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}

func multiply(s []string, c int, sep string) []string {
	l := len(s)
	nl := c * l
	if sep != "" {
		nl += c - 1
	}
	ns := make([]string, 0, nl)
	for i := 0; i < c; i++ {
		ns = append(ns, s...)
		if sep != "" && i != c-1 {
			ns = append(ns, sep)
		}
	}

	return ns
}

func getArrangements(springs string, groups []int) int {
	gk := strings.Join(convertSlice[int, string](groups, strconv.Itoa), ",")
	k := springs + gk
	if _, ok := cache[k]; ok {
		return cache[k]
	}

	if springs == "" {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(groups) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		} else {
			return 1
		}
	}

	count := 0
	group := groups[0]
	c := springs[0:1]
	springsL := len(springs)

	if c == "?" || c == "." {
		count += getArrangements(springs[1:], CopySlice(groups))
	}
	if c == "?" || c == "#" {
		start := min(group, springsL-1)
		if group <= springsL &&
			strings.Contains(springs[:min(group, springsL)], ".") == false &&
			(group == springsL || springs[start:start+1] != "#") {
			count += getArrangements(springs[min(group+1, springsL):], CopySlice(groups[1:]))

		}
	}

	cache[k] = count

	return count
}
