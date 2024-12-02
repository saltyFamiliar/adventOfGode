package four

import (
	"adventOfGode/toolbelt"
	"fmt"
	"strconv"
	"strings"
)

func markFromRow(row []string, toMark string) (newRow []string) {
	for _, val := range row {
		if val == toMark {
			newRow = append(newRow, "x")
		} else {
			newRow = append(newRow, val)
		}
	}
	return newRow
}

func boardWins(board [][]string) bool {
	for y, row := range board {
		// get column by using row number y as x value
		var column []string
		for x := 0; x < len(board); x += 1 {
			column = append(column, board[x][y])
		}
		//fmt.Println("AllinStrArr(", row, ", x): ", toolbelt.AllInStrArr(row, "x"), "len: ", len(row))
		if toolbelt.AllInStrArr(column, "x") || toolbelt.AllInStrArr(row, "x") {
			return true
		}
	}

	return false
}

func score(board [][]string) (score int) {
	for _, row := range board {
		for _, val := range row {
			numericVal, err := strconv.Atoi(val)
			if err != nil {
				score += 0
			} else {
				score += numericVal
			}
		}
	}
	return score
}

func turnsToWin(board [][]string, drawn []string, upperBound int) (turn int, rawScore int) {
	markedBoard := board
	for turn := 0; turn < upperBound; turn += 1 {
		var newBoard [][]string
		for _, row := range markedBoard {
			newBoard = append(newBoard, markFromRow(row, drawn[turn]))
		}
		markedBoard = newBoard

		if boardWins(markedBoard) {
			return turn + 1, score(markedBoard)
		}
	}

	return len(drawn) + 1, score(markedBoard)
}

func getBoards(boardLines []string) (boards [][][]string) {
	var newBoard [][]string
	for _, line := range boardLines {
		if len(line) < 2 {
			boards = append(boards, newBoard)
			newBoard = [][]string{}
		} else {
			lineNumbers := toolbelt.RemoveAllStrArr(strings.Split(line, " "), "")
			newBoard = append(newBoard, lineNumbers)
		}
	}
	boards = append(boards, newBoard)
	return boards
}

func PartOne(fileLines []string) (winningScore int) {
	drawn := strings.Split(fileLines[0], ",")
	boards := getBoards(fileLines[2:])

	fastestWin := len(drawn)
	for _, board := range boards {
		boardTurnsToWin, boardScore := turnsToWin(board, drawn, fastestWin)
		if boardTurnsToWin < fastestWin {
			fastestWin = boardTurnsToWin
			winningScore = boardScore
		}
	}

	lastDrawn, err := strconv.Atoi(drawn[fastestWin-1])
	if err != nil {
		fmt.Println("Error parsing last drawn")
	}

	return winningScore * lastDrawn
}

func PartTwo(fileLines []string) (losingScore int) {
	drawn := strings.Split(fileLines[0], ",")
	boards := getBoards(fileLines[2:])

	slowestWin := 0
	for _, board := range boards {
		boardTurnsToWin, boardScore := turnsToWin(board, drawn, len(drawn))
		if boardTurnsToWin > slowestWin {
			slowestWin = boardTurnsToWin
			losingScore = boardScore
		}
	}

	lastDrawn, err := strconv.Atoi(drawn[slowestWin-1])
	if err != nil {
		fmt.Println("Error parsing last drawn")
	}

	return losingScore * lastDrawn
}
