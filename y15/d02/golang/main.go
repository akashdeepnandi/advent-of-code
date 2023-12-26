package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	file, err := os.Open(*f)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	solve(scanner)
}

func solve(scanner *bufio.Scanner) {
	totalPaper := 0
	totalRibbon := 0
	for scanner.Scan() {
		// for part 1
		l, w, h := getDims(scanner.Text())
		ma := min(l*w, w*h, h*l)
		sa := 2*(l*w+w*h+h*l) + ma
		totalPaper += sa
		// for part 2: vol + shortest dist between sides
		totalRibbon += l*w*h + 2*min(l+w, w+h, h+l)
	}
	fmt.Println("Total Wrapping Paper", totalPaper)
	fmt.Println("Total Ribbon", totalRibbon)
}

func getDims(s string) (int, int, int) {
	dim := stringToIntSlice(strings.Split(s, "x"))
	return dim[0], dim[1], dim[2]
}
