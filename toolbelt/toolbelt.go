package toolbelt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FileToLines(filePath string) (lines []string) {
	openFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}

	fileScanner := bufio.NewScanner(openFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	err = openFile.Close()
	if err != nil {
		log.Println(err)
	}

	return lines
}

func TestPart(partFunc func([]string) int, dirPath string, wanted int) {
	testPath := dirPath + "/test.txt"
	inputPath := dirPath + "/input.txt"
	testResult := partFunc(FileToLines(testPath))
	if testResult == wanted {
		fmt.Println("Test case passed!")
		fmt.Println("Result: ", partFunc(FileToLines(inputPath)))
	} else {
		fmt.Println("Test case failed!\nResult:", testResult)
	}
}

func StrArrToInts(strs []string) (ints []int) {
	for _, str := range strs {
		newInt, err := strconv.Atoi(str)
		if err != nil {
			log.Println("Found non-numeric value in string array")
			return ints
		}
		ints = append(ints, newInt)
	}
	return ints
}

func RemoveFirst[T comparable](arr []T, toRemove T) []T {
	for i, el := range arr {
		if el == toRemove {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func RemoveAllStrArr(arr []string, toRemove string) (newArr []string) {
	for _, el := range arr {
		if el != toRemove {
			newArr = append(newArr, el)
		}
	}
	return newArr
}

func AllInStrArr(arr []string, val string) bool {
	for _, arrElement := range arr {
		if arrElement != val {
			return false
		}
	}
	return true
}
