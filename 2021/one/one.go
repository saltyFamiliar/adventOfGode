package one

import (
	"adventOfGode/toolbelt"
	"strconv"
)

func PartOne(fileLines []string) (depthIncreases int) {
	depth, _ := strconv.Atoi(fileLines[0])
	for _, depthStr := range fileLines[1:] {
		newDepth, _ := strconv.Atoi(depthStr)
		if newDepth > depth {
			depthIncreases += 1
		}
		depth = newDepth
	}
	return depthIncreases
}

func PartTwo(fileLines []string) (depthIncreases int) {
	depths := toolbelt.StrArrToInts(fileLines)
	lastDepth := depths[0] + depths[1] + depths[2]
	newDepth := lastDepth
	for i, depth := range depths[3:] {
		newDepth += depth - depths[i]
		if newDepth > lastDepth {
			depthIncreases += 1
		}
		lastDepth = newDepth
	}
	return depthIncreases
}
