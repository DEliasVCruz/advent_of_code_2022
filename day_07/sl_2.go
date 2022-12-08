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

func strToInt(value string) int64 {
	// convertedStr, err := strconv.Atoi(value)
	convertedStr, err := strconv.ParseInt(value, 10, 64)
	check(err)
	return convertedStr
}

func reverseArray(characters []string) {
	for i, j := 0, len(characters)-1; i < j; i, j = i+1, j-1 {
		characters[i], characters[j] = characters[j], characters[i]
	}
}

func sumSubDirSizes(dir string, dirMap map[string][]string, sizeMap map[string]int) int {
	subDirsTotalSize := 0

	if _, ok := dirMap[dir]; !ok {
		return sizeMap[dir]
	} else {
		subDirsTotalSize += sizeMap[dir]
		for _, dir := range dirMap[dir] {
			subDirsTotalSize += sumSubDirSizes(dir, dirMap, sizeMap)
		}
	}

	return subDirsTotalSize
}

func main() {
	input, err := os.Open("./input.txt")
	check(err)
	defer input.Close()

	var currentPath []string
	var currentDir string
	var directoryFiles = map[string]map[string]string{}
	var absolutePaths = map[string][]string{}

	cdCommand, _ := regexp.Compile(`^\$ cd (.*)`)
	fileType, _ := regexp.Compile(`^(\d+) (.*)`)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		inputLine := scanner.Text()

		switch {
		case cdCommand.MatchString(inputLine):

			dirArg := cdCommand.FindStringSubmatch(inputLine)[1]
			if dirArg == ".." {
				currentDir = strings.Join(currentPath[:len(currentPath)-1], "/")
				currentPath = currentPath[:len(currentPath)-1]
			} else {
				currentPath = append(currentPath, dirArg)
				currentDir = strings.Join(currentPath, "/")
				absolutePaths[currentDir] = make([]string, len(currentPath)-1)
				copy(absolutePaths[currentDir], currentPath[:len(currentPath)-1])
			}

		case fileType.MatchString(inputLine):

			fileInfo := fileType.FindStringSubmatch(inputLine)[1:]
			if _, ok := directoryFiles[currentDir]; !ok {
				directoryFiles[currentDir] = make(map[string]string)
			}
			directoryFiles[currentDir][fileInfo[1]] = fileInfo[0]
		}

	}
	check(scanner.Err())

	directorySizes := make(map[string]int64)

	for directory, files := range directoryFiles {
		var dirSize int64 = 0
		for _, fileSize := range files {
			dirSize += strToInt(fileSize)
		}
		directorySizes[directory] = dirSize
	}

	var dirsTotalSize = map[string]int64{}

	for dir, path := range absolutePaths {
		if _, ok := directorySizes[dir]; ok {
			for i := 1; i <= len(path); i++ {
				dirPath := strings.Join(path[:i], "/")
				dirsTotalSize[dirPath] += directorySizes[dir]
			}
		}
	}

	for dir, size := range directorySizes {
		dirsTotalSize[dir] += size
	}

	unusedSpace := 70000000 - dirsTotalSize["/"]
	neededSpace := 30000000 - unusedSpace

	minSpaceDir := dirsTotalSize["/"]

	for _, size := range dirsTotalSize {
		if size < minSpaceDir && size >= neededSpace {
			minSpaceDir = size
		}
	}

	fmt.Println(minSpaceDir)
}
