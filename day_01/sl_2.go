package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getMin(array [3]int32) int {
	min := array[0]
	index := 0

	for i, value := range array {
		if value < min {
			index = i
		}
	}

	return index
}

func sumArray(array [3]int32) int32 {
	var sum int32 = 0

	for _, value := range array {
		sum = sum + value
	}

	return sum
}

func main() {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	scann := bufio.NewScanner(input)

	top3 := [3]int32{0, 0, 0}
	var sum int32 = 0
	indexMin := 0

	for scann.Scan() {
		if scann.Text() != "" {
			currentNumber, err := strconv.Atoi(scann.Text())
			check(err)

			sum = sum + int32(currentNumber)
		} else {
			if sum > top3[indexMin] {
				top3[indexMin] = sum
				indexMin = getMin(top3)
			}

			sum = 0
		}
	}
	check(scann.Err())

	if sum > top3[indexMin] {
		top3[indexMin] = sum
	}

	total := sumArray(top3)

	fmt.Println(total)
}
