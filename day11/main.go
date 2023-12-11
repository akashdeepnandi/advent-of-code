package main

import (
	"fmt"
	"strings"
)

func main() {
	INPUT := "sinput"
	grid := readLines(INPUT)
	solution(grid)
}

type Pos struct {
	r int
	c int
}

func (p *Pos) equals(o *Pos) bool {
	return p.r == o.r && p.c == o.c
}

func (p *Pos) distance(o Pos) int {
	return Abs(p.r-o.r) + Abs(p.c-o.c)
}

func solution(grid []string) {
	ROWS := len(grid)
	COLS := len(grid[0])
	emptyRows := []int{}
	emptyCols := []int{}
	galaxies := []Pos{}

	for c := 0; c < COLS; c++ {
		isColEmpty := true
		for r := 0; r < ROWS; r++ {
			h := []byte("#")[0]
			if Contains[int](emptyRows, r) == false &&
				strings.Contains(grid[r], "#") == false {
				emptyRows = append(emptyRows, r)
			} else if grid[r][c] == h {
				isColEmpty = false
				galaxies = append(galaxies, Pos{r, c})
			}
		}

		if isColEmpty {
			emptyCols = append(emptyCols, c)
		}
	}

	sum1 := galaxyDistanceSum(2, galaxies, emptyRows, emptyCols)
	sum2 := galaxyDistanceSum(1000000, galaxies, emptyRows, emptyCols)
	fmt.Println("PART 1", sum1)
	fmt.Println("PART 2", sum2)
}

func galaxyDistanceSum(
	expansionFactor int,
	galaxies []Pos,
	emptyRows, emptyCols []int,
) int {
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		p1 := galaxies[i]
		for j := i; j < len(galaxies); j++ {
			p2 := galaxies[j]
			d := p1.distance(p2)
			minr, maxr := getMinMax(p1.r, p2.r)
			minc, maxc := getMinMax(p1.c, p2.c)
			for _, r := range emptyRows {
				if r > minr && r < maxr {
					d += expansionFactor - 1
				}
			}
			for _, c := range emptyCols {
				if c > minc && c < maxc {
					d += expansionFactor - 1
				}
			}
			sum += d
		}
	}

	return sum
}
