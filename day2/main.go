package main

import (
	"bufio"
	"fmt"
	"os"
)

var cubeLimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	INPUT := "input"
	file, err := os.Open(INPUT)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum1 := 0
	sum2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		// get the game index
		_, _, gSubmatches := reMatch("Game (\\d+)", line)
		gIndex := convertToInt(gSubmatches[0][1])

		// For each color, check if the cube count was valid
		// Also retrieve its highest value
		validRedGame, redMax := processColorCubeCount("red", line)
		validGreenGame, greenMax := processColorCubeCount("green", line)
		validBlueGame, blueMax := processColorCubeCount("blue", line)

		// if all color counts are valid the game is valid
		validGame := (validRedGame && validBlueGame && validGreenGame)
		// calculate power of this game
		power := (redMax * greenMax * blueMax)

		if validGame {
			sum1 += gIndex
		}

		sum2 += power
	}

	fmt.Println("Part 1, Sum:", sum1)
	fmt.Println("Part 2, Sum:", sum2)
}

func processColorCubeCount(color, line string) (bool, int) {
	validColorGame := true
	colorMax := -1
	_, _, colorSubmatches := reMatch("(\\d+) "+color, line)
	for _, submatch := range colorSubmatches {
		colorCount := convertToInt(submatch[1])

		// for PART 1
		// its not a validGame if colorCount is more then colorLimit
		if validColorGame && (colorCount > cubeLimit[color]) {
			validColorGame = false
		}

		// for PART 2
		// get the highest color count
		if colorCount > colorMax {
			colorMax = colorCount
		}
	}
	return validColorGame, colorMax
}
