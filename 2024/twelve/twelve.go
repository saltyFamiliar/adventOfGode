package main

import (
	. "aoc_2024/utils"
	"cmp"
	"fmt"
	"maps"
	"slices"
)

func solve(grid Grid) (silver int, gold int) {
	var fences [][][]Point
	surveyRegion := func(pt Point) (cost int, costBulk int, plotsVisited map[Point]bool) {
		plotsVisited, q, p := make(map[Point]bool), []Point{pt}, pt
		var perimeter, area, sides int
		for len(q) > 0 {
			p, q = q[0], q[1:]
			if plotsVisited[p] {
				continue
			}
			area, plotsVisited[p] = area+1, true
			perimeter += len(fences[p.Y][p.X])
			for _, d := range DIRS_CLOCKWISE {
				if !slices.Contains(fences[p.Y][p.X], d) && !plotsVisited[p.AddPt(d)] {
					q = append(q, p.AddPt(d))
				}
			}
		}

		regionPlots := slices.Collect(maps.Keys(plotsVisited))
		regionPlots = slices.DeleteFunc(regionPlots, func(pt Point) bool {
			return len(fences[pt.Y][pt.X]) == 0
		})
		fenceLocsY, fenceLocsX := make(map[float64][]Point), make(map[float64][]Point)
		for _, pt := range regionPlots {
			for _, ptFence := range fences[pt.Y][pt.X] {
				if ptFence.Y != 0 {
					fLoc := (float64(ptFence.Y) / 2.0) + float64(pt.Y)
					if _, ok := fenceLocsY[fLoc]; ok {
						fenceLocsY[fLoc] = append(fenceLocsY[fLoc], pt) //pt.X
					} else {
						fenceLocsY[fLoc] = []Point{pt} //pt.X
					}
				}
				if ptFence.X != 0 {
					fLoc := (float64(ptFence.X) / 2.0) + float64(pt.X)
					if _, ok := fenceLocsX[fLoc]; ok {
						fenceLocsX[fLoc] = append(fenceLocsX[fLoc], pt) //pt.Y
					} else {
						fenceLocsX[fLoc] = []Point{pt} //pt.Y
					}
				}
			}
		}

		for y, _ := range fenceLocsY {
			slices.SortFunc(fenceLocsY[y], func(a, b Point) int {
				return cmp.Compare(a.X, b.X)
			})
			prevPoint := Point{Y: -2, X: -2}
			for _, pt := range fenceLocsY[y] {
				if pt.X != prevPoint.X+1 || pt.Y != prevPoint.Y {
					sides++
				}
				prevPoint = pt
			}
		}
		for x, _ := range fenceLocsX {
			slices.SortFunc(fenceLocsX[x], func(a, b Point) int {
				return cmp.Compare(a.Y, b.Y)
			})
			prevPt := Point{Y: -2, X: -2}
			for _, pt := range fenceLocsX[x] {
				if pt.Y != prevPt.Y+1 || pt.X != prevPt.X {
					sides++
				}
				prevPt = pt
			}
		}
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
			silver, gold = silver+cost, gold+bulkCost
			for k := range newPlotsVisited {
				plotsVisited[k] = true
			}
		}
	}
	return silver, gold
}

func main() {
	grid := GridFrom("twelve/input.txt")
	silver, gold := solve(grid)
	fmt.Printf("silver: %d\ngold: %d\n", silver, gold)
}
