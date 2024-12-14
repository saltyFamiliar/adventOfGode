package main

import (
	. "aoc_2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func safe(levels []int) bool {
	for i, n := range levels[1:] {
		if Abs(n-levels[i]) > 3 || Abs(n-levels[i]) < 1 {
			return false
		}
	}
	return true
}

func solve(reports []string) (silver, gold int) {
	for _, r := range reports {
		levels := MapErr(strconv.Atoi, strings.Split(r, " "))
		levelsB := append([]int{}, levels...)
		slices.Reverse(levelsB)
		if (slices.IsSorted(levels) || slices.IsSorted(levelsB)) && safe(levels) {
			silver, gold = silver+1, gold+1
		} else {
			for i := 0; i < len(levels); i++ {
				butOne := append([]int{}, levels...)
				butOne = append(butOne[:i], butOne[i+1:]...)
				butOneB := append([]int{}, butOne...)
				slices.Reverse(butOneB)
				if (slices.IsSorted(butOne) || slices.IsSorted(butOneB)) && safe(butOne) {
					gold++
					break
				}
			}
		}
	}
	return silver, gold
}

func main() {
	grid := GetLines("two/input.txt")
	silver, gold := solve(grid)
	fmt.Printf("silver: %d\ngold: %d\n", silver, gold)
}
