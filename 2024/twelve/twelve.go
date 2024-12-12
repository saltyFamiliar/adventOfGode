package main

import (
	. "aoc_2024/utils"
	"cmp"
	"fmt"
	"maps"
	"slices"
	"sort"
)

func solve(grid Grid) (silver int, gold int) {
	var fences [][][]Point

	surveyRegion := func(pt Point) (cost int, costBulk int, plotsVisited map[Point]bool) {
		plotsVisited = make(map[Point]bool)
		q := []Point{pt}
		var perimeter, area, sides int
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if plotsVisited[p] {
				continue
			}
			area++
			plotsVisited[p] = true
			pFences := fences[p.Y][p.X]
			perimeter += len(pFences)
			for _, d := range DIRS_CLOCKWISE {
				if !slices.Contains(pFences, d) && !plotsVisited[p.AddPt(d)] {
					q = append(q, p.AddPt(d))
				}
			}
		}
		regionPlots := slices.Collect(maps.Keys(plotsVisited))
		regionPlots = slices.DeleteFunc(regionPlots, func(pt Point) bool {
			return len(fences[pt.Y][pt.X]) == 0
		})

		slices.SortFunc(regionPlots, func(a, b Point) int {
			if a.Y > b.Y {
				return 1
			} else if a.Y < b.Y {
				return -1
			}
			return cmp.Compare(a.X, b.X)
		})
		fenceLocsY := make(map[float64][]int)
		fenceLocsX := make(map[float64][]int)
		for _, pt := range regionPlots {
			for _, ptFence := range fences[pt.Y][pt.X] {
				if ptFence.Y != 0 {
					fLoc := (float64(ptFence.Y) / 2.0) + float64(pt.Y)
					if _, ok := fenceLocsY[fLoc]; ok {
						fenceLocsY[fLoc] = append(fenceLocsY[fLoc], pt.X)
					} else {
						fenceLocsY[fLoc] = []int{pt.X}
					}
				}
				if ptFence.X != 0 {
					fLoc := (float64(ptFence.X) / 2.0) + float64(pt.X)
					if _, ok := fenceLocsX[fLoc]; ok {
						fenceLocsX[fLoc] = append(fenceLocsX[fLoc], pt.Y)
					} else {
						fenceLocsX[fLoc] = []int{pt.Y}
					}
				}
			}
		}

		for y, _ := range fenceLocsY {
			sort.Ints(fenceLocsY[y])
			startX := -2
			for _, x := range fenceLocsY[y] {
				if x != startX+1 && startX != x {
					sides++
				}
				startX = x
			}
		}
		for x, _ := range fenceLocsX {
			sort.Ints(fenceLocsX[x])
			startY := -2
			for _, y := range fenceLocsX[x] {
				if y != startY+1 && startY != y {
					sides++
				}
				startY = y
			}
		}
		fmt.Println(area, sides)
		return perimeter * area, sides * area, plotsVisited
	}

	for y, row := range grid {
		fences = append(fences, [][]Point{})
		for x, ch := range row {
			fences[y] = append(fences[y], []Point{})
			for _, d := range DIRS_CLOCKWISE {
				n := Point{Y: y, X: x}.AddPt(d)
				if !n.In(grid) || grid.At(n) != ch {
					fences[y][x] = append(fences[y][x], d)
				}
			}
		}
	}

	plotsVisited := make(map[Point]bool)
	for pt := range GridIter(grid) {
		if !plotsVisited[pt] {
			cost, bulkCost, newPlotsVisited := surveyRegion(pt)
			silver += cost
			gold += bulkCost
			for k := range newPlotsVisited {
				plotsVisited[k] = true
			}
		}
	}

	return
}

func main() {
	grid := GridFrom("twelve/input.txt")
	silver, gold := solve(grid)
	fmt.Printf("silver: %d\ngold: %d\n", silver, gold)
}
