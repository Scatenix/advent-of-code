package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Left = 0
const Right = 1

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileDistancesPath := getDistancesFilePathFromArgs()
	fileExists(fileDistancesPath)
	file, err := os.Open(fileDistancesPath)
	check(err)

	leftDistances, rightDistances, err := ReadDistances(file)
	check(err)

	similarityScore := 0;
	for _, vi := range leftDistances {
		count :=0
		for _, vj := range rightDistances {
			if vi == vj {
				count++
			}
		}
		similarityScore += vi * count;
	}

	fmt.Printf("The similarity score is: %d\n", similarityScore)
}

func ReadDistances(file *os.File) ([]int, []int, error) {
	var leftDistances []int
	var rightDistances []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		distances := strings.Fields(line)
		leftDistance, err := strconv.Atoi(distances[Left])
		check(err)
		rightDistance, err := strconv.Atoi(distances[Right])
		check(err)

		leftDistances = append(leftDistances, leftDistance)
		rightDistances = append(rightDistances, rightDistance)
	}
	err := file.Close()
	check(err)

	return leftDistances, rightDistances, scanner.Err()
}

func fileExists(file_distances string) {
	if _, err := os.Stat(file_distances); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File: %s does not exist\n", file_distances)
	}
}

func getDistancesFilePathFromArgs() string {
	// Get cli-args
	args := os.Args

	if len(args) < 2 || len(args) > 2 {
		fmt.Println("Please provide a path to the Chief Historian's location IDs")
		os.Exit(1)
	}

	fileDistancesPath := args[1]
	return fileDistancesPath
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}