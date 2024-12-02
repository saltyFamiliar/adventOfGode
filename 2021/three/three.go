package three

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func digitPolarity(d rune) int {
	if d == '0' {
		return -1
	} else {
		return 1
	}
}

func onlyCommonStrs(binStrs []string, digitPosition int, defaultDigit rune, inverse bool) (common []string) {
	if len(binStrs) < 1 {
		log.Println("Got empty list of bin strings")
		return
	} else if digitPosition >= len(binStrs[0]) {
		log.Println("Got invalid digit position for bin string")
		return
	}

	// get most common digit at given position of each binary string
	polarity := 0
	for _, binStr := range binStrs {
		polarity += digitPolarity(rune(binStr[digitPosition]))
	}
	commonDigit := defaultDigit
	if polarity > 0 {
		commonDigit = '1'
		if inverse {
			commonDigit = '0'
		}
	} else if polarity < 0 {
		commonDigit = '0'
		if inverse {
			commonDigit = '1'
		}
	}

	// append all binary strings that have most common digit at position to common
	for _, binStr := range binStrs {
		if rune(binStr[digitPosition]) == commonDigit {
			common = append(common, binStr)
		}
	}

	return common
}

func PartOne(fileLines []string) (result int) {
	strLen := len(fileLines[0])

	digitPolarities := make([]int, strLen)
	for _, bitString := range fileLines {
		for i, digit := range bitString {
			digitPolarities[i] += digitPolarity(digit)
		}
	}

	gamma := 0
	epsilon := 0
	for i, digit := range digitPolarities {
		if digit > 0 {
			gamma += int(math.Pow(2, float64((strLen-1)-i)))
		} else {
			epsilon += int(math.Pow(2, float64((strLen-1)-i)))
		}
	}
	return gamma * epsilon
}

func PartTwo(fileLines []string) (result int) {
	strLen := len(fileLines[0])

	// narrow down binary strings to oxygen rating
	oxygenRatings := fileLines
	for i := 0; len(oxygenRatings) > 1 && i < strLen; i += 1 {
		if i == strLen {
			fmt.Println("Multiple possible oxygen ratings found")
			break
		}
		oxygenRatings = onlyCommonStrs(oxygenRatings, i, '1', false)
	}
	oxygenRatingBin := oxygenRatings[0]

	// narrow down binary strings to cO2 rating
	cO2Ratings := fileLines
	for i := 0; len(cO2Ratings) > 1 && i < strLen; i += 1 {
		if i == strLen {
			fmt.Println("Multiple possible CO2 ratings found")
			break
		}
		cO2Ratings = onlyCommonStrs(cO2Ratings, i, '0', true)
	}
	cO2RatingBin := cO2Ratings[0]

	cO2Rating := 0
	oxygenRating := 0
	for i := 0; i < strLen; i += 1 {
		cO2Digit, _ := strconv.Atoi(string(cO2RatingBin[i]))
		cO2Rating += int(math.Pow(2, float64((strLen-1)-i))) * cO2Digit

		oxygenDigit, _ := strconv.Atoi(string(oxygenRatingBin[i]))
		oxygenRating += int(math.Pow(2, float64((strLen-1)-i))) * oxygenDigit
	}

	return cO2Rating * oxygenRating
}
