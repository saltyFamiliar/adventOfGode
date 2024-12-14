package main

import (
	. "aoc_2024/utils"
	"fmt"
)

func calcAntinodeLocs(pointA, pointB Point) (Point, Point) {
	dx, dy := pointA.X-pointB.X, pointA.Y-pointB.Y
	return Point{Y: pointA.Y + dy, X: pointA.X + dx}, Point{Y: pointB.Y - dy, X: pointB.X - dx}
}

func pointPairIter(locMap map[string][]Point) func(func(Point, Point) bool) {
	return func(yield func(Point, Point) bool) {
		for _, pointList := range locMap {
			for i, ptA := range pointList[:len(pointList)-1] {
				for _, ptB := range pointList[i+1:] {
					if !yield(ptA, ptB) {
						return
					}
				}
			}
		}
	}
}

func main() {
	mat := GridFrom("input.txt")
	locMap := make(map[string][]Point)
	for pt, ch := range GridIter(mat) {
		if ch != "." {
			if _, ok := locMap[ch]; ok {
				locMap[ch] = append(locMap[ch], pt)
			} else {
				locMap[ch] = []Point{pt}
			}
		}
	}

	anLocs := make(map[Point]bool)
	for ptA, ptB := range pointPairIter(locMap) {
		anA, anB := calcAntinodeLocs(ptA, ptB)
		if PtInBounds(anA, mat) {
			anLocs[anA] = true
		}
		if PtInBounds(anB, mat) {
			anLocs[anB] = true
		}
	}
	silver := len(anLocs)

	for ptA, ptB := range pointPairIter(locMap) {
		dy, dx := ptA.Y-ptB.Y, ptA.X-ptB.X
		anX := Point{Y: ptA.Y, X: ptA.X}
		for ; PtInBounds(anX, mat); anX.X, anX.Y = anX.X+dx, anX.Y+dy {
			anLocs[anX] = true
		}
		anX = Point{Y: ptA.Y, X: ptA.X}
		for ; PtInBounds(anX, mat); anX.X, anX.Y = anX.X-dx, anX.Y-dy {
			anLocs[anX] = true
		}
	}

	for pt := range anLocs {
		mat[pt.Y][pt.X] = "#"
	}
	PrintGrid(mat)
	fmt.Printf("silver: %d\ngold: %d\n", silver, len(anLocs))
}
