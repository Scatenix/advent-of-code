package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"slices"
	"strings"
	"time"
)

const DayPart = "2024 Day 9 - Part 1"
const SolutionFormat = ">>> The checksum for the compacted filesystem is : %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-9/resources/puzzle-input"
const EmptySpace = -1

var id = 0

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col string) string {
		return line
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)
	digits := inputToDigits(strings.Split(puzzleInput, "\n")[0])
	digits = shrinkDigits(digits)
	checksum := calcChecksum(digits)

	fmt.Printf(SolutionFormat, checksum)
}

func inputToDigits(input string) []int {
	digits := aocslice.Atoi(strings.Split(input, ""))

	var result = make([]int, 0, len(digits))
	for i := 0; i < len(digits); i++ {
		for j := 0; j < digits[i]; j++ {
			result = append(result, id)
		}
		i++
		if i < len(digits) {
			for j := 0; j < digits[i]; j++ {
				result = append(result, EmptySpace)
			}
		}
		id++
	}
	return result
}

func shrinkDigits(digits []int) []int {
	for i := 0; i < len(digits); i++ {
		if digits[i] == EmptySpace {
			digits = slices.Replace(digits, i, i+1, digits[len(digits)-1])
			digits = digits[:len(digits)-1]
		}
		for digits[len(digits)-1] == EmptySpace {
			digits = digits[:len(digits)-1]
		}
	}
	return digits
}

func calcChecksum(digits []int) int {
	sum := 0
	for i, digit := range digits {
		sum += i * digit
	}
	return sum
}
