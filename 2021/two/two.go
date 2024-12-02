package two

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(fileLines []string) int {
	horizontal := 0
	depth := 0

	for _, line := range fileLines {
		unit, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			fmt.Println("Error parsing units")
		}

		switch line[0] {
		case 'f':
			horizontal += unit
		case 'u':
			depth -= unit
		case 'd':
			depth += unit
		default:
			fmt.Println("Error parsing direction")
			return 0
		}
	}

	return horizontal * depth
}

func PartTwo(fileLines []string) (result int) {
	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range fileLines {
		unit, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			fmt.Println("Error parsing units")
		}

		switch line[0] {
		case 'f':
			horizontal += unit
			depth += unit * aim
		case 'u':
			aim -= unit
		case 'd':
			aim += unit
		default:
			fmt.Println("Error parsing direction")
			return 0
		}
	}

	return horizontal * depth
}
