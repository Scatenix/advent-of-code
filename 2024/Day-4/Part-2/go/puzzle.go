package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strings"
)

const DayPart = "Day 4 - Part 2"
const SolutionFormat = "The cross-count of MASes is: %d\n"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, ret [][]string) [][]string {
		xLine := strings.Split(line, "")
		ret = append(ret, xLine)
		return ret
	}

	wordPuzzle, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	fmt.Println(wordPuzzle)

	solution := 0;
	for y := range wordPuzzle {
		for x := range wordPuzzle[y] {
			if wordPuzzle[y][x] == "A" {
				solution += searchXmasCrossFromA(wordPuzzle, x, y)
			}
		}
	}

	fmt.Printf(SolutionFormat, solution)
}

// searchXmasCrossFromA returns count of MAS Crosses found
func searchXmasCrossFromA(wordPuzzle [][]string, x int, y int) int {
	if !checkBorders(wordPuzzle, x, y) { return 0 }

	mas1 := wordPuzzle[y+1][x-1] + wordPuzzle[y][x] + wordPuzzle[y-1][x+1]
	mas2 := wordPuzzle[y+1][x+1] + wordPuzzle[y][x] + wordPuzzle[y-1][x-1]

	if (mas1 == "MAS" || mas1 == "SAM") && (mas2 == "MAS" || mas2 == "SAM") {
		return 1
	}
	return 0
}

func checkBorders(wordPuzzle [][]string, x int, y int) bool {
	yGreatest := len(wordPuzzle) - 1
	xGreatest := len(wordPuzzle[0]) - 1

	if x-1 < 0 || y-1 < 0 || x+1 > xGreatest || y+1 > yGreatest {
		print ("out")
		return false
	}

	return true
}
