package day1

import (
	"adventOfGode/aocio"
	"testing"
)

var lines []string

func init() {
	lines = aocio.FileToLines("input1.txt")
}

func BenchmarkSolve1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve1(lines)
	}
}
