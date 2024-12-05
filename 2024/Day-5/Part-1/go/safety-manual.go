package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
)

const DayPart = "Day 5 - Part 1"
const SolutionFormat = "The solution is: %d\n"

type manualInstructions struct {
	order map[int]int
	manuals [][]int
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, ret manualInstructions) manualInstructions {
		if strings.Contains(line, "|") {
			instruction := strings.Split(line, "|")
			intI1, err := strconv.Atoi(instruction[0]); aocutil.Check(err)
			intI2, err := strconv.Atoi(instruction[1]); aocutil.Check(err)
			ret.order[intI1] := intI2
		} else if line != "" {
			pages := aocslice.Atoi(strings.Split(line, " "))
			ret.manuals = append(ret.manuals, pages)
		}

		return ret
	}

	var mi = manualInstructions{order: make(map[int]int), manuals: make([][]int, 0)}
	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler, &mi)
	aocutil.Check(err)

	solution := 0;
	for _, v := range puzzleInput.manuals {
		print(v)
	}

	fmt.Printf(SolutionFormat, solution)
}
