package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

var anchorValue = int("a"[0])

func getPriority(value string) int {
	anchorDifference := int(value[0]) - anchorValue
	if anchorDifference < 0 {
		normalizeDifference := int(strings.ToLower(value)[0]) - anchorValue
		return normalizeDifference + 1 + 26
	} else {
		return anchorDifference + 1
	}
}

func main() {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	prioritySum := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		inputText := scanner.Text()
		inputMedian := len(inputText) / 2
		stringArray := strings.Split(inputText, "")

		firstSack := stringArray[:inputMedian]
		secondStack := stringArray[inputMedian:]

		firstSetSack := getSet(firstSack)
		secondSetSack := getSet(secondStack)

		repetitionCounterArray := make(map[string]int)

		for value := range firstSetSack {
			repetitionCounterArray[value] += 1
		}

		for value := range secondSetSack {
			repetitionCounterArray[value] += 1
			if repetitionCounterArray[value] > 1 {
				prioritySum += getPriority(value)
				break
			}
		}

	}
	check(scanner.Err())

	fmt.Println(prioritySum)
}
