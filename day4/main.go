package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	INPUT := "input"
	cards := readLines(INPUT)
	solution(cards)
}

func solution(cards []string) {
	// PART 1 -- START
	sum1 := 0
	// FOR PART 2, initializing an card copy tracker with start count 1
	cardsLen := len(cards)
	cardCopyTracker := make([]int, cardsLen)
	for i := range cardCopyTracker {
		cardCopyTracker[i] = 1
	}

	for cardIdx, card := range cards {
		_, allNumsStr := stringTwoSplit(card, ": ")
		winningNumsStr, myNumsStr := stringTwoSplit(allNumsStr, " | ")
		winningNums := stringToIntSlice(strings.Split(winningNumsStr, " "))
		myNums := stringToIntSlice(strings.Split(myNumsStr, " "))
		matchCount := 0
		cardPoints := 0
		for _, num := range myNums {
			isWinNum := Contains[int](winningNums, num)
			if isWinNum {
				matchCount++
			}
		}

		if matchCount != 0 {
			cardPoints = int(math.Pow(2, float64(matchCount)-1))
			// FOR PART 2
			// find the range of card Indexes to increase copy count
			startIdx := cardIdx + 1
			endIdx := min(cardsLen, cardIdx+matchCount)
			currentCardCopies := cardCopyTracker[cardIdx]
			for i := startIdx; i <= endIdx; i++ {
				// increase count by adding copy of current card
				cardCopyTracker[i] += currentCardCopies
			}
		}
		sum1 += cardPoints
	}
	fmt.Println("Part 1 Sum", sum1)
	sum2 := 0
	for _, v := range cardCopyTracker {
		sum2 += v
	}
	fmt.Println("Part 2 Sum", sum2)
	// PART 1 -- END
}
