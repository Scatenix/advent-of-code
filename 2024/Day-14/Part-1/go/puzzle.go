package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	"advent-of-code/aocutil/go/aoc/map2D"
	"advent-of-code/aocutil/go/aoc/math"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/* Conclusion
Massive takeaway: Check your constants. My first solution that solved the example was correct.
Yet I thought it was incorrect on the real data, but I just forgot to adjust the constants, that
define the playing field size
*/

const DayPart = "2024 Day 14 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"

const XWidth = 101
const YWidth = 103

type robot struct {
	start map2D.Coord
	vec   map2D.Coord
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col []string) []string {
		col = append(col, line)
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	robots := parseInput(puzzleInput)

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		end := map2D.Coord{X: r.start.X, Y: r.start.Y}
		for i := 0; i < 100; i++ {
			end = map2D.AddVector(end, r.vec)
		}

		warpedPos := map2D.Coord{0, 0}
		if end.X > 0 {
			warpedPos.X = end.X % XWidth
		} else if end.X < 0 {
			pos := XWidth - math.Abs(end.X%XWidth)
			if pos == XWidth {
				warpedPos.X = 0
			} else {
				warpedPos.X = pos
			}
		} // if end.X is 0, leave warpedPos at 0,0
		if end.Y > 0 {
			warpedPos.Y = end.Y % YWidth
		} else if end.Y < 0 {
			pos := YWidth - math.Abs(end.Y%YWidth)
			if pos == YWidth {
				warpedPos.Y = 0
			} else {
				warpedPos.Y = pos
			}
		} // if end.Y is 0, leave warpedPos at 0,0

		midX := XWidth / 2
		midY := YWidth / 2

		if warpedPos.X > midX && warpedPos.Y > midY {
			q4++
		} else if warpedPos.X < midX && warpedPos.Y < midY {
			q1++
		} else if warpedPos.X > midX && warpedPos.Y < midY {
			q2++
		} else if warpedPos.X < midX && warpedPos.Y > midY {
			q3++
		}
	}

	fmt.Printf(SolutionFormat, q1*q2*q3*q4)
}

func parseInput(puzzleInput []string) []robot {
	robots := make([]robot, 0)
	for _, line := range puzzleInput {
		robots = append(robots, robot{})
		start := strings.Split(strings.Split(line, " ")[0], "=")[1]
		vec := strings.Split(strings.Split(line, " ")[1], "=")[1]

		robots[len(robots)-1].start.X, _ = strconv.Atoi(strings.Split(start, ",")[0])
		robots[len(robots)-1].start.Y, _ = strconv.Atoi(strings.Split(start, ",")[1])

		robots[len(robots)-1].vec.X, _ = strconv.Atoi(strings.Split(vec, ",")[0])
		robots[len(robots)-1].vec.Y, _ = strconv.Atoi(strings.Split(vec, ",")[1])
	}
	return robots
}
