package main

import (
	. "aoc_2024/utils"
	"fmt"
)

func solve(lines []string) (silver, gold int) {
	for i := 0; i < len(lines); i += 4 {
		var ax, ay, bx, by, px, py float64
		Must(fmt.Sscanf(lines[i], "Button A: X+%f, Y+%f", &ax, &ay))
		Must(fmt.Sscanf(lines[i+1], "Button B: X+%f, Y+%f", &bx, &by))
		Must(fmt.Sscanf(lines[i+2], "Prize: X=%f, Y=%f", &px, &py))

		a := ((px * by) - (py * bx)) / ((by * ax) - (ay * bx))
		if float64(int(a)) == a {
			b := int((px - (ax * a)) / bx)
			silver = silver + (int(a) * 3) + b
		}
		px += 10000000000000
		py += 10000000000000
		a = ((px * by) - (py * bx)) / ((by * ax) - (ay * bx))
		if float64(int(a)) == a {
			b := (px - (ax * a)) / bx
			if float64(int(b)) == b {
				gold = gold + (int(a) * 3) + int(b)
			}
		}
	}

	return silver, gold
}

func main() {
	grid := GetLines("thirteen/input.txt")
	silver, gold := solve(grid)
	fmt.Printf("silver: %d\ngold: %d\n", silver, gold)
}
