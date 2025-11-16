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
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-5/resources/puzzle-input"

type manualInstructions struct {
	order   map[int][]int
	manuals [][]int
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, ret manualInstructions) manualInstructions {
		if strings.Contains(line, "|") {
			if ret.order == nil {
				ret.order = make(map[int][]int)
			}

			instruction := strings.Split(line, "|")
			intI1, err := strconv.Atoi(instruction[0])
			aocutil.Check(err)
			intI2, err := strconv.Atoi(instruction[1])
			aocutil.Check(err)
			ret.order[intI1] = append(ret.order[intI1], intI2)
		} else if line != "" {
			pages := aocslice.Atoi(strings.Split(line, ","))
			ret.manuals = append(ret.manuals, pages)
		}

		return ret
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	middleSum := 0
	for _, v := range puzzleInput.manuals {
		if iterPages(v, puzzleInput) {
			middleSum += v[len(v)/2]
		}
	}

	fmt.Printf(SolutionFormat, middleSum)
}

func iterPages(v []int, puzzleInput manualInstructions) bool {
	manualWasWrong := false
	for y := 0; y <= len(v); y++ {
		for _, page := range v {
			pageOrder := puzzleInput.order[page]
			for _, po := range pageOrder {
				if slices.Contains(v, po) {
					before := slices.Index(v, po)
					after := slices.Index(v, page)
					if before < after {
						manualWasWrong = true
						tmp := v[before]
						v[before] = v[after]
						v[after] = tmp
					}
				}
			}
		}
	}
	if manualWasWrong {
		return true
	}
	return false
}
