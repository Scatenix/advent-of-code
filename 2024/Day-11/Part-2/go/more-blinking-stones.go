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
const Blinks = 45

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
		stones += blink(v, 0)
	}

	//fmt.Printf(SolutionFormat, puzzleInput.Len())
	fmt.Printf("--------------")
	fmt.Printf(SolutionFormat, stones)
}

func blink(val, blinked int) int {
	if blinked == Blinks {
		return 1
	}
	strVal := strconv.Itoa(val)
	if val == 0 {
		return blink(1, blinked+1)
	} else if len(strVal)%2 == 0 {
		firstHalf, _ := strconv.Atoi(strVal[0 : len(strVal)/2])
		secondHalf, _ := strconv.Atoi(strVal[len(strVal)/2:])
		tmp := blink(firstHalf, blinked+1)
		return blink(secondHalf, blinked+1) + tmp
	} else {
		return blink(val*2024, blinked+1)
	}
}
