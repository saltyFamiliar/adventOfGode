package day7

import (
	"adventOfGode/2023/ergo"
	"slices"
	"strings"
)

type round struct {
	hand string
	bid  int
}

func (r round) getHandRank(jacksWild bool) (unique int) {
	set := make(map[int32]int)
	jacks := 0
	for _, ch := range r.hand {
		if _, ok := set[ch]; ok {
			set[ch] += 1
		} else if jacksWild && ch == 'J' {
			jacks += 1
		} else {
			set[ch] = 1
			unique += 1
		}
	}
	if jacksWild {
		if jacks == 5 {
			return 10
		}
		var modeCard int32
		maxCount := 0
		for card, count := range set {
			if count > maxCount {
				maxCount = count
				modeCard = card
			}
		}
		set[modeCard] += jacks
	}

	if unique == 1 {
		return 10
	} else if unique == 2 {
		for _, v := range set {
			if v == 1 || v == 4 {
				return 9
			} else {
				return 8
			}
		}
	} else if unique == 3 {
		for _, v := range set {
			if v == 2 {
				return 6
			}
		}
		return 7
	} else if unique == 4 {
		return 5
	}

	return 4
}

func Solve1() (totalScore int) {
	scanner := ergo.GetInputScanner("solutions/day7/input.txt")
	var rounds []round

	faceCards := map[uint8]uint8{
		'T': 100,
		'J': 101,
		'Q': 102,
		'K': 103,
		'A': 104,
	}

	for scanner.Scan() {
		hand, bidStr, _ := strings.Cut(scanner.Text(), " ")

		rounds = append(rounds, round{hand, ergo.EzIntParse(bidStr)})
	}

	slices.SortFunc(rounds, func(a, b round) int {
		aRank, bRank := a.getHandRank(false), b.getHandRank(false)
		if aRank > bRank {
			return -1
		} else if aRank < bRank {
			return 1
		}

		for i, _ := range a.hand {
			aCard, bCard := a.hand[i], b.hand[i]
			if aCard > 60 {
				aCard = faceCards[aCard]
			}
			if bCard > 60 {
				bCard = faceCards[bCard]
			}

			if aCard > bCard {
				return -1
			} else if aCard < bCard {
				return 1
			}
		}
		return 0
	})

	totalRounds := len(rounds)
	for i, r := range rounds {
		totalScore += (totalRounds - i) * r.bid
	}

	return totalScore
}

func Solve2() (totalScore int) {
	scanner := ergo.GetInputScanner("solutions/day7/input.txt")
	var rounds []round

	faceCards := map[uint8]uint8{
		'T': 100,
		'J': 0,
		'Q': 102,
		'K': 103,
		'A': 104,
	}

	for scanner.Scan() {
		hand, bidStr, _ := strings.Cut(scanner.Text(), " ")

		rounds = append(rounds, round{hand, ergo.EzIntParse(bidStr)})
	}

	slices.SortFunc(rounds, func(a, b round) int {
		aRank, bRank := a.getHandRank(true), b.getHandRank(true)
		if aRank > bRank {
			return -1
		} else if aRank < bRank {
			return 1
		}

		for i, _ := range a.hand {
			aCard, bCard := a.hand[i], b.hand[i]
			if aCard > 60 {
				aCard = faceCards[aCard]
			}
			if bCard > 60 {
				bCard = faceCards[bCard]
			}

			if aCard > bCard {
				return -1
			} else if aCard < bCard {
				return 1
			}
		}
		return 0
	})

	totalRounds := len(rounds)
	for i, r := range rounds {
		totalScore += (totalRounds - i) * r.bid
	}

	return totalScore
}
