package toolbelt

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
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
		passStyle := color.New(color.FgGreen, color.Bold)
		passStyle.Println("Test case passed!")
		fmt.Println("Input Result:", partFunc(FileToLines(inputPath)))
	} else {
		failStyle := color.New(color.FgRed, color.Bold)
		failStyle.Println("Test case failed!")
		fmt.Println("Test Result:", testResult)
	}
}

func StrArrToInts(ss []string) []int {
	ns := make([]int, 0, len(ss))
	for _, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err.Error())
		}
		ns = append(ns, n)
	}
	return ns
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

// Deprecated. Use generic All
func AllInStrArr(arr []string, val string) bool {
	for _, arrElement := range arr {
		if arrElement != val {
			return false
		}
	}
	return true
}

func All[E comparable](s []E, target E) bool {
	for _, e := range s {
		if e != target {
			return false
		}
	}
	return true
}

func StrArrToIntMap(arr []string) map[int]int {
	result := make(map[int]int)
	for _, digit := range arr {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			fmt.Println("Error converting digit to int")
		}
		result[digitInt] += 1
	}

	return result
}

func MapValSum[T comparable](_map map[T]int) (sum int) {
	for _, val := range _map {
		sum += val
	}

	return sum
}
