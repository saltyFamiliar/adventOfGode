package day9

import (
	"adventOfGode/2023/ergo"
	"adventOfGode/toolbelt"
	"fmt"
	"strings"
)

func differences(seq []int) func(func(int) bool) {
	return func(yield func(int) bool) {
		if len(seq) < 2 {
			return
		}
		for i := 1; i < len(seq); i++ {
			if !yield(int(float64(seq[i] - seq[i-1]))) {
				return
			}
		}
	}
}

func lineDifferences(line []int) []int {
	diffs := make([]int, 0, len(line))
	for d := range differences(line) {
		diffs = append(diffs, d)
	}
	return diffs
}

func predictNextValue(line []int) (prediction int) {
	newDiffs := make([]int, len(line))
	copy(newDiffs, line)
	for !toolbelt.All(newDiffs, 0) {
		newDiffs = lineDifferences(newDiffs)
		if len(newDiffs) == 0 {
			break
		}
		prediction += newDiffs[len(newDiffs)-1]
	}
	prediction += line[len(line)-1]
	return prediction
}

func predictPrevValue(line []int) (prediction int) {
	newDiffs := make([]int, len(line))
	copy(newDiffs, line)
	firstDigits := []int{newDiffs[0]}
	for !toolbelt.All(newDiffs, 0) {
		newDiffs = lineDifferences(newDiffs)
		if len(newDiffs) == 0 {
			break
		}
		firstDigits = append(firstDigits, newDiffs[0])
	}

	prediction = firstDigits[len(firstDigits)-1]
	for i := len(firstDigits) - 2; i >= 0; i-- {
		prediction = firstDigits[i] - prediction
	}

	fmt.Println(prediction)
	return prediction
}

func Solve1() (answer int) {
	for _, line := range ergo.GetFileLines("solutions/day9/input.txt") {
		lineNumeric := toolbelt.StrArrToInts(strings.Fields(line))
		answer += predictNextValue(lineNumeric)
	}
	return answer
}

func Solve2() (answer int) {
	for _, line := range ergo.GetFileLines("solutions/day9/input.txt") {
		lineNumeric := toolbelt.StrArrToInts(strings.Fields(line))
		answer += predictPrevValue(lineNumeric)
	}
	return answer
}
