package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func main() {
	INPUT := "sinput"
	file, err := os.Open(INPUT)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	solution(scanner)
}

type Hand struct {
	cards    string
	strength int
}

type HandBid struct {
	hand Hand
	bid  int
}

func createHandBid(h, bs string, wildcard bool) HandBid {
	hand := Hand{h, 0}
	if wildcard {
		hand.addStrengthWildcard()
	} else {
		hand.addStrength()
	}
	return HandBid{hand, convertToInt(bs)}
}

func getTotalWinnins(HB []HandBid, part string) {
	// for reverse order, return false on higher value
	sort.Slice(HB, func(i, j int) bool {
		hand1 := HB[i].hand
		hand2 := HB[j].hand
		if hand1.strength == hand2.strength {
			return hand1.cards < hand2.cards
		}
		return hand1.strength < hand2.strength
	})
	sum1 := 0
	for i, hb := range HB {
		bid := hb.bid
		sum1 += (i + 1) * bid
	}

	println(part, sum1)

}

func solution(scanner *bufio.Scanner) {
	HB1 := []HandBid{}
	HB2 := []HandBid{}
	for scanner.Scan() {
		line := scanner.Text()
		h, bs := stringTwoSplit(line, " ")
		hb1 := createHandBid(h, bs, false)
		hb2 := createHandBid(h, bs, true)
		HB1 = append(HB1, hb1)
		HB2 = append(HB2, hb2)
	}
	getTotalWinnins(HB1, "Part 1")
	getTotalWinnins(HB2, "Part 2")
}

var cardMap1 = map[string]string{
	"A": "E",
	"K": "D",
	"Q": "C",
	"J": "B",
	"T": "A",
}

// FOR PART 1
func (h *Hand) addStrength() {
	cardsA := strings.Split(h.cards, "")
	cardCount := make(map[string]int)
	for i, c := range cardsA {
		l := cardMap1[c]
		k := c
		if l != "" {
			cardsA[i] = l
			k = l
		}
		cardCount[k]++
	}
	counts := getMapValues(cardCount)
	strength := getStrength(counts)
	h.cards = strings.Join(cardsA, "")
	h.strength = strength
}

var cardMap2 = map[string]string{
	"A": "E",
	"K": "D",
	"Q": "C",
	"J": "1", // assigned J lower value for PART 2
	"T": "A",
}

// FOR PART 2
func (h *Hand) addStrengthWildcard() {
	cardsA := strings.Split(h.cards, "")
	cardCount := make(map[string]int)
	cardArr := []string{}
	for i, c := range cardsA {
		l := cardMap2[c]
		k := c
		if l != "" {
			cardsA[i] = l
			k = l
		}
		cardCount[k]++
		cardArr = append(cardArr, k)
	}

	h.cards = strings.Join(cardsA, "")
	if len(cardArr) > 1 {
		highCard := "a" // assign a value which will never be in cardCount
		for _, card := range cardArr {
			if card != "1" && cardCount[card] > cardCount[highCard] {
				highCard = card
			}
		}

		if highCard != "a" && cardCount["1"] > 0 {
			cardCount[highCard] += cardCount["1"]
			delete(cardCount, "1")
		}

	}

	counts := getMapValues(cardCount)
	strength := getStrength(counts)
	h.strength = strength
}

func getStrength(c []int) int {
	if Contains(c, 5) { // five of kind
		return 6
	} else if Contains(c, 4) { // four of kind
		return 5
	} else if Contains(c, 3) { // full house or three kind
		if Contains(c, 2) {
			return 4
		} else {
			return 3
		}
	} else if Contains(c, 2) {
		if len(c) == 3 {
			return 2
		} else {
			return 1
		}
	}

	return 0
}
