package main

import (
	"adventOfGode/2021/three"
	"adventOfGode/toolbelt"
)

func main() {
	dirPath := "2021/two"
	toolbelt.TestPart(three.PartOne, dirPath, 198)
	toolbelt.TestPart(three.PartTwo, dirPath, 230)
}
