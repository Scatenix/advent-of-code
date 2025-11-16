package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	"advent-of-code/aocutil/go/aoc/map2D"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"time"
)

const DayPart = "2024 Day 12 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-12/resources/puzzle-input"

var directions = []map2D.Coord{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col [][]rune) [][]rune {
		col = append(col, []rune(line))
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	area, perimeter := getFenceNumbers(puzzleInput)
	totalFenceCost := calculateTotalFenceCost(area, perimeter)

	fmt.Printf(SolutionFormat, totalFenceCost)
}

// getAreadData returns (area, perimeter)
func getAreaData(puzzleInput [][]rune, symbol rune, coord map2D.Coord, visitedLocally, visitedGlobally map[map2D.Coord]bool) (int, int) {
	if visitedLocally[coord] == true {
		return 0, 0
	} else {
		if map2D.OutOfBounds(puzzleInput, coord) || puzzleInput[coord.Y][coord.X] != symbol {
			return 0, 1
		}
		visitedGlobally[coord] = true
		visitedLocally[coord] = true
		area := 1
		perimeter := 0
		for _, dir := range directions {
			checkPlot := map2D.AddVector(map2D.Coord{coord.X, coord.Y}, dir)
			a, p := getAreaData(puzzleInput, symbol, checkPlot, visitedLocally, visitedGlobally)
			area += a
			perimeter += p
		}
		return area, perimeter
	}
}

func getFenceNumbers(puzzleInput [][]rune) (map[int]int, map[int]int) {
	fenceArea := make(map[int]int)
	fencePerimeter := make(map[int]int)
	visited := make(map[map2D.Coord]bool)
	id := 0
	for y, plotRow := range puzzleInput {
		for x, plot := range plotRow {
			if visited[map2D.Coord{x, y}] {
				continue
			}

			a, p := getAreaData(puzzleInput, plot, map2D.Coord{x, y}, make(map[map2D.Coord]bool), visited)
			fenceArea[id] = a
			fencePerimeter[id] = p
			id++
		}
	}
	return fenceArea, fencePerimeter
}

func calculateTotalFenceCost(fenceArea, fencePerimeter map[int]int) int {
	totalCost := 0
	for p := range fenceArea {
		totalCost += fenceArea[p] * fencePerimeter[p]
	}
	return totalCost
}
