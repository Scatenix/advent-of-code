package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DayPart = "2024 Day 11 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"
const Blinks = 75

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

	stones := 0

	for _, v := range puzzleInput {
		stones += blink(v, 0, make(map[[2]int]int))
	}

	fmt.Printf(SolutionFormat, stones)
}

// TODO: add version without memoization to here, to be able to see how exactly it works!
func blink(val, blinked int, memoizeMap map[[2]int]int) int {
	if memoizeMap[[2]int{val, blinked}] != 0 {
		return memoizeMap[[2]int{val, blinked}]
	}
	if blinked == Blinks {
		return 1
	}
	strVal := strconv.Itoa(val)
	if val == 0 {
		return blink(1, blinked+1, memoizeMap)
	} else if len(strVal)%2 == 0 {
		firstHalf, _ := strconv.Atoi(strVal[0 : len(strVal)/2])
		secondHalf, _ := strconv.Atoi(strVal[len(strVal)/2:])
		tmp1 := blink(firstHalf, blinked+1, memoizeMap)
		tmp2 := blink(secondHalf, blinked+1, memoizeMap)
		memoizeMap[[2]int{val, blinked}] += tmp1 + tmp2
		return tmp1 + tmp2
	} else {
		return blink(val*2024, blinked+1, memoizeMap)
	}
}
