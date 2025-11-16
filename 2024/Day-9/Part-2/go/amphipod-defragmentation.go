package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strings"
	"time"
)

const DayPart = "2024 Day 9 - Part 2"
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
	digits = shrinkDigitsDefragmented(digits)
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

func shrinkDigitsDefragmented(digits []int) []int {
	for i := len(digits) - 1; i > -1; i-- {
		if digits[i] != EmptySpace && i >= 1 {
			digitCount := 1
			for digits[i] == digits[i-1] {
				digitCount++
				i--
				if i <= 1 {
					break
				}
			}
			freeSpace := 0
			for j := 0; j < len(digits); j++ {
				if j >= i {
					break
				}
				if digits[j] == EmptySpace {
					freeSpace++
				} else {
					freeSpace = 0
				}
				if freeSpace == digitCount {
					digits = aocslice.ReplaceRange(digits, j-digitCount+1, digits[i:i+digitCount])

					placeHolderSlice := make([]int, digitCount)
					for i := 0; i < digitCount; i++ {
						placeHolderSlice[i] = EmptySpace
					}
					digits = aocslice.ReplaceRange(digits, i, placeHolderSlice)
					break
				}
			}
		}
	}
	return digits
}

func calcChecksum(digits []int) int {
	sum := 0
	for i, digit := range digits {
		if digit != EmptySpace {
			sum += i * digit
		}
	}
	return sum
}
