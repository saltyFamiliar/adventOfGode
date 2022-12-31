package main

import (
	"adventOfGode/toolbelt"
	"strconv"
)

func partOne(fileLines []string) (depthIncreases int) {
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

func partTwo(fileLines []string) (depthIncreases int) {
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

func main() {
	dirPath := "2021/one"
	toolbelt.TestPart(partOne, dirPath, 7)
	toolbelt.TestPart(partTwo, dirPath, 5)
}
