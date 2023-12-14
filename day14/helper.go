package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToInt(s string) int {
	s2n := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
	if s2n[s] != 0 || s == "0" {
		return s2n[s]
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return num
}

func readFile(fname string) string {
	data, err := os.ReadFile(fname)
	check(err)
	return strings.TrimSpace(string(data))
}

func readLines(fname string) []string {
	data := readFile(fname)
	lines := strings.Split(data, "\n")
	// popping off the blank line
	return lines
}

func stringTwoSplit(s string, sep string) (string, string) {
	split := strings.Split(s, sep)
	if len(split) < 2 {
		return "", ""
	}
	return split[0], split[1]
}

func reMatch(pattern string, s string) ([]string, [][]int, [][]string) {
	r, _ := regexp.Compile(pattern)
	return r.FindAllString(s, -1), r.FindAllStringIndex(s, -1), r.FindAllStringSubmatch(s, -1)
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func getMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getMapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for k := range m {
		v := m[k]
		values = append(values, v)
	}
	return values
}

func stringToIntSlice(s []string) []int {
	var iSlice []int

	for _, v := range s {
		cleanVal := strings.TrimSpace(v)
		if cleanVal == "" {
			continue
		}
		iSlice = append(iSlice, convertToInt(cleanVal))
	}

	return iSlice
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func getMinMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func gridPrint[T any](g [][]T) {
	R := len(g)
	C := len(g[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			fmt.Print(g[r][c])
		}
		fmt.Println()
	}

	fmt.Println()
}

func gridTranspose[T any](g [][]T) [][]T {
	R := len(g)
	C := len(g[0])
	ng := make([][]T, C)
	for c := 0; c < C; c++ {
		row := make([]T, C)
		for r := 0; r < R; r++ {
			row[r] = g[r][c]
		}
		ng[c] = row
	}

	return ng
}

func reverseSlice[T any](s []T) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

func sortString(input string, reverse bool) string {
	// Convert the string to a slice of characters
	chars := strings.Split(input, "")
	// Sort the slice of characters
	sort.Strings(chars)
	// Convert the sorted slice of characters back to a string
	if reverse {
		reverseSlice[string](chars)
	}
	sortedString := strings.Join(chars, "")
	return sortedString
}
