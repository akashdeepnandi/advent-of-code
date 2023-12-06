package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	INPUT := "input"
	lines := readLines(INPUT)
	solveQuadratic(lines)
}

func solveQuadratic(lines []string) {
	matches, _, _ := reMatch("(\\d+)", lines[0])
	times := stringToIntSlice(matches)
	matches, _, _ = reMatch("(\\d+)", lines[1])
	distances := stringToIntSlice(matches)

	part1MarginError := 1
	for i, time := range times {
		distance := distances[i]
		wins := getWinsUsingRoots(float64(time), float64(distance))
		part1MarginError *= wins
	}

	time := getInputNum(lines[0])
	distance := getInputNum(lines[1])
	part2Win := getWinsUsingRoots(float64(time), float64(distance))

	fmt.Println("Part 1: ", part1MarginError)
	fmt.Println("Part 2: ", part2Win)

}

func solveBruteForce(lines []string) {
	matches, _, _ := reMatch("(\\d+)", lines[0])
	times := stringToIntSlice(matches)
	matches, _, _ = reMatch("(\\d+)", lines[1])
	distances := stringToIntSlice(matches)

	waysToWin := []int{}
	for i, time := range times {
		distance := distances[i]
		wins := getWaysToWin(time, distance)
		waysToWin = append(waysToWin, wins)
	}

	part1MarginError := 1
	for _, ways := range waysToWin {
		part1MarginError *= ways
	}

	time := getInputNum(lines[0])
	distance := getInputNum(lines[1])
	part2Win := getWaysToWin(time, distance)

	fmt.Println("Part 1: ", part1MarginError)
	fmt.Println("Part 2: ", part2Win)

}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func getWinsUsingRoots(time, distance float64) int {
	// time * x - x ** 2 > distance
	// x ^ 2 + (- time) * x + distance > 0 - this is quadratic inequality
	a := 1.0
	b := -time
	c := distance
	sqrt := math.Sqrt(b*b - 4*a*c)
	// find the roots
	r1 := ((-b - sqrt) / 2 * a) // smaller root
	r2 := ((-b + sqrt) / 2 * a) // bigger root
	// check if the roots are integers
	if isInt(r1) {
		r1++
	}
	if isInt(r2) {
		r2--
	}
	// get the upper and lower limit of exclusive values
	r1 = math.Ceil(r1)    // lower limit
	r2 = math.Floor(r2)   // upper limit
	return int(r2-r1) + 1 // get the range length - upper - lower + 1
}

func isInt(n float64) bool {
	return n == float64(int(n))
}

func getInputNum(line string) int {
	m, _, _ := reMatch("\\d+.*", line)
	m = strings.Split(m[0], " ")
	num := ""
	for _, s := range m {
		t := strings.TrimSpace(s)
		if t != "" {
			num += t
		}
	}

	return convertToInt(num)
}

func getWaysToWin(time, distance int) int {
	winCount := 0
	for speed := 1; speed < time; speed++ {
		distanceTravelled := speed * (time - speed)
		if distanceTravelled > distance {
			winCount++
		}
	}
	if winCount > 0 {
		return winCount
	}
	return 1
}
