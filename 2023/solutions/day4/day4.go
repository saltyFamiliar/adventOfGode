package day4

import (
	"2023/ergo"
	"slices"
	"strings"
)

func Solve1() (totalScore int) {
	scanner := ergo.GetInputScanner("solutions/day4/input.txt")

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ": ")
		cardNums := lineParts[1]
		winners, yours, ok := strings.Cut(cardNums, " | ")
		if !ok {
			panic("Separator not found")
		}

		winnerSlc, yoursSlc := strings.Fields(winners), strings.Fields(yours)
		cardScore := 0
		for _, yourNum := range yoursSlc {
			if slices.Contains(winnerSlc, yourNum) {
				if cardScore == 0 {
					cardScore += 1
				} else {
					cardScore = cardScore * 2
				}
			}
		}
		totalScore += cardScore
	}

	return totalScore
}

func Solve2() (totalScore int) {
	scanner := ergo.GetInputScanner("solutions/day4/input.txt")

	cardMap := map[int]int{}

	for gameNum := 1; scanner.Scan(); gameNum++ {
		if _, ok := cardMap[gameNum]; !ok {
			cardMap[gameNum] = 1
		}

		line := scanner.Text()
		lineParts := strings.Split(line, ": ")
		cardNums := lineParts[1]
		winners, yours, ok := strings.Cut(cardNums, " | ")
		if !ok {
			panic("Separator not found")
		}

		winnerSlc, yoursSlc := strings.Fields(winners), strings.Fields(yours)
		cardScore := 0
		for _, yourNum := range yoursSlc {
			if slices.Contains(winnerSlc, yourNum) {
				cardScore += 1
			}
		}

		for i := gameNum + 1; i <= gameNum+cardScore; i++ {
			if _, ok := cardMap[i]; !ok {
				cardMap[i] = 1 + cardMap[gameNum]
			} else {
				cardMap[i] += cardMap[gameNum]
			}
		}

		totalScore += cardMap[gameNum]
	}
	return totalScore
}
