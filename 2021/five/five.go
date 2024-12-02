package five

import (
	. "adventOfGode/toolbelt"
	"fmt"
	"strconv"
	"strings"
)

func parsePtFromStr(pointStr string) (result Pt) {
	pointCoords := strings.Split(pointStr, ",")
	x, err := strconv.Atoi(pointCoords[0])
	if err != nil {
		fmt.Println("Couldn't convert first value to int")
	}
	y, err := strconv.Atoi(pointCoords[1])
	if err != nil {
		fmt.Println("Couldn't convert second value to int")
	}

	return Pt{X: x, Y: y}
}

func parseLine(line string) (from Pt, to Pt) {
	points := strings.Split(line, " -> ")
	return parsePtFromStr(points[0]), parsePtFromStr(points[1])
}

func PartOne(fileLines []string) (result int) {
	var ventPoints []Pt
	for _, line := range fileLines {
		from, to := parseLine(line)
		if from.Y == to.Y || from.X == to.X {
			ventPoints = append(ventPoints, PtsBetween(from, to)...)
		}
	}

	return len(FindDuplicatePts(ventPoints))
}

func PartTwo(fileLines []string) (result int) {
	var ventPoints []Pt
	for _, line := range fileLines {
		from, to := parseLine(line)
		ventPoints = append(ventPoints, PtsBetween(from, to)...)
	}

	return len(FindDuplicatePts(ventPoints))
}
