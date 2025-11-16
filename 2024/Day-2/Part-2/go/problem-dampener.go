package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocmath "advent-of-code/aocutil/go/aoc/math"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
)

const DayPart = "Day 2 - Part 2"
const SolutionFormat = "Taking the problem dampener into account, the count of the safe reports now is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-2/resources/puzzle-input"

// Usage: go-app <PATH_TO_PUZZLE_FILE>
func main() {
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, ret [][]int) [][]int {
		ret = append(ret, []int{})

		s_report := strings.Fields(line)
		for _, v := range s_report {
			level, err := strconv.Atoi(v)
			aocutil.Check(err)
			ret[len(ret)-1] = append(ret[len(ret)-1], level)
		}

		return ret
	}

	reports, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	safeReports := 0
	for _, v := range reports {
		if checkReport(v) {
			safeReports++
			continue
		}
		for i := 0; i < len(v); i++ {
			vDampened := aocslice.RemoveIndex(v, i)
			if checkReport(vDampened) {
				safeReports++
				break
			}
		}
	}

	fmt.Printf(SolutionFormat, safeReports)
}

func checkReport(v []int) bool {
	levelsIncreasing := v[0] < v[1]
	for i := range v {
		if i < len(v)-1 {
			if levelsIncreasing != (v[i] < v[i+1]) {
				return false
			}
			diff := aocmath.Abs(v[i] - v[i+1])
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
}
