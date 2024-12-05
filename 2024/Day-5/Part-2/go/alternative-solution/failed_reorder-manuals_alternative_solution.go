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

const DayPart = "Day 5 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"

type manualInstructions struct {
	order []int
	orderMap map[int]int
	manuals [][]int
}

/*
I feel like the idea was great, but I could not get it to work on the big dataset (worked on the small example tough)
Idea was to get a map with a perfectly ordered list of pages on how they need to be ordered for the manuals (for the manuals)
and using that to order the actual manuals.
I think the ordered list of pages was correct (but very hard to say on the big data)
I guess the error is somewhere in the last for loop

While maybe not more efficient, my first attempt was at least less code and did even work
 */

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col manualInstructions) manualInstructions {
		if strings.Contains(line, "|") {
			instruction := strings.Split(line, "|")
			intI1, err := strconv.Atoi(instruction[0]); aocutil.Check(err)
			intI2, err := strconv.Atoi(instruction[1]); aocutil.Check(err)

			if !slices.Contains(col.order, intI1) {
				col.order = append(col.order, intI1)
			}
			if !slices.Contains(col.order, intI2) {
				col.order = append(col.order, intI2)
			}

			i1 := slices.Index(col.order, intI1)
			i2 := slices.Index(col.order, intI2)
			if i1 > i2 {
				col.order = aocslice.RemoveIndexImpure(col.order, i1)
				col.order = slices.Insert(col.order, i2, intI1)
			}
		} else if line != "" {
			pages := aocslice.Atoi(strings.Split(line, ","))
			col.manuals = append(col.manuals, pages)
		}

		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)
	puzzleInput.orderMap = make(map[int]int)
	for i, v := range puzzleInput.order {
		puzzleInput.orderMap[v] = i
	}

	middleSum := 0;
	for _, v := range puzzleInput.manuals {
		tmpSlice := make([]int, len(puzzleInput.orderMap))
		for _, page := range v {
			tmpSlice[puzzleInput.orderMap[page]] = page
		}
		tmpSlice = aocslice.RemoveZeros(tmpSlice)
		if !compareSlices(tmpSlice, v) {
			middleSum += tmpSlice[len(tmpSlice)/2]
		}
	}

	fmt.Printf(SolutionFormat, middleSum)
}

func compareSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}