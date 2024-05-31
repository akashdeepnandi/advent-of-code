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
	// file, err := os.Open(*f)
	// check(err)
	// defer file.Close()
	//
	// scanner := bufio.NewScanner(file)
	//
	// solve(scanner)
	solve(lines)
}

func solve(lines []string) {
	p1 := 0
	p2 := 0
	for _, line := range lines {
		curLen := len(line)
		x, _ := strconv.Unquote(line)
		p1 += curLen - len(x)

		y := strconv.Quote(line)
		p2 += len(y) - curLen
	}

	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
}

// func solve(scanner *bufio.Scanner) {
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fmt.Println(line)
// 	}
// \"c\x04 - 3 | 7 - 1 - 3  = 4
// \x67lwkks  - 6 | 9 - 3 = 6
// }
