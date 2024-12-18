package main

import (
	. "aoc_2024/utils"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strings"
)

func parse(lines []string, doubleWidth bool) (grid [][]string, moves []string, r robot) {
	var i int
	for ; i < len(lines) && len(lines[i]) > 2; i++ {
		if !doubleWidth {
			grid = append(grid, strings.Split(lines[i], ""))
		} else {
			var newLine []string
			for x, ch := range strings.Split(lines[i], "") {
				if ch == "#" {
					newLine = append(newLine, "#", "#")
				} else if ch == "." {
					newLine = append(newLine, ".", ".")
				} else if ch == "O" {
					newLine = append(newLine, "[", "]")
				} else if ch == "@" {
					newLine = append(newLine, ".", ".")
					r = robot{pt: Point{Y: i, X: x * 2}}
				}
			}
			grid = append(grid, newLine)
		}
	}
	for _, ml := range lines[i+1:] {
		moves = append(moves, strings.Split(ml, "")...)
	}
	return grid, moves, r
}

type robot struct {
	pt Point
}

func (r *robot) act(cmd string, grid [][]string) bool {
	var moveDir Point
	switch cmd {
	case "^":
		moveDir = Point{Y: -1}
	case ">":
		moveDir = Point{X: 1}
	case "v":
		moveDir = Point{Y: 1}
	case "<":
		moveDir = Point{X: -1}
	}
	gridAt := func(pt Point) string {
		return grid[pt.Y][pt.X]
	}
	safeAdd := func(pts []Point, pt Point) []Point {
		for _, p := range pts {
			if p == pt {
				return pts
			}
		}
		return append(pts, pt)
	}
	searchPts := []Point{r.pt.AddPt(moveDir)}
	//movePts := []Point{r.pt}
	var movePts []Point
	var sp Point
	for len(searchPts) > 0 {
		sp, searchPts = searchPts[0], searchPts[1:]
		if gridAt(sp) == "#" {
			return false
		}
		if gridAt(sp) == "." {
			continue
		}
		newSp := sp.AddPt(moveDir)
		searchPts = safeAdd(searchPts, newSp)
		movePts = safeAdd(movePts, sp)
		if gridAt(sp) == "O" {
			continue
		}
		if gridAt(sp) == "[" && moveDir != LEFT {
			searchPts = safeAdd(searchPts, newSp.AddPt(RIGHT))
			movePts = safeAdd(movePts, sp.AddPt(RIGHT))
		} else if gridAt(sp) == "]" && moveDir != RIGHT {
			searchPts = safeAdd(searchPts, newSp.AddPt(LEFT))
			movePts = safeAdd(movePts, sp.AddPt(LEFT))
		}
	}
	for len(movePts) > 0 {
		mp := movePts[len(movePts)-1]
		movePts = movePts[:len(movePts)-1]
		nxt := mp.AddPt(moveDir)
		grid[mp.Y][mp.X], grid[nxt.Y][nxt.X] = grid[nxt.Y][nxt.X], grid[mp.Y][mp.X]
	}
	r.pt = r.pt.AddPt(moveDir)
	return true
}

const (
	SCALE     = 12
	FPS       = 1200
	unitSize  = int32(SCALE / 3)
	unitSizeF = float32(unitSize)
)

func render(grid [][]string, r robot) {
	rl.ClearBackground(rl.Black)
	rl.BeginDrawing()
	for y, row := range grid {
		for x, ch := range row {
			y, x := int32(y), int32(x)
			ySc, xSc := y*SCALE, x*SCALE
			centerY, centerX := ySc+SCALE/2, xSc+SCALE/2
			if ch == "#" {
				rl.DrawCircle(centerX, centerY, unitSizeF, rl.White)
			} else if ch == "O" {
				rl.DrawCircle(centerX, centerY, unitSizeF, rl.Green)
			} else if x == int32(r.pt.X) && y == int32(r.pt.Y) {
				rl.DrawCircle(centerX, centerY, unitSizeF, rl.Red)
			} else if ch == "[" {
				rl.DrawRectangle(xSc+SCALE/3, centerY, unitSize*2, unitSize, rl.Green)
			} else if ch == "]" {
				rl.DrawRectangle(xSc, centerY, unitSize*2, unitSize, rl.Green)
			}
		}
	}
	rl.EndDrawing()
}

func solve(grid [][]string, moves []string, r robot) (silver, gold int) {
	for i := 0; !rl.WindowShouldClose() && i < len(moves); i++ {
		render(grid, r)
		r.act(moves[i], grid)
	}
	render(grid, r)
	//time.Sleep(100 * time.Second)
	for y, row := range grid {
		for x, ch := range row {
			if ch == "O" || ch == "[" || ch == "]" {
				silver += (100 * y) + x
			}
		}
	}
	for y, row := range grid {
		for x, ch := range row {
			if ch == "[" {
				gold += (y * 100) + x
			}
		}
	}
	return silver, gold
}

func main() {
	lines := GetLines("fifteen/input.txt")
	grid, moves, r := parse(lines, true)
	var HEIGHT, WIDTH = int32(len(grid)), int32(len(grid[0]))
	rl.InitWindow(WIDTH*SCALE, HEIGHT*SCALE, "Day 15")
	defer rl.CloseWindow()
	rl.SetTargetFPS(FPS)
	silver, gold := solve(grid, moves, r)
	fmt.Printf("silver: %d\ngold: %d\n", silver, gold)
}
