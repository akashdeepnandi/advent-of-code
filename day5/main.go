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
	solution(scanner)
}

func solution(scanner *bufio.Scanner) {
	mapKeys := []string{"sS", "sF", "fW", "wL", "lT", "tH", "hL"}

	mapKeyIndex := 0
	seeds := []int{}
	inputMap := make(map[string][][]int)
	for scanner.Scan() {
		mk := mapKeys[mapKeyIndex]
		v := inputMap[mk]

		line := scanner.Text()
		re := "\\d+"
		matches, _, _ := reMatch(re, line)
		if len(seeds) == 0 {
			seeds = stringToIntSlice(matches)
			scanner.Scan()
			continue
		}
		if len(matches) > 0 {
			v = append(v, stringToIntSlice(matches))
			inputMap[mk] = v
		}
		if line == "" {
			mapKeyIndex++
			continue
		}
	}

	minLoc1 := -1
	for _, seed := range seeds {
		loc := getSeedMinLoc(seed, inputMap)
		if minLoc1 == -1 || loc < minLoc1 {
			minLoc1 = loc
		}
	}

	minLoc2 := -1
	fmt.Println("Lowest location is", minLoc1)
	fmt.Println("Lowest location is", minLoc2)
}

func addRange(ranges [][]int, nr []int) [][]int {
	for _, r := range ranges {
		if r[0] == nr[0] && r[1] == nr[1] {
			return ranges
		}
	}

	return append(ranges, nr)
}

func getSeedMinLoc(seed int, inputMap map[string][][]int) int {
	soil := getValueFromRange(seed, inputMap["sS"])
	fertilizer := getValueFromRange(soil, inputMap["sF"])
	water := getValueFromRange(fertilizer, inputMap["fW"])
	light := getValueFromRange(water, inputMap["wL"])
	temp := getValueFromRange(light, inputMap["lT"])
	humidity := getValueFromRange(temp, inputMap["tH"])
	loc := getValueFromRange(humidity, inputMap["hL"])
	return loc
}

func getValueFromRange(lookUp int, lookUpRange [][]int) int {
	for _, input := range lookUpRange {
		dStart, sStart, mrange := getInputVals(input)
		if lookUp >= sStart && lookUp < sStart+mrange {
			return dStart + lookUp - sStart
		}
	}
	return lookUp
}

// returns destStart, sourceStart, mrange
func getInputVals(r []int) (int, int, int) {
	return r[0], r[1], r[2]
}
