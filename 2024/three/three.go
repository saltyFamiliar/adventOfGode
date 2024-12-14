package main

import (
	. "aoc_2024/utils"
	"fmt"
	"regexp"
)

func solve(lines []string) (silver, gold int) {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
	able := true
	for _, l := range lines {
		for _, m := range re.FindAllString(l, -1) {
			if m == "do()" {
				able = true
			} else if m == `don't()` {
				able = false
			} else {
				var a, b int
				Must(fmt.Sscanf(m, "mul(%d,%d)", &a, &b))
				silver += a * b
				if able {
					gold += a * b
				}
			}
		}
	}
	return silver, gold
}

func main() {
	grid := GetLines("three/input.txt")
	silver, gold := solve(grid)
	fmt.Printf("silver: %d\ngold: %d\n", silver, gold)
}
