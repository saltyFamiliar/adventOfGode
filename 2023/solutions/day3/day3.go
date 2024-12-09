package day3

import (
	"adventOfGode/toolbelt"
	"slices"
	"strconv"
	"unicode"
)

func IsSymbol(b byte) bool {
	return !unicode.IsDigit(rune(b)) && b != '.'
}

func consumeNum(y, x int, schematic [][]byte, consume bool) int {
	var rightBytes []byte
	var leftBytes []byte
	for i := x; i < len(schematic[y]) && unicode.IsDigit(rune(schematic[y][i])); i++ {
		rightBytes = append(rightBytes, schematic[y][i])
		if consume {
			schematic[y][i] = byte('.')
		}
	}
	for i := x - 1; i >= 0 && unicode.IsDigit(rune(schematic[y][i])); i-- {
		leftBytes = slices.Insert(leftBytes, 0, schematic[y][i])
		if consume {
			schematic[y][i] = byte('.')
		}
	}
	digits := string(append(leftBytes, rightBytes...))
	num, err := strconv.Atoi(digits)
	toolbelt.Must("parse number", err)
	return num
}

func consumeNumsAroundCoord(y, x int, schematic [][]byte, consume bool) (nums []int) {
	yStart := y - 1
	if y == 0 {
		yStart = y
	}
	yEnd := y + 1
	if y == len(schematic)-1 {
		yEnd = y
	}
	xStart := x - 1
	if x == 0 {
		xStart = x
	}
	xEnd := x + 1
	if x == len(schematic[0])-1 {
		xEnd = x
	}
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			r := rune(schematic[y][x])
			if unicode.IsDigit(r) {
				nums = append(nums, consumeNum(y, x, schematic, consume))
				for x <= xEnd && unicode.IsDigit(rune(schematic[y][x])) {
					x += 1
				}
			}
		}
	}
	return nums
}

func Solve1(matrix [][]byte) (partNumberSum int) {
	for y, line := range matrix {
		for x, b := range line {
			if !IsSymbol(b) {
				continue
			}
			for _, n := range consumeNumsAroundCoord(y, x, matrix, true) {
				partNumberSum += n
			}
		}
	}
	return partNumberSum
}

func Solve2(matrix [][]byte) (gearSum int) {
	for y, line := range matrix {
		for x, b := range line {
			if rune(b) != '*' {
				continue
			}

			nums := consumeNumsAroundCoord(y, x, matrix, false)
			if len(nums) == 2 {
				gearSum += nums[0] * nums[1]
			}
		}
	}

	return gearSum
}
