package main

import (
	"2023/solutions/day8"
	"time"
)

func main() {
	startTime := time.Now()
	println("The answer is ", +day8.Solve2())
	print("Time taken: ", time.Since(startTime))
}
