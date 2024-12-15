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
	searchPt := r.pt.AddPt(moveDir)
	for ; gridAt(searchPt) != "."; searchPt = searchPt.AddPt(moveDir) {
		if gridAt(searchPt) == "#" {
			return false
		}
	}
	invDir := Point{Y: -moveDir.Y, X: -moveDir.X}
	spaces := Abs(searchPt.Y-r.pt.Y) + Abs(searchPt.X-r.pt.X)
	for i := 0; i < spaces; i++ {
		grid[searchPt.Y][searchPt.X],
			grid[searchPt.Y+invDir.Y][searchPt.X+invDir.X] =
			grid[searchPt.Y+invDir.Y][searchPt.X+invDir.X],
			grid[searchPt.Y][searchPt.X]
		searchPt = searchPt.AddPt(invDir)
	}
	r.pt = r.pt.AddPt(moveDir)
	return true
}

const (
	SCALE     = 12
	FPS       = 2
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
				rl.DrawCircle(centerX, centerY, unitSizeF, rl.Brown)
			} else if x == int32(r.pt.X) && y == int32(r.pt.Y) {
				rl.DrawCircle(centerX, centerY, unitSizeF, rl.Red)
			} else if ch == "[" {
				rl.DrawRectangle(xSc+SCALE/3, centerY, unitSize*2, unitSize, rl.Brown)
			} else if ch == "]" {
				rl.DrawRectangle(xSc, centerY, unitSize*2, unitSize, rl.Brown)
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
	for y, row := range grid {
		for x, ch := range row {
			if ch == "O" {
				silver += (100 * y) + x
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
