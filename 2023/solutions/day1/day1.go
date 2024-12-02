package day1

import (
	"2023/ergo"
	"bufio"
	"fmt"
	"os"
)

func Solve1() (totalValue int) {
	f, err := os.Open("solutions/day1/input1.txt")
	ergo.Must("open file", err)

	r := bufio.NewScanner(f)
	for r.Scan() {
		byteBuff := r.Bytes()
		var lineValue int

		for _, b := range byteBuff {
			d, err := ergo.RuneToIntIfDigit(rune(b))
			if err != nil {
				continue
			}

			lineValue = d * 10
			break
		}

		for i := len(byteBuff) - 1; i >= 0; i-- {
			d, err := ergo.RuneToIntIfDigit(rune(byteBuff[i]))
			if err != nil {
				continue
			}

			lineValue += d
			break
		}

		totalValue += lineValue
		fmt.Println(r.Text())
		println(lineValue)
	}

	return totalValue
}

func Solve2() (totalValue int) {
	numberNames := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	shortLen, longLen := 3, 5

	f, err := os.Open("solutions/day1/input1.txt")
	ergo.Must("open file", err)

	r := bufio.NewScanner(f)
	for r.Scan() {
		lineStr := r.Text()
		fmt.Println(lineStr)

		var lineValue int
		byteBuff := r.Bytes()
	firstNum:
		for i, b := range byteBuff {
			d, err := ergo.RuneToIntIfDigit(rune(b))
			if err == nil {
				lineValue = d * 10
				break
			}

			for j := i + shortLen; j <= i+longLen && j < len(byteBuff); j++ {
				possibleName := string(byteBuff[i:j])
				for k, name := range numberNames {
					if possibleName == name {
						println(possibleName)
						lineValue = (k + 1) * 10
						break firstNum
					}
				}
			}

		}

	secondNum:
		for i := len(byteBuff) - 1; i >= 0; i-- {
			d, err := ergo.RuneToIntIfDigit(rune(byteBuff[i]))
			if err == nil {
				lineValue += d
				break
			}

			for j := i + shortLen; j <= i+longLen && j <= len(byteBuff); j++ {
				possibleName := string(byteBuff[i:j])
				for k, name := range numberNames {
					if possibleName == name {
						println(possibleName)
						lineValue += k + 1
						break secondNum
					}
				}
			}

		}

		totalValue += lineValue
		println(lineValue)
	}

	return totalValue
}
