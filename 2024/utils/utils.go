package utils

import (
	"fmt"
	"os"
	"strings"
)

var (
	UP    = Point{-1, 0}
	RIGHT = Point{0, 1}
	DOWN  = Point{1, 0}
	LEFT  = Point{0, -1}
)

var (
	DIRS_CLOCKWISE = []Point{UP, RIGHT, DOWN, LEFT}
)

func SplitAtIndex(s string, i int) (string, string) {
	return s[:i], s[i:]
}

func Count[T comparable](t []T) map[T]int {
	counts := make(map[T]int)
	for _, element := range t {
		counts[element]++
	}
	return counts
}

func SumInt(ints []int) (sum int) {
	for _, i := range ints {
		sum += i
	}
	return sum
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Abs[T Ordered](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

type Grid [][]string

func (grid Grid) At(pt Point) string {
	return grid[pt.Y][pt.X]
}

func (grid *Grid) Set(pt Point, new string) string {
	old := (*grid)[pt.Y][pt.X]
	(*grid)[pt.Y][pt.X] = new
	return old
}

func (grid *Grid) Row(y int) []string {
	return (*grid)[y]
}

func (grid Grid) RowCopy(y int) []string {
	cpy := make([]string, len(grid[y]))
	copy(cpy, grid[y])
	return cpy
}

type Point struct {
	Y int
	X int
}

func (pt Point) Add(dy, dx int) Point {
	return Point{pt.Y + dy, pt.X + dx}
}

func (pt Point) AddPt(dPt Point) Point {
	return Point{pt.Y + dPt.Y, pt.X + dPt.X}
}

func Must[T any](t T, e error) T {
	if e != nil {
		panic(e.Error())
	}
	return t
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

func ForEach[T any](s []T, fn func(T)) {
	for _, v := range s {
		fn(v)
	}
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

func GridFrom(file string) Grid {
	lines := GetLines(file)
	var grid Grid
	for _, l := range lines {
		grid = append(grid, strings.Split(l, ""))
	}
	return grid
}

func PrintGrid(grid Grid) {
	var labelsX []int

	for i := range len(grid[0]) {
		labelsX = append(labelsX, i%10)
	}

	fmt.Println(" ", labelsX)
	for i, l := range grid {
		fmt.Println(i%10, l)
	}
}

func GridIter(grid Grid) func(func(pt Point, ch string) bool) {
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

func (pt Point) In(grid Grid) bool {
	return pt.Y >= 0 && pt.Y < len(grid) && pt.X >= 0 && pt.X < len(grid[0])
}
