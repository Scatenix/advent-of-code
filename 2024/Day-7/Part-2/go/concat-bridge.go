package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocmath "advent-of-code/aocutil/go/aoc/math"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DayPart = "2024 Day 7 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-7/resources/puzzle-input"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, totalResult int64) int64 {
		result, _ := strconv.ParseInt(strings.Split(line, ":")[0], 10, 64)
		operands := aocslice.Atoi64(strings.Split(line, " ")[1:])

		return totalResult + checkForResult(result, operands)
	}

	totalResult, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	fmt.Printf(SolutionFormat, totalResult)
}

func checkForResult(result int64, operands []int64) int64 {
	gaps := len(operands) - 1

	for i := 0; i < aocmath.Pow(3, gaps); i++ {
		check := operands[0]
		operators := aocmath.IntToBaseStringWithPadding(i, gaps, 3)
		for op, v := range operands[1:] {
			if operators[op] == '0' {
				check = check * v
			} else if operators[op] == '1' {
				check = check + v
			} else {
				check = aocmath.ConcatInt64(check, v)
			}
		}
		if check == result {
			return result
		}
	}
	return 0
}
