package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	"advent-of-code/aocutil/go/aoc/math"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/* Conclusion
I like my solution overall, as I think it's near the best that is achievable through the "drawing on an array" approach.

However, A solution where only coordinates are tracked in a Set, where wire 1 is going, without the overhead of empty space,
is way cheaper and easier, since this approach allows for negative coordinates.
A hash-set is probably the way to go since we will need to constantly check wire 2 against wire 1 with unpredictable positions of the intersections.
*/

const DayPart = "2019 Day 3 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2019/Day-3/resources/puzzle-input"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col [][]string) [][]string {
		path := strings.Split(line, ",")
		col = append(col, path)
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	ph, pw := determinePanelSize(puzzleInput)
	frontPanel := make([][]int, pw)
	for i := range frontPanel {
		frontPanel[i] = make([]int, ph)
	}

	midPointH := ph / 2
	midPointW := pw / 2
	x := midPointW
	y := midPointH
	frontPanel[x][y] = 9

	possibleSolution := make([]int, 0)
	for _, v := range puzzleInput[0] {
		dir := string(v[0])
		length, _ := strconv.Atoi(v[1:])

		for i := 0; i < length; i++ {
			vecx, vecy := direction(dir)
			x += vecx
			y += vecy

			frontPanel[x][y] = 1
		}
	}

	x = midPointW
	y = midPointH

	for _, v := range puzzleInput[1] {
		dir := string(v[0])
		length, _ := strconv.Atoi(v[1:])

		for i := 0; i < length; i++ {
			vecx, vecy := direction(dir)
			x += vecx
			y += vecy

			// The size of front panel is exactly as large as it needs to be for wire 1.
			// Thus, if wire 2 is about to be outside that scope, we do not care about it.
			if math.Abs(x) > len(frontPanel) || math.Abs(y) > len(frontPanel[0]) {
				break
			}

			// Comments are for debugging (printing the map). Printing takes very long on the actual, large puzzle input.
			if frontPanel[x][y] == 1 {
				//frontPanel[x][y] = 8
				possibleSolution = append(possibleSolution, math.Abs(x-midPointW)+math.Abs(y-midPointH))
			} /*else {
				frontPanel[x][y] = 2
			}*/
		}
	}

	//map2D.PrintCleanMap(frontPanel, false)
	fmt.Printf(SolutionFormat, math.GetSmallest(possibleSolution))
}

func direction(dir string) (int, int) {
	if dir == "U" {
		return -1, 0
	} else if dir == "D" {
		return 1, 0
	} else if dir == "R" {
		return 0, 1
	} else {
		return 0, -1
	}
}

// determinePanelSize get the exact needed panel size to optimize memory usage.
// We don't need the size for the second wire, as we do not care for wire 2 when it is outside wire 1's scope.
func determinePanelSize(puzzleInput [][]string) (int, int) {
	height, width := 0, 0
	curPosX, curPosY := 0, 0

	for _, v := range puzzleInput[0] {
		steps, _ := strconv.Atoi(v[1:])
		x, y := direction(string(v[0]))
		curPosX += steps * x
		curPosY += steps * y
		height = max(math.Abs(height), math.Abs(curPosY))
		width = max(math.Abs(width), math.Abs(curPosX))
	}

	// Multiply by 2 because we now know the max steps in one direction, but we set our origin point exactly at the middle of the 2D field
	return height*2 + 1, width*2 + 1
}
