package day2

import (
	"adventOfGode/toolbelt"
	"strconv"
	"strings"
)

var cubeAmounts = map[string]int64{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func roundIsValid(draws []string) bool {
	for _, draw := range draws {
		drawParts := strings.Fields(draw)
		drawDigit, drawColor := drawParts[0], drawParts[1]
		drawNum, err := strconv.ParseInt(drawDigit, 10, 0)
		toolbelt.Must("parse draw number", err)
		if cubeAmounts[drawColor] < drawNum {
			return false
		}
	}
	return true
}

func gameIsValid(rounds []string) bool {
	for _, round := range rounds {
		draws := strings.Split(round, ", ")
		if !roundIsValid(draws) {
			return false
		}
	}
	return true
}

func Solve1(lines []string) (sumOfPossible int) {
	for gameNum, line := range lines {
		gameNum += 1
		lineParts := strings.Split(line, ":")
		rounds := strings.Split(lineParts[1], ";")
		if gameIsValid(rounds) {
			sumOfPossible += gameNum
		}
	}
	return sumOfPossible
}

func getMinCubeSet(rounds []string) map[string]int64 {
	cubeSet := map[string]int64{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, round := range rounds {
		draws := strings.Split(round, ", ")
		for _, draw := range draws {
			drawParts := strings.Fields(draw)
			drawColor := drawParts[1]
			drawNum, err := strconv.ParseInt(drawParts[0], 10, 0)
			toolbelt.Must("parse number", err)
			if cubeSet[drawColor] < drawNum {
				cubeSet[drawColor] = drawNum
			}
		}
	}
	return cubeSet
}

func Solve2(lines []string) (sumOfPowers int64) {
	for _, line := range lines {
		lineParts := strings.Split(line, ":")
		rounds := strings.Split(lineParts[1], ";")
		minCubeSet := getMinCubeSet(rounds)
		power := int64(1)
		for _, minAmt := range minCubeSet {
			power *= minAmt
		}
		sumOfPowers += power
	}
	return sumOfPowers
}
