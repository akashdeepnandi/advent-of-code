package main

import (
	"fmt"
	"strings"
)

func main() {
	// INPUT := "sinput"
	INPUT := "sinput"
	data := readFile(INPUT)
	lines := strings.Split(data, "\n")
	solution(lines[:len(lines)-1])

}

func solution(lines []string) {
	llen := len(lines)
	sum := 0
	sumGearRatio := 0
	gearMap := make(map[string][]int)

	// loop over all lines
	for i := 0; i < llen; i++ {
		line := lines[i]
		// get the numbers out of the line and their matched indices
		nums, indices := reMatch("\\d+", line)
		chars := strings.Split(line, "")

		for j, num := range nums {
			isPart := false
			gearIndex := ""
			idx := indices[j]
			start := idx[0] // get match start
			end := idx[1]   // get match end

			// check left -> start - 1
			if start-1 >= 0 {
				c := chars[start-1]

				m, _ := reMatch("[^\\d\\.]", c)
				if len(m) > 0 {
					isPart = true
					if m[0] == "*" {
						gearIndex = fmt.Sprintf("x:%d,y:%d", start-1, i)
					}
				}
			}

			// check right -> end
			if end != len(chars) {
				c := chars[end]

				m, _ := reMatch("[^\\d\\.]", c)
				if len(m) > 0 {
					isPart = true
					if m[0] == "*" {
						gearIndex = fmt.Sprintf("x:%d,y:%d", end, i)
					}
				}
			}

			aStart := max(0, start-1)
			aEnd := min(len(chars)-1, end+1)

			// check top -> i - 1, start - 1 to end
			if i != 0 {
				tChars := strings.Split(lines[i-1], "")

				for x := aStart; x < aEnd; x++ {
					c := tChars[x]

					m, _ := reMatch("[^\\d\\.]", c)
					if len(m) > 0 {
						isPart = true
						if m[0] == "*" {
							gearIndex = fmt.Sprintf("x:%d,y:%d", x, i-1)
						}
						break
					}

				}
			}

			// check bottom -> i - 1, start - 1 to end
			if i < len(lines)-1 {
				bChars := strings.Split(lines[i+1], "")

				for x := aStart; x < aEnd; x++ {
					c := bChars[x]

					m, _ := reMatch("[^\\d\\.]", c)
					if len(m) > 0 {
						isPart = true
						if m[0] == "*" {
							gearIndex = fmt.Sprintf("x:%d,y:%d", x, i+1)
						}
						break
					}

				}

			}

			if isPart {
				n := convertToi(num)
				sum += n
				// fmt.Println(gearIndex)
				if gearIndex != "" {
					gearMap[gearIndex] = append(gearMap[gearIndex], n)
					if len(gearMap[gearIndex]) == 2 {
						p0 := gearMap[gearIndex][0]
						p1 := gearMap[gearIndex][1]
						gearRatio := p0 * p1

						sumGearRatio += gearRatio
					}
				}
			}
		}
	}

	fmt.Println("Total sum:", sum)
	fmt.Println("Total sum of gearRatio:", sumGearRatio)
}
