package main

import (
	"bufio"
	"fmt"
	"os"
)

func main1() {
	INPUT := "sinput"
	file, err := os.Open(INPUT)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	solution(scanner)
}

func solution1(scanner *bufio.Scanner) {
	seeds := []int{}
	maps := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		re := "\\d+"
		matches, _, _ := reMatch(re, line)
		if len(seeds) == 0 {
			seeds = stringToIntSlice(matches)
			scanner.Scan()
			continue
		} else if len(matches) > 0 {
			maps = append(maps, stringToIntSlice(matches))
		}
	}

	fmt.Println(seeds)
	fmt.Println(maps)
}
