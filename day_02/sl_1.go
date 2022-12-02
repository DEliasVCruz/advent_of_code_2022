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
	var sideTheirs int
	var sideOurs int

	shapeValuesOurs := map[string]int{"X": 1, "Y": 2, "Z": 3}
	shapeValuesTheirs := map[string]int{"A": 1, "B": 2, "C": 3}
	winCondition := map[int]int{3: 2, 4: 1, 5: 3}

	totalScore := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {

		sides = strings.Split(scanner.Text(), " ")
		sideTheirs = shapeValuesTheirs[sides[0]]
		sideOurs = shapeValuesOurs[sides[1]]

		totalScore += sideOurs

		if sideTheirs == sideOurs {
			totalScore += 3
		} else if sideOurs == winCondition[sideTheirs+sideOurs] {
			totalScore += 6
		}
	}
	check(scanner.Err())

	fmt.Println(totalScore)
}
