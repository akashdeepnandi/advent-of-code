package main

import (
	"flag"
	"fmt"
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
	for _, line := range lines {
		fmt.Println(line)
	}
}

// func solve(scanner *bufio.Scanner) {
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fmt.Println(line)
// 	}
// }
