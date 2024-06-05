package day6

import (
	"adventOfGode/2023/ergo"
	"strings"
)

func distFunc(totalTime, timeHeld int) int {
	return (totalTime - timeHeld) * timeHeld
}

func Solve1() int {
	scanner := ergo.GetInputScanner("solutions/day6/input.txt")

	scanner.Scan()
	timesLine := scanner.Text()
	_, timesStr, _ := strings.Cut(timesLine, ":")
	times := strings.Fields(timesStr)

	scanner.Scan()
	distancesLine := scanner.Text()
	_, distancesStr, _ := strings.Cut(distancesLine, ":")
	distances := strings.Fields(distancesStr)

	result := 1
	for i, totalTime := range times {
		waysToWin := 0
		distance := ergo.EzIntParse(distances[i])
		totalTime := ergo.EzIntParse(totalTime)
		for j := 0; j < totalTime; j++ {
			if distFunc(totalTime, j) > distance {
				waysToWin++
			}
		}
		result *= waysToWin
	}

	return result
}

func Solve2() int {
	scanner := ergo.GetInputScanner("solutions/day6/input.txt")

	scanner.Scan()
	timesLine := scanner.Text()
	_, timesStr, _ := strings.Cut(timesLine, ":")
	times := strings.Fields(timesStr)
	time := strings.Join(times, "")
	totalTime := ergo.EzIntParse(time)

	scanner.Scan()
	distancesLine := scanner.Text()
	_, distancesStr, _ := strings.Cut(distancesLine, ":")
	distances := strings.Fields(distancesStr)
	distance := strings.Join(distances, "")
	record := ergo.EzIntParse(distance)

	waysToWin := 0
	for i := 0; i < totalTime; i++ {
		if distFunc(totalTime, i) > record {
			waysToWin++
		}
	}

	return waysToWin
}
