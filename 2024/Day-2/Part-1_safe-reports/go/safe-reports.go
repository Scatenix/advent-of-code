package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DayPart = "Day 2 - Part 1"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	puzzleFilePath := getPuzzleFilePathFromArgs()
	fileExists(puzzleFilePath)
	file, err := os.Open(puzzleFilePath)
	check(err)

	reports, err := ReadReports(file)
	check(err)

	safeReports := 0;
	for _, v := range reports {
		levelsIncreasing := v[0] < v[1]
		for i := range v {
			if i < len(v) - 1 {
				if levelsIncreasing != (v[i] < v[i+1]) {
					break
				}

				diff := abs(v[i] - v[i+1])
				if diff < 1 || diff > 3 {
					break
				}
			} else {
				safeReports++
			}
		}
	}

	fmt.Printf("The count of the safe reports is: %d\n", safeReports)
}

func ReadReports(file *os.File) ([][]int, error) {
	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		reports = append(reports, []int{})

		s_report := strings.Fields(lineText)
		for _, v := range s_report {
			level, err := strconv.Atoi(v)
			check(err)
			reports[len(reports)-1] = append(reports[len(reports)-1], level)
		}
	}
	err := file.Close()
	check(err)

	return reports, scanner.Err()
}

func fileExists(file string) {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File: %s does not exist\n", file)
	}
}

func getPuzzleFilePathFromArgs() string {
	// Get cli-args
	args := os.Args

	if len(args) < 2 || len(args) > 2 {
		fmt.Println("Please provide a path to the puzzle file of " + DayPart)
		os.Exit(1)
	}

	puzzleFilePath := args[1]
	return puzzleFilePath
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}