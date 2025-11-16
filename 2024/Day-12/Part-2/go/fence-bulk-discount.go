package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	"advent-of-code/aocutil/go/aoc/map2D"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"time"
)

/* Conclusion
- Recursion on steroids... I don't even know how I managed to do this only with slight hints.

Takeaway:
- recursion with a returned tupel is insanely head-breaking
- Counting corners if you need sides of a shape/region in an array is easier than actually trying to find the sides
	- Finally managed to get rid of global variable for recursive function
- Taking shortcuts in the form of global variables is... well... also a solution :D
*/

const DayPart = "2024 Day 12 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-12/resources/puzzle-input"

var directions = map2D.GetDirectionalVectors()

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
			return 0, 0
		}
		visitedGlobally[coord] = true
		visitedLocally[coord] = true
		area := 1
		perimeter := 0
		for _, dir := range directions {
			checkPlot := map2D.AddVector(map2D.Coord{coord.X, coord.Y}, dir)
			a, p := getAreaData(puzzleInput, symbol, checkPlot, visitedLocally, visitedGlobally)
			area += a
			perimeter += map2D.SearchCorner(puzzleInput, symbol, coord, dir) + p
		}
		return area, perimeter
	}
}

// This is the original function written for this exact puzzle. Leaving it here for reference.
// Was later replaced with map2D.SearchCorner(...) because I felt this could be useful in the future
func searchCorner(puzzleInput [][]rune, symbol rune, coord map2D.Coord, dir map2D.Coord) int {
	pos1 := map2D.AddVector(coord, dir)
	pos3 := map2D.AddVector(pos1, map2D.RotateRight(dir))
	pos2 := map2D.AddVector(coord, map2D.RotateRight(dir))
	s1, s2, s3 := '.', '.', '.'
	if map2D.WithinBounds(puzzleInput, pos1) {
		s1 = puzzleInput[pos1.Y][pos1.X]
	}
	if map2D.WithinBounds(puzzleInput, pos2) {
		s2 = puzzleInput[pos2.Y][pos2.X]
	}
	if map2D.WithinBounds(puzzleInput, pos3) {
		s3 = puzzleInput[pos3.Y][pos3.X]
	}
	if (s1 == symbol && s2 == symbol && s3 != symbol) || (s1 != symbol && s2 != symbol) {
		return 1
	}
	return 0
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
