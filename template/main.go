package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	INPUT := "sinput"
	file, err := os.Open(INPUT)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1(scanner)
}

func part1(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func part2(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
