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

	index := 0
	prioritySum := 0
	var processedLInes [3]string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		processedLInes[index] = scanner.Text()

		if index == 2 {
			repetitionCounterArray := make(map[string]int)
			for _, line := range processedLInes {
				arrayLine := strings.Split(line, "")
				setLine := getSet(arrayLine)
				for value := range setLine {
					repetitionCounterArray[value] += 1
				}
			}

			for key, value := range repetitionCounterArray {
				if value == 3 {
					prioritySum += getPriority(key)
				}
			}

			index = 0
		} else {
			index += 1
		}

	}
	check(scanner.Err())

	fmt.Println(prioritySum)
}
