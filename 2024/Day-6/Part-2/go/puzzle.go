package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strings"
	"time"
)

const DayPart = "2024 Day 6 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"

const Guard = "^"
const Wall = "#"
const Walkway = "."

type coord struct {
	x int
	y int
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
    defer aocperf.TimeTracker(time.Now(), "Main")
    defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col [][]string) [][]string {
		col = append(col, strings.Split(line, ""))
		return col
	}

	guardMap, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	infinitumObstacle := 0
	pos := locateStart(guardMap)
	walkVec := coord{x: 0, y:-1}
	dc := aocslice.DeepCopy2D(guardMap)

	visited := map[coord]coord{{pos.x, pos.y}: {walkVec.x, walkVec.y}}
	_, visited = walk(dc, coord{pos.x, pos.y}, walkVec, visited, 0)

	for place := range visited {
		newMap := aocslice.DeepCopy2D(guardMap)
		newMap[place.y][place.x] = Wall

		success, _ := walk(newMap, coord{pos.x, pos.y}, walkVec, map[coord]coord{{0, 0}: {0, 0}}, 0)
		if !success {
			infinitumObstacle++
		}

		newMap[place.y][place.x] = Walkway
	}
	fmt.Printf(SolutionFormat, infinitumObstacle)
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

func walk(guardMap [][]string, pos coord, vec coord, visited map[coord]coord, depth int) (bool, map[coord]coord) {
	nextY, nextX := pos.y+vec.y, pos.x+vec.x

	if depth > 20000 {
		return false, visited
	}
	if checkForMapEnd(guardMap, pos) {
		return true, visited
	}
	if guardMap[nextY][nextX] == Wall {
		vec = turnLeft(vec)
	} else {
		pos.y, pos.x = nextY, nextX
		//visited = append(visited, [][]int{{nextY, nextX}}...)
		visited[pos] = coord{vec.x, vec.y}
	}

	depth++
	return walk(guardMap, pos, vec, visited, depth)
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