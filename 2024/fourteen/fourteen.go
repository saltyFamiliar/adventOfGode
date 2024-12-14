package main

import (
	. "aoc_2024/utils"
	"fmt"
	"os"
	"os/exec"
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
	HEIGHT = 103
	WIDTH  = 101
	//HEIGHT = 7
	//WIDTH  = 11
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
	for t := 1; t < 10000; t++ {
		for _, r := range allRobots {
			r.move(grid)
		}
		overlap := false
		for _, row := range grid {
			for _, robots := range row {
				if len(robots) > 1 {
					overlap = true
					break
				}
			}
			if overlap {
				break
			}
		}
		if overlap {
			continue
		}
		//if t < 70 {
		//	continue
		//}
		printGrid(grid)
		if t == 70 {
			break
		}
		fmt.Println(t)
		time.Sleep(100 * time.Second)
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}

	//printGrid(grid)
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
