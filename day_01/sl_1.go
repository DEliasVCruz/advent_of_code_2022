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

func new_max(max int32, number int32) int32 {

	if max < number {
		return number
	}

	return max
}

func main() {
	input, err := os.Open("./input.txt")

	check(err)

	defer input.Close()

	scann := bufio.NewScanner(input)

	var max int32 = 0
	var sum int32 = 0

	for scann.Scan() {
		if scann.Text() != "" {
			current_number, err := strconv.Atoi(scann.Text())
			check(err)
			sum = sum + int32(current_number)
		} else {
			max = new_max(max, sum)
			sum = 0
		}
	}
	check(scann.Err())

	max = new_max(max, sum)
	fmt.Println(max)

}
