package day1

import (
	"adventOfGode/toolbelt"
	"fmt"
)

func Solve1(lines []string) (totalValue int) {
	for _, line := range lines {
		var lineValue int
		for _, b := range line {
			if d, ok := toolbelt.RuneToIntIfDigit(b); ok {
				lineValue = d * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if d, ok := toolbelt.RuneToIntIfDigit(rune(line[i])); ok {
				lineValue += d
				break
			}
		}
		totalValue += lineValue
	}
	return totalValue
}

var numberNames = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}
var nameLenMin, nameLenMax = 3, 5

func isNumberNameOrDigit(line string, i int) (n int, ok bool) {
	if d, ok := toolbelt.ByteToIntIfDigit(line[i]); ok {
		return d, ok
	}
	for j := i + nameLenMin; j <= i+nameLenMax && j <= len(line); j++ {
		possibleName := line[i:j]
		fmt.Println(possibleName)
		for d, name := range numberNames {
			if possibleName == name {
				return d + 1, true
			}
		}
	}
	return -1, false
}

func Solve2(lines []string) (totalValue int) {
	for _, line := range lines {
		var lineValue int
		for i := range line {
			if n, ok := isNumberNameOrDigit(line, i); ok {
				lineValue = n * 10
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if n, ok := isNumberNameOrDigit(line, i); ok {
				lineValue += n
				break
			}
		}
		totalValue += lineValue
	}

	return totalValue
}
