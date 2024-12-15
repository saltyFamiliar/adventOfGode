package main

import (
	. "aoc_2024/utils"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"slices"
	"time"
)

type robot struct {
	vy int
	vx int
	py int
	px int
	id int
}

func (r *robot) move(grid [][][]robot) {
	newLocY := (r.py + r.vy) % HEIGHT
	if newLocY < 0 {
		newLocY += HEIGHT
	}
	newLocX := (r.px + r.vx) % WIDTH
	if newLocX < 0 {
		newLocX += WIDTH
	}
	for i := 0; i < len(grid[r.py][r.px]); i++ {
		if grid[r.py][r.px][i].id == r.id {
			grid[r.py][r.px] = slices.Delete(grid[r.py][r.px], i, i+1)
			break
		}
	}
	r.px = newLocX
	r.py = newLocY
	grid[r.py][r.px] = append(grid[r.py][r.px], *r)
}

const (
	HEIGHT         = 103
	WIDTH          = 101
	SCALE          = 12
	RENDER_FPS     = 60
	SIMULATION_FPS = 60
	//HEIGHT = 7
	//WIDTH  = 11
	//HEIGHT = 800
	//WIDTH  = 800
)

func printGrid(grid [][][]robot) {
	for _, row := range grid {
		for _, robots := range row {
			if len(robots) == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d", len(robots))
			}
		}
		fmt.Printf("\n")
	}
}

func solve(lines []string) (silver, gold int) {
	var allRobots []*robot
	grid := make([][][]robot, HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		grid[i] = make([][]robot, WIDTH)
	}

	for _, l := range lines {
		var px, py, vx, vy int
		Must(fmt.Sscanf(l, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy))
		r := robot{vy, vx, py, px, len(allRobots)}
		allRobots = append(allRobots, &r)
		grid[py][px] = append(grid[py][px], r)
	}

	var simulationIter int
	for ; simulationIter < 7200; simulationIter++ {
		for i := range allRobots {
			allRobots[i].move(grid)
		}
	}

	rl.InitWindow(WIDTH*SCALE, HEIGHT*SCALE, "Robot Simulation")
	rl.SetTargetFPS(RENDER_FPS)
	defer rl.CloseWindow()
	simulationInterval := float32(1.0 / SIMULATION_FPS)
	var simulationTimer float32
	for !rl.WindowShouldClose() {
		simulationTimer += rl.GetFrameTime()
		for simulationTimer >= simulationInterval {
			for i := range allRobots {
				allRobots[i].move(grid)
			}
			simulationTimer -= simulationInterval
			simulationIter++
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for y, row := range grid {
			for x, robots := range row {
				if len(robots) == 1 {
					rl.DrawCircle(int32(x*SCALE+SCALE/2), int32(y*SCALE+SCALE/2), float32(SCALE/3), rl.White)
				} else if len(robots) == 2 {
					rl.DrawCircle(int32(x*SCALE+SCALE/2), int32(y*SCALE+SCALE/2), float32(SCALE/3), rl.Green)
				} else if len(robots) >= 3 {
					rl.DrawCircle(int32(x*SCALE+SCALE/2), int32(y*SCALE+SCALE/2), float32(SCALE/3), rl.Red)
				}
			}
		}
		rl.DrawText(fmt.Sprintf("T = %d", simulationIter), 0, 0, 10, rl.White)
		rl.EndDrawing()
		if simulationIter == 7383 {
			time.Sleep(10 * time.Second)
		}
	}

	var ul, ur, dl, dr int
	for _, r := range allRobots {
		if r.px < (WIDTH/2) && r.py < (HEIGHT/2) {
			ul++
		} else if r.px < (WIDTH/2) && r.py > (HEIGHT/2) {
			dl++
		} else if r.px > (WIDTH/2) && r.py < (HEIGHT/2) {
			ur++
		} else if r.px > (WIDTH/2) && r.py > (HEIGHT/2) {
			dr++
		}
	}

	return ul * ur * dl * dr, gold
}

func main() {
	lines := GetLines("fourteen/input.txt")
	silver, gold := solve(lines)
	fmt.Printf("silver: %d\ngold: %d\n", silver, gold)
}
