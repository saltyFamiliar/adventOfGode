package main

import (
	"adventOfGode/2023/solutions/day9"
	"adventOfGode/aocio"
	"fmt"
	"golang.org/x/exp/constraints"
)

type SolverInput interface {
	[]string | [][]byte
}

type SolverFunc[SI SolverInput, T constraints.Integer] interface {
	func(SI) T
}

func solveFuncAdapter[SF SolverFunc[SI, T], T constraints.Integer, SI SolverInput](sf SF, filePath string) {
	var result T
	switch f := any(sf).(type) {
	case func([]string) T:
		lines := aocio.FileToLines(filePath)
		result = f(lines)
	case func([][]byte) T:
		lines := aocio.FileToByteMatrix(filePath)
		result = f(lines)
	default:
		panic("function type unsupported")
	}

	fmt.Println("The answer is ", result)
}

func main() {
	solveFuncAdapter(day9.Solve2, "solutions/day9/input.txt")
}
