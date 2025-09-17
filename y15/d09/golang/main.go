package main

import (
	"flag"
	"fmt"
	"slices"
	"strings"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	lines := readLines(*f)
	solve(lines)
}

func permute[T any](arr []T, i int, result *[][]T) {
	if i == len(arr) {
		// copy the input slice
		permutation := append([]T{}, arr...)
		*result = append(*result, permutation)
		return
	}

	for j := i; j < len(arr); j++ {
		// swap
		arr[i], arr[j] = arr[j], arr[i]
		// permute subtree
		permute(arr, i+1, result)
		// restore
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func solve(lines []string) {
	graph := make(map[string]map[string]int)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		u, v, d := parts[0], parts[2], convertToInt(parts[len(parts)-1])
		if graph[u] == nil {
			graph[u] = make(map[string]int)
		}
		if graph[v] == nil {
			graph[v] = make(map[string]int)
		}
		graph[u][v] = d
		graph[v][u] = d
	}

	nodes := getMapKeys(graph)
	routes := make([][]string, 0)
	permute(nodes, 0, &routes)

	distances := make([]int, 0)

	for _, route := range routes {
		routeDist := 0
		for i := 0; i < len(route)-1; i++ {
			u, v := route[i], route[i+1]
			routeDist += graph[u][v]
		}
		distances = append(distances, routeDist)
	}

	part1 := slices.Min(distances)
	fmt.Println("Part 1", part1)

	part2 := slices.Max(distances)
	fmt.Println("Part 2", part2)
}
