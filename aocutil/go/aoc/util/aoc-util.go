package util

import (
	aocargs "advent-of-code/aocutil/go/aoc/args"
	aocio "advent-of-code/aocutil/go/aoc/io"
	"os"
)

// Check for error being present and panic in case it is
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// AocSetup provides a basic setup for Advent of Code challenges.
// Puzzle file will be gathered from cli args.
func AocSetup(dayPart string) *os.File {
	puzzleFilePath := aocargs.GetPuzzleFilePathFromArgs(dayPart)
	aocio.FileExists(puzzleFilePath)
	file, err := os.Open(puzzleFilePath)
	Check(err)
	return file
}
