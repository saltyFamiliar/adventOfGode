package main

import (
	. "aoc_2024/utils"
	"fmt"
	"maps"
	"slices"
	"strings"
)

func solve(grid [][]string) (silver, gold int) {
	type PointDir struct {
		pos Point
		dir Point
	}
	gridAt := func(pt Point) string {
		return grid[pt.Y][pt.X]
	}
	scoreGrid := make([][]map[Point]int, 0, len(grid))
	for y, row := range grid {
		scoreGrid = append(scoreGrid, []map[Point]int{})
		for range row {
			scoreGrid[y] = append(scoreGrid[y], make(map[Point]int))
		}
	}
	start := Point{Y: len(grid) - 2, X: 1}
	scoreGrid[start.Y][start.X][RIGHT] = 0

	q := []PointDir{{pos: start, dir: RIGHT}}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for _, dir := range DIRS_CLOCKWISE {
			ahead := n.pos.AddPt(dir)
			var aheadCost int
			if gridAt(ahead) != "#" {
				if dir == n.dir {
					aheadCost = scoreGrid[n.pos.Y][n.pos.X][n.dir] + 1
				} else {
					aheadCost = scoreGrid[n.pos.Y][n.pos.X][n.dir] + 1001
				}
				if _, ok := scoreGrid[ahead.Y][ahead.X][dir]; !ok {
					scoreGrid[ahead.Y][ahead.X][dir] = aheadCost
					q = append(q, PointDir{pos: ahead, dir: dir})
				} else if scoreGrid[ahead.Y][ahead.X][dir] > aheadCost {
					scoreGrid[ahead.Y][ahead.X][dir] = aheadCost
					q = append(q, PointDir{pos: ahead, dir: dir})
				}
			}
		}
	}
	costs := make([]int, 0, 4)
	for _, cost := range scoreGrid[1][len(scoreGrid[1])-2] {
		costs = append(costs, cost)
	}
	silver = slices.Min(costs)
	goodSeats := make(map[Point]bool)
	gs := []Point{{Y: 1, X: len(grid[1]) - 2}}
	for len(gs) > 0 {
		s := gs[0]
		gs = gs[1:]
		goodSeats[s] = true
		if grid[s.Y][s.X] == "S" {
			continue
		}
		grid[s.Y][s.X] = "O"
		minCost := slices.Min(slices.Collect(maps.Values(scoreGrid[s.Y][s.X])))
		for k, v := range scoreGrid[s.Y][s.X] {
			revDir := Point{Y: -k.Y, X: -k.X}
			if v == minCost {
				gs = append(gs, s.AddPt(revDir))
			} else {
				nxt := s.AddPt(k)
				if ok, _ := goodSeats[nxt]; !ok {
					continue
				}
				minNxt := slices.Min(slices.Collect(maps.Values(scoreGrid[nxt.Y][nxt.X])))
				var nxtDir Point
				for dir, cost := range scoreGrid[nxt.Y][nxt.X] {
					if cost == minNxt {
						nxtDir = dir
					}
				}
				if v == minCost+1000 && k == nxtDir {
					gs = append(gs, s.AddPt(revDir))
				}
			}
		}
	}
	gold = len(goodSeats)

	return silver, gold
}

func main() {
	lines := GetLines("sixteen/input.txt")
	var grid [][]string
	for _, row := range lines {
		grid = append(grid, strings.Split(row, ""))
	}
	silver, gold := solve(grid)
	fmt.Printf("silver: %d\ngold: %d\n", silver, gold)
}
