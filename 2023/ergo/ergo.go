package ergo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ":" + err.Error())
	}
}

func RuneToIntIfDigit(r rune) (int, error) {
	if !unicode.IsDigit(r) {
		return 0, fmt.Errorf("Not a digit")
	}

	return int(r - '0'), nil
}

func GetInputScanner(filePath string) *bufio.Scanner {
	file, err := os.Open(filePath)
	Must("open file", err)
	scanner := bufio.NewScanner(file)

	return scanner
}

func EzIntParse(digits string) int {
	num, err := strconv.Atoi(digits)
	Must("parse number", err)
	return num
}
