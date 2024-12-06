package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"atomicgo.dev/cursor"
	//"github.com/gookit/color"
	"fmt"
	"strings"
)

const DayPart = "2024 Day 6 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"

const Guard = "^"
const PassMarker = "X"
const Wall = "#"
const Wallo = "O"
const Walkway = "."

const ColorGreen = "\033[32m"
const ColorReset = "\033[0m"

type coord struct {
	x int
	y int
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
    //defer aocperf.TimeTracker(time.Now(), "Main")
    //defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col [][]string) [][]string {
		col = append(col, strings.Split(line, ""))
		return col
	}

	guardMap, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	infinitumObstacle := 0
	// We start by walking down, because the map was read backwards

	// first walk to see where we want to put the barriets (we only put them were the guard actually wants to walk
	pos := locateStart(guardMap)
	walkVec := coord{x: 0, y:-1}
	dc := aocslice.DeepCopy2D(guardMap)
	walk(dc, pos, walkVec, 0)
	walkedPositions := getWalkedPositions(dc)

	for y, row := range guardMap {
		for x := range row {
			pos := locateStart(guardMap)
			walkVec = coord{x: 0, y:-1}
			newMap := aocslice.DeepCopy2D(guardMap)

			//fmt.Printf("Currently at: %d %d\n", x, y)

			if newMap[y][x] == Wall {
				continue
			}
			newMap[y][x] = Wallo

			printMap(newMap)
			cursor.Move(pos.x+1, len(guardMap)-pos.y-1)
			cursor.Hide()

			if !walk(newMap, pos, walkVec, 0) {
				infinitumObstacle++
			}
			fmt.Print("\033[0m")
			cursor.Move(-130, 0)
			newMap[y][x] = Walkway
		}
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

// walk the map and leave an x
func walk(guardMap [][]string, pos coord, vec coord, depth int) bool {
	if depth > 20000 {
		return false
	}

	//time.Sleep(1 * time.Millisecond)
	//time.Sleep(100 * time.Microsecond)
	//printMap(guardMap)
	printPassMarker(vec)

	if checkForMapEnd(guardMap, pos) {
		guardMap[pos.y][pos.x] = PassMarker
		return true
	}

	if guardMap[pos.y+vec.y][pos.x+vec.x] == Wall || guardMap[pos.y+vec.y][pos.x+vec.x] == Wallo {
		vec = turnLeft(vec)
	}

	guardMap[pos.y][pos.x] = PassMarker
	pos = addVec(pos, vec)
	return walk(guardMap, pos, vec, depth+1)
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

func getWalkedPositions(guardMap [][]string) []coord {
	walkedPos := make([]coord, 0)
	for y := 0; y<len(guardMap); y++ {
		for x := 0; x<len(guardMap[0]); x++ {
			if guardMap[y][x] == PassMarker {
				walkedPos = append(walkedPos, coord{x: x, y: y})
			}
		}
	}
	return walkedPos
}

func printMap(guardMap [][]string) {
	cursor.UpAndClear(129)
	for y := 0; y<len(guardMap); y++ {
		for x := 0; x<len(guardMap[0]); x++ {
			if guardMap[y][x] == PassMarker {
				//fmt.Print("\033[33m", PassMarker)
				fmt.Print("\033[46m%s\033[0m", PassMarker)
			} else {
				fmt.Print(guardMap[y][x])
				//fmt.Print("\033[0m", guardMap[y][x])
			}
		}
		fmt.Print("\n")
	}
}

func printPassMarker(vec coord) {
	cursor.Move(vec.x-1, -vec.y)
	//fmt.Print("ESC[{line};{column}H", PassMarker)
	//fmt.Print("\033[%d;%dH\033[46m", 50, 50, PassMarker)
	fmt.Print("\033[46m", PassMarker)
}