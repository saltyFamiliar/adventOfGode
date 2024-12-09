package utils

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	Y int
	X int
}

func GetLines(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func Map[T any, U any](s []T, fn func(T) U) []U {
	var result []U
	for _, v := range s {
		result = append(result, fn(v))
	}
	return result
}

func MapErr[T any, U any](fn func(T) (U, error), s []T) []U {
	var result []U
	for _, v := range s {
		r, err := fn(v)
		if err != nil {
			panic("Failed to apply function to element: " + err.Error())
		}
		result = append(result, r)
	}
	return result
}

func GetGrid(file string) [][]string {
	lines := GetLines(file)
	var matrix [][]string
	for _, l := range lines {
		matrix = append(matrix, strings.Split(l, ""))
	}
	return matrix
}

func PrintGrid(grid [][]string) {
	var labelsX []int

	for i := range len(grid[0]) {
		labelsX = append(labelsX, i%10)
	}

	fmt.Println(" ", labelsX)
	for i, l := range grid {
		fmt.Println(i%10, l)
	}
}

func GridIter(grid [][]string) func(func(pt Point, ch string) bool) {
	return func(yield func(pt Point, ch string) bool) {
		for y, row := range grid {
			for x := range row {
				if !yield(Point{y, x}, grid[y][x]) {
					return
				}
			}
		}
	}
}

func PtInBounds(pt Point, mat [][]string) bool {
	return pt.Y >= 0 && pt.Y < len(mat) && pt.X >= 0 && pt.X < len(mat[0])
}
