package args

import (
	"fmt"
	"os"
)

// GetPuzzleFilePathFromArgs gathers the puzzle file from args at index 1
func GetPuzzleFilePathFromArgs(dayPart string) string {
	// Get cli-args
	args := os.Args

	if len(args) < 2 || len(args) > 2 {
		fmt.Println("Please provide a path to the puzzle file of " + dayPart)
		os.Exit(1)
	}

	puzzleFilePath := args[1]
	return puzzleFilePath
}