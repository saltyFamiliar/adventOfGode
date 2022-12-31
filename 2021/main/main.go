package main

import (
	"adventOfGode/2021/two"
	"adventOfGode/toolbelt"
)

func main() {
	dirPath := "2021/two"
	toolbelt.TestPart(two.PartOne, dirPath, 150)
	toolbelt.TestPart(two.PartTwo, dirPath, 900)
}
