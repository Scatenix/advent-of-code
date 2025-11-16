package args

import (
	"fmt"
	"os"
)

// GetPuzzleFilePathFromArgs gathers the puzzle file from args at index 1
func GetPuzzleFilePathFromArgs(dayPart string) string {
	// Get cli-args
	args := os.Args

	if len(args) < 2 {
		// Immediately return to use fallback puzzle file
		return ""
	}

	if len(args) > 2 {
		fmt.Println("Please provide a path as the only argument to this go app to the puzzle input file of " + dayPart)
		fmt.Println("Omitting the first argument will use a fallback puzzle file.")
		os.Exit(1)
	}

	puzzleFilePath := args[1]
	return puzzleFilePath
}
