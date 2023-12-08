package main

import (
	"fmt"
	"strings"
)

func main() {
	INPUT := "input"
	lines := readLines(INPUT)
	solution(lines)
}

type Node struct {
	left  string
	right string
}

func solution(lines []string) {
	var instructions []string
	network := make(map[string]Node)
	startNodes := []string{}
	for i, line := range lines {
		if i == 0 {
			instructions = strings.Split(line, "")
			continue
		} else if line != "" {
			m, _, _ := reMatch("(\\w+)", line)
			n, l, r := getNodeVal(m)
			network[n] = Node{l, r}
			if strings.HasSuffix(n, "A") {
				startNodes = append(startNodes, n)
			}
		}
	}

	// FOR PART 1
	steps := 0
	startNode := "AAA"
	steps = findSteps(network, startNode, instructions, false)
	fmt.Println("Part 1", steps)
	// FOR PART 1

	// FOR PART 2
	stepsTaken := []int{}
	for _, currentNode := range startNodes {
		steps = findSteps(network, currentNode, instructions, true)
		stepsTaken = append(stepsTaken, steps)
	}
	fmt.Println(lcmN(stepsTaken))
	fmt.Println("Part 2", steps)
	// FOR PART 2
}

func findSteps(network map[string]Node, startNode string, instructions []string, isPart2 bool) int {
	steps := 0
	for {
		currentNode := startNode
		for _, nav := range instructions { // nav = "L" | "R"
			steps++
			node := network[currentNode]
			if nav == "L" {
				currentNode = node.left
			} else {
				currentNode = node.right
			}
			if currentNode == "ZZZ" && isPart2 == false {
				return steps
			}
			if strings.HasSuffix(currentNode, "Z") && isPart2 {
				return steps
			}
		}
		startNode = currentNode
	}
}

func getNodeVal(m []string) (string, string, string) {
	return m[0], m[1], m[2]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	// a * b = lcm(a, b) * gcd(a, b)
	return (a * b) / gcd(a, b)
}

func lcmN(n []int) int {
	if len(n) == 2 {
		return lcm(n[0], n[1])
	}
	return lcm(n[0], lcmN(n[1:]))
}
