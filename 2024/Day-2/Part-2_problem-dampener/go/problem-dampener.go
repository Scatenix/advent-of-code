package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DayPart = "Day 2 - Part 2"

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
		if checkReport(v) {
			safeReports++
			continue
		}
		for i := 0; i < len(v); i++ {
			vDampened := removeIndex(v, i)
			if checkReport(vDampened) {
				safeReports++
				break
			}
		}
	}

	fmt.Printf("Taking the problem dampener into account, the count of the safe reports is: %d\n", safeReports)
}

func checkReport(v []int) bool {
	levelsIncreasing := v[0] < v[1]
	for i := range v {
		if i < len(v)-1 {
			if levelsIncreasing != (v[i] < v[i+1]) {
				return false
			}
			diff := abs(v[i] - v[i+1])
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
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

// =============================================================================================================
// A deep dive into properly removing an item at index from a slice (vector)
// =============================================================================================================

// A proper pure removeIndex function which does not modify the input but rather creates a deep copy
// This is what one would expect from this function.
// However, if the input slice is also used for the output slice, a version without deep copy can
// be a better choice performance and space wise
func removeIndex(s []int, index int) []int {
	// create a completely new slice with the length of the input slice
	deepCopySlice := make([]int, 0, len(s)-1)
	// append the input slice until excluding index to the new slice
	deepCopySlice = append(deepCopySlice, s[:index]...)
	// append the input slice from index+1 to the end to the new slice
	return append(deepCopySlice, s[index+1:]...)
}

// Shallow copy only!!!
// Assumption: Use only if the input slice will also be used for the returned slice (i.e. mySlice = removeIndex(mySlice, 4)
// ----------------------
// Since the slice (vector) is already implicitly a pointer (because all arrays are just that),
// This function is essentially impure by default.
// the original slice s will be modified as well, leading to drastic side effects,
// when trying to keep the original slice as is and using the output for a new slice.
// A deep copy is required here for a expected function
func removeIndexImpure(s []int, index int) []int {
 	// needed instead if the slice type is a pointer or a struct with pointer fields to avoid memory leaks:
	//copy(a[i:], a[i+1:])
	//a[len(a)-1] = nil // or the zero value of T
	//a = a[:len(a)-1]

	// Fine for primitives
	return append(s[:index], s[index+1:]...)
}