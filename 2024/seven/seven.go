package main

import (
	. "aoc_2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func validTestValue(testVal int, calibVals []int, numOps int) bool {
	maxVal := int(math.Pow(float64(numOps), float64(len(calibVals))))
	for n := 0; n < maxVal; n++ {
		total, opNum := calibVals[0], n
		for _, cv := range calibVals[1:] {
			switch opNum % numOps {
			case 0:
				total += cv
			case 1:
				total *= cv
			default:
				total, _ = strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(cv))
			}
			opNum /= numOps
		}
		if total == testVal {
			return true
		}
	}
	return false
}

func main() {
	lines := GetLines("input.txt")
	var wg sync.WaitGroup
	wg.Add(len(lines))
	var mu1, mu2 sync.Mutex
	var total1, total2 int
	for _, line := range lines {
		lineArr := strings.Split(line, ":")
		testVal, _ := strconv.Atoi(lineArr[0])
		calibVals := MapErr(strconv.Atoi, strings.Fields(lineArr[1]))
		go func() {
			defer wg.Done()
			if validTestValue(testVal, calibVals, 2) {
				mu1.Lock()
				total1 += testVal
				mu1.Unlock()
			}
			if validTestValue(testVal, calibVals, 3) {
				mu2.Lock()
				total2 += testVal
				mu2.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("silver: %d\ngold: %d\n", total1, total2)
}
