package main

import (
	"container/heap"
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

var R, C int

type Pos struct {
	r  int // row index
	c  int // col index
	dr int // row dir
	dc int // col dir
	n  int // num steps
	h  int // heat total
}

func (p *Pos) next(grid [][]int) []Pos {
	if p.n < 3 && (p.dr != 0 || p.dc != 0) {
		nr, nc := p.r+p.dr, p.c+p.dc
		if nr >= 0 && nr < R && nc >= 0 && nc < C {
			// ch := grid[nr][nc]
			return []Pos{{nr, nc, p.dr, p.dc, p.n + 1, p.h + grid[nr][nc]}}
		}
	}
	pos := []Pos{}
	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		dr, dc := dir[0], dir[1]
		sameDir := dr == p.dr && dc == p.dc
		revDir := dr == -p.dr && dc == -p.dc
		// if dir is same or reverse, skip
		if !sameDir && !revDir {
			nr, nc := p.r+dr, p.c+dc
			if nr >= 0 && nr < R && nc >= 0 && nc < C {
				pos = append(pos, Pos{nr, nc, dr, dc, 1, p.h + grid[nr][nc]})
			}
		}
	}

	return pos
}

func (p *Pos) key() string {
	return (strconv.Itoa(p.r) + "," + strconv.Itoa(p.c) + "," +
		strconv.Itoa(p.dr) + "," + strconv.Itoa(p.dc)) +
		"," + strconv.Itoa(p.n)
}

func (p *Pos) getHeat(grid [][]int) {
	queue := []Pos{*p}
	seen := make(map[string]int)
	cur := Pos{}

	for len(queue) > 0 {
		cur, queue = pqPop(queue)
		if cur.r == R-1 && cur.c == C-1 {
			fmt.Println("Found", cur.h)
			break
		}
		if seen[cur.key()] == 1 {
			continue
		}

		seen[cur.key()] = 1
		if cur.n < 3 && !(cur.dr == 0 && cur.dc == 0) {
			nr := cur.r + cur.dr
			nc := cur.c + cur.dc
			if nr >= 0 && nr < R && nc >= 0 && nc < C {
				h := cur.h + grid[cur.r][cur.c]
				queue = append(queue, Pos{nr, nc, cur.dr, cur.dc, cur.n + 1, h})
			}

		}

		for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			ndr, ndc := dir[0], dir[1]
			sameDir := ndr == p.dr && ndc == p.dc
			revDir := ndr == -p.dr && ndc == -p.dc
			// if dir is same or reverse, skip
			if sameDir || revDir {
				continue
			}
			nr := cur.r + ndr
			nc := cur.c + ndc
			if nr >= 0 && nr < R && nc >= 0 && nc < C {
				h := cur.h + grid[cur.r][cur.c]
				queue = append(queue, Pos{nr, nc, ndr, ndc, 1, h})

			}
		}

		// next := cur.next(grid)
		// for _, np := range next {
		// 	queue = append(queue, np)
		// }

		// fmt.Println(queue)
	}
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func pqPop(q []Pos) (Pos, []Pos) {
	minPos := Pos{h: MaxInt}
	minI := -1
	for i, p := range q {
		if p.h < minPos.h {
			minPos = p
			minI = i
		}
	}

	return minPos, append(q[:minI], q[minI+1:]...)
}

// func solve(lines []string) {
// 	grid := make([][]int, len(lines))
// 	for i, line := range lines {
// 		grid[i] = stringToIntSlice(strings.Split(line, ""))
// 	}
//
// 	R = len(grid)
// 	C = len(grid[0])
// 	start := Pos{} // golang sets everything to zero
// 	start.getHeat(grid)
// }

// Define a PriorityQueue type which implements the heap.Interface
type PriorityQueue []*Pos

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Here, we're prioritizing based on the 'priority' field
	return pq[i].h < pq[j].h
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	p := x.(*Pos)
	*pq = append(*pq, p)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func solve(lines []string) {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = stringToIntSlice(strings.Split(line, ""))
	}

	R = len(grid)
	C = len(grid[0])
	start := Pos{0, 0, 0, 0, 0, 0} // golang sets everything to zero

	seen := make(map[string]int)
	// q := []Pos{start}
	// cur := Pos{}

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &start)

	i := 0
	// for len(q) > 0 {
	for pq.Len() > 0 {
		i++
		cur := heap.Pop(&pq).(*Pos)
		// cur, q = pqPop(q)
		r, c, dr, dc, n, h := cur.r, cur.c, cur.dr, cur.dc, cur.n, cur.h

		if r == R-1 && c == C-1 {
			fmt.Println("Reached", h)
		}

		k := cur.key()

		if seen[k] == 1 {
			continue
		}

		seen[k] = 1

		if n < 3 && (dr != 0 || dc != 0) {
			nr := r + dr
			nc := r + dc
			if nr >= 0 && nr < R && nc >= 0 && nc < C {
				h += grid[nr][nc]

				// q = append(q, Pos{nr, nc, dr, dc, n + 1, h})
				heap.Push(&pq, &Pos{nr, nc, dr, dc, n + 1, h})
			}
		}

		for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			ndr, ndc := dir[0], dir[1]
			sameDir := ndr == dr && ndc == dc
			revDir := ndr == -dr && ndc == -dc
			// if sameDir || revDir {
			// 	continue
			// }
			if !sameDir && !revDir {
				nr := r + ndr
				nc := r + ndc
				if nr >= 0 && nr < R && nc >= 0 && nc < C {
					h += grid[nr][nc]
					// fmt.Println(dr, dc, ndr, ndc, h)
					// q = append(q, Pos{nr, nc, ndr, ndc, 1, h})
					heap.Push(&pq, &Pos{nr, nc, ndr, ndc, 1, h})
				}
			}
		}

		// break
	}
}
