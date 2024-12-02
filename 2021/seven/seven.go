package seven

import (
	"adventOfGode/toolbelt"
	"math"
	"sort"
	"strings"
)

func median(arr []int) int {
	sort.Ints(arr)
	numInts := len(arr)
	if numInts%2 == 0 {
		leftPosition := arr[numInts/2]
		rightPosition := arr[(numInts/2)+1]

		leftOccurrences := 0
		rightOccurrences := 0
		for _, pos := range arr {
			if pos == leftPosition {
				leftOccurrences += 1
			}
			if pos == rightPosition {
				rightOccurrences += 1
			}
		}

		if leftOccurrences > rightOccurrences {
			return leftPosition
		} else {
			return rightPosition
		}

	} else {
		return arr[numInts/2]
	}
}

func PartOne(fileLines []string) (fuelUsage int) {
	positions := toolbelt.StrArrToInts(strings.Split(fileLines[0], ","))
	medianPos := median(positions)

	for _, pos := range positions {
		fuelUsage += int(math.Abs(float64(pos) - float64(medianPos)))
	}

	return fuelUsage
}

func PartTwo(fileLines []string) (result int) {

	return
}

// 0, 1, 1, 2, 2, 2, 4, 7, 14, 16
