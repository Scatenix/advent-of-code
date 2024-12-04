package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
)

const DayPart = "Day 4 - Part 2"
const SolutionFormat = "The solution is: %d\n"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, ret string) string {
		ret += line
		return ret
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	solution := 0;
	for _, v := range puzzleInput {
		num, err := strconv.Atoi(string(v))
		aocutil.Check(err)
		solution += num
	}

	fmt.Printf(SolutionFormat, solution)
}
