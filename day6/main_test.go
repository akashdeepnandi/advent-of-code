package main

import (
	"testing"
)

var num = 1000

func BenchmarkQuadratic(b *testing.B) {
	lines := readLines("input")
	for i := 0; i < b.N; i++ {
		solveQuadratic(lines)
	}
}

func BenchmarkBruteForce(b *testing.B) {
	lines := readLines("input")
	for i := 0; i < b.N; i++ {
		solveBruteForce(lines)
	}
}
