package aocio

import (
	"adventOfGode/toolbelt"
	"bufio"
	"os"
)

func FileToLines(filePath string) (lines []string) {
	openFile, err := os.Open(filePath)
	toolbelt.Must("open file", err)
	fileScanner := bufio.NewScanner(openFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			lines = append(lines, fileScanner.Text())
		}
	}
	toolbelt.Must("close file", openFile.Close())
	return lines
}

func FileToByteMatrix(filePath string) (matrix [][]byte) {
	openFile, err := os.Open(filePath)
	toolbelt.Must("open file", err)
	defer func() {
		toolbelt.Must("close file", openFile.Close())
	}()

	fileScanner := bufio.NewScanner(openFile)
	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		byteRow := make([]byte, len(line))
		copy(byteRow, line)
		matrix = append(matrix, byteRow)
	}

	return matrix
}
