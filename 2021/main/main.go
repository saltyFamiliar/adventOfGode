package main

import (
	"adventOfGode/2021/seven"
	"adventOfGode/toolbelt"
)

func main() {
	dirPath := "2021/seven"
	toolbelt.TestPart(seven.PartOne, dirPath, 37)
	toolbelt.TestPart(seven.PartTwo, dirPath, -1)
}
