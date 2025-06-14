package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

const DayPart = "2024 Day 11 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col []int) []int {
		col = append(col, aocslice.Atoi(strings.Split(line, " "))...)
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	for i := 0; i < 25; i++ {
		for s := 0; s < len(puzzleInput); s++ {
			strVal := strconv.Itoa(puzzleInput[s])
			if puzzleInput[s] == 0 {
				puzzleInput[s] = 1
			} else if len(strVal)%2 == 0 {
				firstHalf, _ := strconv.Atoi(strVal[0 : len(strVal)/2])
				secondHalf, _ := strconv.Atoi(strVal[len(strVal)/2:])
				puzzleInput[s] = firstHalf
				puzzleInput = slices.Insert(puzzleInput, s+1, secondHalf)
				s++
			} else {
				puzzleInput[s] = puzzleInput[s] * 2024
			}
		}
	}

	fmt.Printf(SolutionFormat, len(puzzleInput))
}
