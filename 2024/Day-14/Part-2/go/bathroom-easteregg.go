package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	"advent-of-code/aocutil/go/aoc/map2D"
	"advent-of-code/aocutil/go/aoc/math"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/* Conclusion 2024 Day 14 Part 2
Simple takeaway here: When part 1 says: do it in 100 seconds, that does not mean that part 2 also has to be done in 100 seconds,
if it is not explicitly stated...
Overall I think this solution is somewhat messy, but it is not too long and does it's job.
The variable names got a bit out of hand in this one. But I am not to fond of this puzzle (lots of silly errors today...)
*/

const DayPart = "2024 Day 14 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-14/resources/puzzle-input"

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
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col []string) []string {
		col = append(col, line)
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	robots := parseInput(puzzleInput)
	roboMap := make([]map2D.Coord, len(robots))
	for i, r := range robots {
		roboMap[i] = map2D.Coord{X: r.start.X, Y: r.start.Y}
	}

	searchForTheTree(robots, roboMap)
}

// Exits the app when the tree is found.
func searchForTheTree(robots []robot, roboMap []map2D.Coord) {
	// The loop count does not matter at all here, since we exit as soon as we have our answer. And the answer WILL come at some point
	// in part one, we would loop to 100 here
	// Also, as opposed to Part 1, we need a know the state of each robot for each second.
	// That is why the for loops are the other way around compared to part 1, where this does not matter,
	// because we simply needed to know in which quadrant each individual robot ended up.
	for second := 0; true; second++ {
		for i, r := range robots {
			newStart := map2D.AddVector(roboMap[i], r.vec)

			if newStart.X >= XWidth {
				newStart.X = newStart.X - XWidth
			} else if newStart.X < 0 {
				newStart.X = XWidth + newStart.X
			}
			if newStart.Y >= YWidth {
				newStart.Y = newStart.Y - YWidth
			} else if newStart.Y < 0 {
				newStart.Y = YWidth - math.Abs(newStart.Y)
			}
			roboMap[i] = newStart
		}
		if isThereAStraightLine(roboMap) {
			// second+1 because the for loop starts at 0, but the actual second is one ahead
			fmt.Printf(SolutionFormat, second+1)
			os.Exit(0)
		}
	}
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

// This function is honestly kind of "cheated" because I knew from the subreddit how the tree looked like.
// Of course, it makes sense that we have at one row of 10 robots in a line at the x-axis to see a tree,
// but the tree could have been also a lot smaller
// Also scanning the y-axis instead would probably make even more sense, since a tree is taller than it is wide
func isThereAStraightLine(Cmap []map2D.Coord) bool {
	map2d := convertTo2D(Cmap)

	xCount := 0
	for _, y := range map2d {
		for x := 0; x < len(y); x++ {
			if y[x] == "X" {
				xCount++
				if xCount > 10 {
					map2D.PrintMap(map2d, false)
					return true
				}
			} else {
				xCount = 0
			}
		}
	}
	return false
}

func convertTo2D(Cmap []map2D.Coord) [][]string {
	map2d := make([][]string, 103)
	for i := range map2d {
		map2d[i] = make([]string, 101)
		for j := range map2d[i] {
			map2d[i][j] = "."
		}
	}

	for _, c := range Cmap {
		map2d[c.Y][c.X] = "X"
	}
	return map2d
}
