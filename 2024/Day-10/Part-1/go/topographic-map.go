package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	"advent-of-code/aocutil/go/aoc/map2D"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strings"
	"time"
)

const DayPart = "2024 Day 10 - Part 1"
const SolutionFormat = ">>> The sum of trailhead scores is: %d\n"

var StartVector = map2D.Coord{1, 0}
var Score = 0
var Paths = make(map[int][]map2D.Coord, 0)
var PathId = 0

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
    defer aocperf.TimeTracker(time.Now(), "Main")
    defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col [][]int) [][]int {
		col = append(col, aocslice.Atoi(strings.Split(line, "")))
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	score := 0
	for y := 0; y < len(puzzleInput); y++ {
		for x := 0; x < len(puzzleInput[y]); x++ {
			if puzzleInput[y][x] == 0 {
				score += searchPath(puzzleInput, map2D.Coord{x, y}, 0)
			}
		}
	}

	print(Paths)
	fmt.Printf(SolutionFormat, Score)
}

func searchPath(input [][]int, pos map2D.Coord, score int) int {
	vec := StartVector
	for i := 0; i < 4; i++ {

		nextPos := map2D.AddVector(pos, vec)
		if map2D.OutOfBounds(input, nextPos) {
			continue
		}

		if input[pos.Y][pos.X]+1 == input[nextPos.Y][nextPos.X] {
			if input[nextPos.Y][nextPos.X] == 9 {
				// found end of path
				Score++
				continue
			}

			Paths[Score] = append(Paths[Score], nextPos)
			// found next path way
			score = score + searchPath(input, nextPos, score)
		}

		vec = rotate(vec)
	}
	return score

}

// this function is tied to StartVector
func rotate(vec map2D.Coord) map2D.Coord {
	if vec.X == 1 && vec.Y == 0 {
		return map2D.Coord{0, 1}
	} else if vec.X == 0 && vec.Y == 1 {
		return map2D.Coord{-1, 0}
	} else if vec.X == -1 && vec.Y == 0 {
		return map2D.Coord{0, -1}
	} else if vec.X == 0 && vec.Y == -1 {
		// using this as an exit condition, since we alredy used vec(1,0)(=StartVector)
		return map2D.Coord{-1, -1}
	}
	panic("rotate input is not allowed")
}