package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AISolve() {

	// Initialize variables
	numRed := 12
	numGreen := 13
	numBlue := 14

	var possibleGames []int
	var sum int

	// Open file
	file, _ := os.Open("solutions/day2/input.txt")

	// Read each game
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// Get game and rounds string
		game := scanner.Text()
		gameInfo := strings.Split(game, ":")
		roundsString := gameInfo[1]

		// Split rounds into array
		rounds := strings.Split(roundsString, ";")

		// Validate each round
		valid := true

		for _, round := range rounds {

			// Initialize cube counts for round
			cubeCounts := make(map[string]int)

			// Split round into cube draws
			cubeDraws := strings.Split(round, ",")

			// Process each cube draw
			for _, draw := range cubeDraws {
				draw = strings.TrimSpace(draw)
				cubeInfo := strings.Split(draw, " ")

				// Convert count to int
				count, err := strconv.Atoi(cubeInfo[0])
				if err != nil {
					// Handle error
				}

				// Get color
				color := cubeInfo[1]

				// Increment count for color
				cubeCounts[color] += count

			}

			// Validate cube counts
			if cubeCounts["red"] > numRed ||
				cubeCounts["green"] > numGreen ||
				cubeCounts["blue"] > numBlue {
				valid = false
				break
			}

		}

		// If valid, process game number
		if valid {
			gameNum := strings.Split(game, " ")[1]
			num, _ := strconv.Atoi(gameNum[:len(gameNum)-1])
			possibleGames = append(possibleGames, num)
			sum += num
		}

	}

	// Print output
	fmt.Println(sum)
	fmt.Println("done")

}
