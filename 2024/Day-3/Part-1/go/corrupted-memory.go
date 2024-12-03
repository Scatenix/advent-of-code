package main

import (
	aocargs "advent-of-code/aocutil/go/aoc/args"
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocmath "advent-of-code/aocutil/go/aoc/math"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const DayPart = "Day 3 - Part 1"
const SolutionFormat = "The sum of all multiplications is: %d\n"

func main() {
	puzzleFilePath := aocargs.GetPuzzleFilePathFromArgs(DayPart)
	aocutil.Check(aocio.FileExists(puzzleFilePath))
	file, err := os.Open(puzzleFilePath)
	aocutil.Check(err)
	aocmath.Abs(1)

	puzzleLineHandler := func(line string, ret string) string {
		ret += line
		return ret
	}

	corruptMem, err := aocio.ReadPuzzleFile[string](file, puzzleLineHandler)
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
