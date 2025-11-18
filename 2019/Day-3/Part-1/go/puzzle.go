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
func determinePanelSize(puzzleInput [][]string) (int, int) {
	height0 := 0
	width0 := 0
	height1 := 0
	width1 := 0
	curPosX := 0
	curPosY := 0
	for _, v := range puzzleInput[0] {
		steps, _ := strconv.Atoi(v[1:])
		x, y := direction(string(v[0]))
		curPosX += steps * x
		curPosY += steps * y
		height0 = math.Greater(math.Abs(height0), math.Abs(curPosY))
		width0 = math.Greater(math.Abs(width0), math.Abs(curPosX))
	}
	curPosX = 0
	curPosY = 0
	for _, v := range puzzleInput[1] {
		steps, _ := strconv.Atoi(v[1:])
		x, y := direction(string(v[0]))
		curPosX += steps * x
		curPosY += steps * y
		height1 = math.Greater(math.Abs(height1), math.Abs(curPosY))
		width1 = math.Greater(math.Abs(width1), math.Abs(curPosX))
	}

	// Multiply by 2 because we now know the max steps in one direction, but we set our origin point exactly at the middle of the 2D field
	// Plus 1 because we start to "walk" at 0, but set the length of the panel with these values.
	return (math.Greater(math.Abs(height0), math.Abs(height1)) * 2) + 1, (math.Greater(math.Abs(width0), math.Abs(width1)) * 2) + 1
}
