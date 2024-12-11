package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DayPart = "2024 Day 11 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col *list.List) *list.List {
		col = list.New()
		for _, v := range aocslice.Atoi(strings.Split(line, " ")) {
			col.PushBack(v)
		}
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	for i := 0; i < 25; i++ {
		for e := puzzleInput.Front(); e != nil; e = e.Next() {
			strVal := strconv.Itoa(e.Value.(int))
			if e.Value == 0 {
				e.Value = 1
			} else if len(strVal)%2 == 0 {
				firstHalf, _ := strconv.Atoi(strVal[0 : len(strVal)/2])
				secondHalf, _ := strconv.Atoi(strVal[len(strVal)/2:])
				e.Value = firstHalf
				puzzleInput.InsertAfter(secondHalf, e)
				e = e.Next()
			} else {
				e.Value = e.Value.(int) * 2024
			}
		}
		println("blink ", i)
	}

	fmt.Printf(SolutionFormat, puzzleInput.Len())
}
