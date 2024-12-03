package main

import (
	aocargs "advent-of-code/aocutil/go/aoc/args"
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocmath "advent-of-code/aocutil/go/aoc/math"
	aocutil "advent-of-code/aocutil/go/aoc/util"

	"os"
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

	//sumMultiplications := 0;
	print(corruptMem)


	//fmt.Printf(SolutionFormat, sumMultiplications)
}


