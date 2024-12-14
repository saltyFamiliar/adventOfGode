package main

import (
	. "aoc_2024/utils"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

func solve(input string) (int, int) {
	line := strings.Fields(input)
	stones := Map(line, func(s string) int {
		return Must(strconv.Atoi(s))
	})
	stoneMap := Count(stones)
	var silver int
	for i := 0; i < 75; i++ {
		newStones := make(map[int]int)
		for stone, count := range stoneMap {
			numDigits := int(math.Log10(float64(stone))) + 1
			if stone == 0 {
				newStones[1] += count
			} else if numDigits%2 == 0 {
				div := int(math.Pow10(numDigits / 2))
				stoneA, stoneB := stone/div, stone%div
				newStones[stoneA] += count
				newStones[stoneB] += count
			} else {
				newStones[stone*2024] += count
			}
		}
		stoneMap = newStones
		if i == 24 {
			silver = SumInt(slices.Collect(maps.Values(stoneMap)))
		}
	}
	return silver, SumInt(slices.Collect(maps.Values(stoneMap)))
}

func main() {
	silver, gold := solve(GetLines("input.txt")[0])
	fmt.Printf("Silver: %d\nGold: %d\n", silver, gold)
}
