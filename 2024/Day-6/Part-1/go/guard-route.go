package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strings"
	"time"
)

const DayPart = "2024 Day 6 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-6/resources/puzzle-input"

const Guard = "^"
const PassMarker = "X"
const Wall = "#"

type coord struct {
	x int
	y int
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col [][]string) [][]string {
		col = append(col, strings.Split(line, ""))
		return col
	}

	guardMap, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	pos := locateStart(guardMap)
	// We start by walking down, because the map was read backwards
	walkVec := coord{x: 0, y: -1}

	fmt.Println("Before walking:")
	printMap(guardMap)

	walk(guardMap, pos, walkVec)
	walkedPositions := countWalkedPositions(guardMap)

	fmt.Println("\nAfter walking:")
	printMap(guardMap)
	fmt.Printf(SolutionFormat, walkedPositions)
}

// locateStart returnx (x, y)
func locateStart(guardMap [][]string) coord {
	for y := range guardMap {
		for x := range guardMap[y] {
			if guardMap[y][x] == Guard {
				return coord{x: x, y: y}
			}
		}
	}
	panic("could not locate start")
}

// walk the map and leave an x
func walk(guardMap [][]string, pos coord, vec coord) {
	if checkForMapEnd(guardMap, pos) {
		guardMap[pos.y][pos.x] = PassMarker
		return
	}

	if guardMap[pos.y+vec.y][pos.x+vec.x] == Wall {
		vec = turnLeft(vec)
	}

	guardMap[pos.y][pos.x] = PassMarker
	pos = addVec(pos, vec)
	walk(guardMap, pos, vec)
}

func addVec(pos coord, vec coord) coord {
	pos.x += vec.x
	pos.y += vec.y
	return pos
}

func checkForMapEnd(guardMap [][]string, pos coord) bool {
	xLen := len(guardMap[0])
	yLen := len(guardMap)
	if pos.x == 0 || pos.y == 0 || pos.x == xLen-1 || pos.y == yLen-1 {
		return true
	}
	return false
}

// The puzzle tells us to turn right, but because the map was read in backwards,
// we also need to do everything the other way round.
func turnLeft(vec coord) coord {
	tmpVecX := vec.x
	if vec.x != 0 {
		vec.x = 0
	} else {
		vec.x = -vec.y
	}
	if vec.y != 0 {
		vec.y = 0
	} else {
		vec.y = tmpVecX
	}
	return vec
}

func countWalkedPositions(guardMap [][]string) int {
	passes := 0
	for y := 0; y < len(guardMap); y++ {
		for x := 0; x < len(guardMap[0]); x++ {
			if guardMap[y][x] == PassMarker {
				passes++
			}
		}
	}
	return passes
}

func printMap(guardMap [][]string) {
	for y := 0; y < len(guardMap); y++ {
		for x := 0; x < len(guardMap[0]); x++ {
			fmt.Printf(guardMap[y][x])
		}
		fmt.Printf("\n")
	}
}
