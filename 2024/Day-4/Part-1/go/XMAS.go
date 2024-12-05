package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strings"
)

const DayPart = "Day 4 - Part 1"
const SolutionFormat = "The word-count of XMAS is: %d\n"

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
			if wordPuzzle[y][x] == "X" {
				solution += searchXmasFromX(wordPuzzle, x, y)
			}
		}
	}

	fmt.Printf(SolutionFormat, solution)
}

// searchXmasFromX returns count of XMAS words found
func searchXmasFromX(wordPuzzle [][]string, x int, y int) int {
	xmasCount := 0
	for direction := 0; direction < 8; direction++ {
		xStep := directionalPendulum(direction, false)
		yStep := directionalPendulum(direction, true)

		if !checkBorders(wordPuzzle, x+xStep*3, y+yStep*3) { continue }

		if wordPuzzle[y+yStep][x+xStep] == "M" {
			if wordPuzzle[y+yStep*2][x+xStep*2] == "A" {
				if wordPuzzle[y+yStep*3][x+xStep*3] == "S" {
					xmasCount++
				}
			}
		}
	}

	return xmasCount
}

func directionalPendulum(direction int, yOffset bool) int {
	if yOffset {
		switch direction {
		case 2, 6:
			return 0
		case 0, 1, 7:
			return 1
		case 3, 4, 5:
			return -1
		}
	}

	switch direction {
	case 0, 4:
		return 0
	case 1, 2, 3:
		return 1
	case 5, 6, 7:
		return -1
	}

	panic("direction for directionalPendulum must be between 0 and 7.")
}

func checkBorders(wordPuzzle [][]string, x int, y int) bool {
	yGreatest := len(wordPuzzle) - 1
	xGreatest := len(wordPuzzle[0]) - 1

	if x < 0 || y < 0 || x > xGreatest || y > yGreatest {
		return false
	}

	return true
}
