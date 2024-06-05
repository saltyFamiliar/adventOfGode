package main

import (
	"adventOfGode/2023/solutions/day9"
	"time"
)

func main() {
	startTime := time.Now()
	println("The answer is ", +day9.Solve2())
	print("Time taken: ", time.Since(startTime))
}
