package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func reverseArray(characters []string) {
	for i, j := 0, len(characters)-1; i < j; i, j = i+1, j-1 {
		characters[i], characters[j] = characters[j], characters[i]
	}
}

func main() {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	crateSectionCounter := 1
	stacks := make(map[int][]string)

	var instructions [3]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		inputLine := scanner.Text()

		if crateSectionCounter < 9 {
			for indx, char := range inputLine {
				trueIndex := indx - 1
				if ((trueIndex)%4) == 0 && string(char) != " " {
					stackNumber := trueIndex/4 + 1
					stacks[stackNumber] = append(stacks[stackNumber], string(char))
				}
			}
			crateSectionCounter++
			continue
		} else if crateSectionCounter < 11 {
			crateSectionCounter++
			continue
		} else if crateSectionCounter == 11 {
			for key := range stacks {
				reverseArray(stacks[key])
			}
			crateSectionCounter++
		}

		regex, _ := regexp.Compile(`\d+`)
		matches := regex.FindAllString(inputLine, -1)

		for indx, value := range matches {
			instructions[indx] = strToInt(value)
		}

		fromStackIndex := instructions[1]
		toStackIndex := instructions[2]

		removedIndex := len(stacks[fromStackIndex]) - instructions[0]
		crate := stacks[fromStackIndex][removedIndex:]

		stacks[fromStackIndex] = stacks[fromStackIndex][:removedIndex]
		for _, value := range crate {
			stacks[toStackIndex] = append(stacks[toStackIndex], value)

		}

	}
	check(scanner.Err())

	var stringSlice [9]string

	for i := 0; i < 9; i++ {
		stringSlice[i] = stacks[i+1][len(stacks[i+1])-1]
	}

	fmt.Println(strings.Join(stringSlice[:], ""))
}
