package main

import (
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

func reverseArray(characters []string) {
	for i, j := 0, len(characters)-1; i < j; i, j = i+1, j-1 {
		characters[i], characters[j] = characters[j], characters[i]
	}
}

func main() {
	data, err := os.ReadFile("./input.txt")
	check(err)

	bufferSlice := strings.Split(string(data), "")

	for index := range bufferSlice {
		paquet := bufferSlice[index : index+14]
		paquetSet := getSet(paquet)
		if len(paquetSet) == 14 {
			fmt.Println(index + 14)
			break
		}
	}

}
