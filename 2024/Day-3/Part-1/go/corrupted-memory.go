package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

const DayPart = "Day 3 - Part 1"
const SolutionFormat = "The sum of all multiplications is: %d\n"

// Usage: go-app <PATH_TO_PUZZLE_FILE>
func main() {
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, ret string) string {
		ret += line
		return ret
	}

	corruptMem, err := aocio.ReadPuzzleFile[string](puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	sumOfAllMultiplications := 0;
	rgxMulCommand := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	// the param n means how many matches should be made at max. -1 means infinite matches are allowed.
	mulCommands := rgxMulCommand.FindAllString(corruptMem, -1)

	rgxMulNumbers := regexp.MustCompile(`\d{1,3}`)
	for _, mul := range mulCommands {
		rgxMulNumbers := rgxMulNumbers.FindAllString(mul, -1)

		n1, err := strconv.Atoi(rgxMulNumbers[0])
		aocutil.Check(err)
		n2, err := strconv.Atoi(rgxMulNumbers[1])
		aocutil.Check(err)

		sumOfAllMultiplications += n1 * n2
	}

	fmt.Printf(SolutionFormat, sumOfAllMultiplications)
}
