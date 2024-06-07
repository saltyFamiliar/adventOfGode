package testing

import (
	"adventOfGode/aocio"
	"fmt"
	"github.com/fatih/color"
)

func TestPart(partFunc func([]string) int, dirPath string, wanted int) {
	testPath := dirPath + "/test.txt"
	testResult := partFunc(aocio.FileToLines(testPath))
	if testResult == wanted {
		passStyle := color.New(color.FgGreen, color.Bold)
		passStyle.Println("Test case passed!")
		inputPath := dirPath + "/input.txt"
		fmt.Println("Input Result:", partFunc(aocio.FileToLines(inputPath)))
	} else {
		failStyle := color.New(color.FgRed, color.Bold)
		failStyle.Println("Test case failed!")
		fmt.Println("Test Result:", testResult)
	}
}
