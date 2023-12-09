package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	INPUT := "input"
	file, err := os.Open(INPUT)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	solution(scanner)
}

func solution(scanner *bufio.Scanner) {
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		m := strings.Split(line, " ")
		nums := stringToIntSlice(m)
		lastVals := [][]int{}
		fl := []int{nums[0], nums[len(nums)-1]}
		// Storing the first and last values for each generations
		lastVals = append(lastVals, fl)
		for {
			newNums := []int{}
			allZero := true
			for i := 0; i < len(nums)-1; i++ {
				a := nums[i]
				b := nums[i+1]
				diff := b - a
				if diff != 0 {
					allZero = false
				}
				newNums = append(newNums, diff)
			}
			nums = newNums
			fl := []int{nums[0], nums[len(nums)-1]}
			lastVals = append(lastVals, fl)
			if allZero {
				break
			}

		}

		firstVal := 0
		lastVal := 0
		for i := len(lastVals) - 1; i >= 0; i-- {
			previousVals := lastVals[i]
			// FOR PART 1
			previousFirst := previousVals[0]
			firstVal = previousFirst - firstVal
			// FOR PART 2
			previousLast := previousVals[1]
			lastVal = lastVal + previousLast
		}

		// FOR PART 1
		sum1 += lastVal

		// FOR PART 2
		sum2 += firstVal
	}

	fmt.Println("Part 1", sum1)
	fmt.Println("Part 1", sum2)
}
