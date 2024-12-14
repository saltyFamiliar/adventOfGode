package main

import (
	. "aoc_2024/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solve(input []string) (int, int) {
	var left, right []int
	for _, line := range input {
		l, r, _ := strings.Cut(line, "   ")
		left = append(left, Must(strconv.Atoi(l)))
		right = append(right, Must(strconv.Atoi(r)))
	}
	sort.Ints(left)
	sort.Ints(right)
	counts := Count(right)
	var silver, gold int
	for i := 0; i < len(left); i++ {
		silver += Abs(left[i] - right[i])
		gold += left[i] * counts[left[i]]
	}
	return silver, gold
}

func main() {
	silver, gold := solve(GetLines("input.txt"))
	fmt.Printf("Silver: %d\nGold: %d\n", silver, gold)
}
