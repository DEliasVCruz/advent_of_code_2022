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

func main() {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	var sides []string

	shapeValuesTheirs := map[string]int{"A": 1, "B": 2, "C": 3}
	winCondition := map[string]int{"A": 2, "B": 3, "C": 1}
	loseCondition := map[string]int{"A": 3, "B": 1, "C": 2}

	totalScore := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {

		sides = strings.Split(scanner.Text(), " ")

		switch sides[1] {
		case "Z":
			totalScore += 6
			totalScore += winCondition[sides[0]]
		case "X":
			totalScore += loseCondition[sides[0]]
		default:
			totalScore += 3
			totalScore += shapeValuesTheirs[sides[0]]
		}

	}
	check(scanner.Err())

	fmt.Println(totalScore)
}
