package main

import (
	"adventOfGode/2021/four"
	"adventOfGode/toolbelt"
)

func main() {
	dirPath := "2021/four"
	toolbelt.TestPart(four.PartOne, dirPath, 4512)
	toolbelt.TestPart(four.PartTwo, dirPath, 1924)
}
