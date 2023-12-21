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

type Rule struct {
	part     string
	operator string
	value    int
	dest     string // can be A/R/Rule.name
}

type Workflow struct {
	rules    []Rule
	lastDest string // can be A/R/Rule.name
}

func makeWorkflow(s string) (string, Workflow) {
	name, rest := stringTwoSplit(s, "{")
	w := Workflow{}
	rules := []Rule{}
	parts := strings.Split(rest, ",")
	for i, r := range parts {
		if i == len(parts)-1 {
			r = r[:len(r)-1]
			w.lastDest = r
		} else {
			rest, dest := stringTwoSplit(r, ":")
			operator := findOperator(rest)
			part, value := stringTwoSplit(rest, operator)
			rules = append(rules, Rule{part, operator, convertToInt(value), dest})
		}
	}

	w.rules = rules
	return name, w
}

func findOperator(s string) string {
	if strings.Contains(s, "=") {
		return "="
	} else if strings.Contains(s, "<") {
		return "<"
	} else {
		return ">"
	}
}

type Rating struct {
	part  string
	value int
}

func (ru *Rule) operate(value int) bool {
	if ru.operator == "=" && value == ru.value {
		return true
	} else if ru.operator == ">" && value > ru.value {
		return true
	} else if ru.operator == "<" && value < ru.value {
		return true
	}
	return false
}

func makeRatings(s string) []Rating {
	s = s[1 : len(s)-1] // remove { }
	temp := strings.Split(s, ",")
	ratings := []Rating{}
	for _, r := range temp {
		part, value := stringTwoSplit(r, "=")
		ratings = append(ratings, Rating{part, convertToInt(value)})
	}
	return ratings
}

func makeRating(s string) map[string]int {
	s = s[1 : len(s)-1] // remove { }
	temp := strings.Split(s, ",")
	rating := make(map[string]int)
	for _, r := range temp {
		part, value := stringTwoSplit(r, "=")
		rating[part] = convertToInt(value)
	}
	return rating
}

// func (w *Workflow) matchWorkflow(ratings []Rating) (bool, string) {
// 	dest := ""
// 	match := true
//
// 	for _, r := range ratings {
// 		for i, rule := range w.rules {
// 			if rule.operate(r) { // rule does match
// 				dest = rule.dest
// 			} else {
// 				if i == len(w.rules)-1 {
// 					dest = w.lastDest
// 				} else {
// 					match = false
// 					break
// 				}
//
// 			}
// 		}
// 	}
//
// 	return match, dest
// }

func solve(scanner *bufio.Scanner) {
	// workflows := []Workflow{}
	workflows := make(map[string]Workflow)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		name, w := makeWorkflow(line)
		workflows[name] = w
		// workflows = append(workflows, makeWorkflow(line))
	}

	// machineParts := []map[string]int{}
	sum1 := 0
	for scanner.Scan() {
		line := scanner.Text()
		rating := makeRating(line)
		start := "in"
		accepted := false
		for {
			w := workflows[start]
			dest := ""
			for _, r := range w.rules {
				if rating[r.part] != 0 {
					match := r.operate(rating[r.part])
					// fmt.Println(start, w, r.part, match)
					if match {
						dest = r.dest
						break
					}
				}
			}
			if dest == "" {
				dest = w.lastDest
			}

			start = dest
			// fmt.Println(dest)
			if dest == "A" {
				accepted = true
				break
			}
			if dest == "R" {
				break
			}
		}

		if accepted {
			for _, v := range rating {
				sum1 += v
			}
		}
	}

	fmt.Println(sum1)

	// for _, ratings := range machineParts {
	// 	fmt.Println(ratings)
	// 	start := "in"
	// 	for {
	// 		w := workflows[start]
	// 		dest := ""
	// 		for _, r := range w.rules {
	// 			if ratings[r.part] != 0 {
	// 				match := r.operate(ratings[r.part])
	// 				// fmt.Println(start, w, r.part, match)
	// 				if match {
	// 					dest = r.dest
	// 					break
	// 				}
	// 			}
	// 		}
	// 		if dest == "" {
	// 			dest = w.lastDest
	// 		}
	//
	// 		start = dest
	// 		// fmt.Println(dest)
	// 		if dest == "A" || dest == "R" {
	// 			break
	// 		}
	// 	}
	// }
}
