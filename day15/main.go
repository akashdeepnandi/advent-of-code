package main

import (
	"flag"
	"fmt"
	"strings"
)

type LabelLens struct {
	label       string
	focalLength int
}

func (ll *LabelLens) hasLabel(label string) bool {
	return ll.label == label
}

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	lines := readLines(*f)
	solve(lines[0])
}

func solve(input string) {
	steps := strings.Split(input, ",")
	sum1 := 0
	boxes := make([][]LabelLens, 256)
	for _, step := range steps {
		hash := getHash(step)
		sum1 += hash

		// for part 2
		if strings.Contains(step, "-") {
			label, _ := stringTwoSplit(step, "-")
			boxNum := getHash(label)
			labelIndex := findLabel(label, boxes[boxNum])
			if labelIndex != -1 {
				boxes[boxNum] = append(boxes[boxNum][0:labelIndex], boxes[boxNum][labelIndex+1:]...)
			}
		} else if strings.Contains(step, "=") {
			label, lens := stringTwoSplit(step, "=")
			boxNum := getHash(label)
			labelIndex := findLabel(label, boxes[boxNum])
			ll := LabelLens{label, convertToInt(lens)}
			if labelIndex != -1 {
				boxes[boxNum][labelIndex] = ll
			} else {
				boxes[boxNum] = append(boxes[boxNum], ll)
			}
		}
	}

	sum2 := 0

	for i, box := range boxes {
		focusPower := calFocusPower(i, box)
		sum2 += focusPower
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func calFocusPower(boxNum int, box []LabelLens) int {
	if len(box) == 0 {
		return 0
	}
	power := 0
	for i, ll := range box {
		power += (boxNum + 1) * ((i + 1) * ll.focalLength)
	}

	return power
}

func findLabel(label string, box []LabelLens) int {
	for i, labelLens := range box {
		if labelLens.hasLabel(label) {
			return i
		}
	}
	return -1
}

func getHash(s string) int {
	num := 0
	for _, c := range s {
		num += int(c)
		num *= 17
		num %= 256
	}

	return num
}
