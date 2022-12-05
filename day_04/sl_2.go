package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type void struct{}

var member void

func getSet(array []string) map[string]void {
	set := make(map[string]void)
	for _, value := range array {
		set[value] = member
	}

	return set
}

func strToInt(value string) int {
	convertedStr, err := strconv.Atoi(value)
	check(err)
	return convertedStr
}

func main() {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	var fullyContainedRangeCounter = 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rangePairs := strings.Split(scanner.Text(), ",")
		firstPair := rangePairs[0]
		secondPair := rangePairs[1]

		firstPairRange := strings.Split(firstPair, "-")
		secondPairRange := strings.Split(secondPair, "-")

		firstPairRangeLower := strToInt(firstPairRange[0])
		firstPairRangeUpper := strToInt(firstPairRange[1])
		secondPairRangeLower := strToInt(secondPairRange[0])
		secondPairRangeUpper := strToInt(secondPairRange[1])

		if secondPairRangeLower >= firstPairRangeLower {
			if secondPairRangeLower <= firstPairRangeUpper {
				fullyContainedRangeCounter += 1
			}
		} else if firstPairRangeLower <= secondPairRangeUpper {
			fullyContainedRangeCounter += 1
		}

		// if (secondPairRangeLower <= firstPairRangeUpper) && (secondPairRangeLower >= firstPairRangeLower) {
		// 	fullyContainedRangeCounter += 1
		// } else if (secondPairRangeUpper <= firstPairRangeUpper) && (secondPairRangeUpper >= firstPairRangeLower) {
		// 	fullyContainedRangeCounter += 1
		// }

	}
	check(scanner.Err())

	fmt.Println(fullyContainedRangeCounter)
}
