package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	lines := readLines(*f)
	solve(lines)
}

type Pos struct {
	r  int
	c  int
	dr int
	dc int
}

func (p *Pos) dirIsVertical() bool {
	return p.dr != 0
}

func (p *Pos) dirIsHorizontal() bool {
	return p.dc != 0
}

func (p *Pos) next(grid [][]string) []Pos {
	R := len(grid)
	C := len(grid[0])
	nr, nc := p.r+p.dr, p.c+p.dc
	if nr >= 0 && nr < R && nc >= 0 && nc < C {
		ch := grid[nr][nc]
		if ch == "." || (ch == "|" && p.dr != 0) || (ch == "-" && p.dc != 0) {
			return []Pos{{nr, nc, p.dr, p.dc}}
		} else if ch == "/" {
			ndr, ndc := -p.dc, -p.dr
			return []Pos{{nr, nc, ndr, ndc}}
		} else if ch == "\\" {
			ndr, ndc := p.dc, p.dr
			return []Pos{{nr, nc, ndr, ndc}}
		} else if ch == "|" {
			return []Pos{{nr, nc, 1, 0}, {nr, nc, -1, 0}}
		} else if ch == "-" {
			return []Pos{{nr, nc, 0, 1}, {nr, nc, 0, -1}}
		}
	}
	return []Pos{}
}

func (p *Pos) key() string {
	return (strconv.Itoa(p.r) + "," + strconv.Itoa(p.c) + "," +
		strconv.Itoa(p.dr) + "," + strconv.Itoa(p.dc))
}

func (p *Pos) energizedTilesCount(grid [][]string) int {
	queue := []Pos{*p}
	seen := make(map[string]int)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		next := p.next(grid)
		for _, np := range next {
			if seen[np.key()] != 1 {
				queue = append(queue, np)
				seen[np.key()] = 1
			}
		}
	}

	points := make(map[string]int)

	for s := range seen {
		t := strings.Split(s, ",")
		points[t[0]+","+t[1]] = 1
	}

	return len(points)
}

func solve(lines []string) {
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}
	R := len(grid)
	C := len(grid[0])

	start := Pos{0, -1, 0, 1}
	fmt.Println("Part 1:", start.energizedTilesCount(grid))

	maxCount := 0
	count := 0
	for r := 0; r < R; r++ {
		// for left side
		start = Pos{r, -1, 0, 1}
		count = start.energizedTilesCount(grid)
		maxCount = max(maxCount, count)
		// for right side
		start = Pos{r, C, 0, -1}
		count = start.energizedTilesCount(grid)
		maxCount = max(maxCount, count)
	}

	// check for columns
	for c := 0; c < C; c++ {
		// for down
		start = Pos{-1, c, 1, 0}
		count = start.energizedTilesCount(grid)
		maxCount = max(maxCount, count)
		// for up
		start = Pos{R, c, -1, 0}
		count = start.energizedTilesCount(grid)
		maxCount = max(maxCount, count)
	}

	fmt.Println("Part 2:", maxCount)
}
