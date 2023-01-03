package six

import (
	. "adventOfGode/toolbelt"
	"strings"
)

func updateFishTimerCounts(timers map[int]int) {
	for dayNum := 0; dayNum < 9; dayNum += 1 {
		timers[dayNum-1] = timers[dayNum]
	}
	timers[6] += timers[-1]
	timers[8] = timers[-1]
	timers[-1] = 0
}

func PartOne(fileLines []string) int {
	fishTimers := StrArrToIntMap(strings.Split(fileLines[0], ","))

	for day := 0; day < 80; day += 1 {
		updateFishTimerCounts(fishTimers)
	}

	return MapValSum(fishTimers)
}

func PartTwo(fileLines []string) int {
	fishTimers := StrArrToIntMap(strings.Split(fileLines[0], ","))

	for day := 0; day < 256; day += 1 {
		updateFishTimerCounts(fishTimers)
	}

	return MapValSum(fishTimers)
}
