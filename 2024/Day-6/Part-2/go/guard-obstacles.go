package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"atomicgo.dev/cursor" // Doc: https://github.com/atomicgo/cursor
	"fmt"
	"strings"
	"time"
)

// TODO: This go file contains some comments, starting with "NOTE:" with suggestions to improve the code for performance gains

// visual only works in a terminal which can fit the whole map at once. Else there are massive bugs with it
const visual = false
const timeBetweenRenderInMicroSeconds = 100
const showFromRun = 2

var VisualRun = 1

const DayPart = "2024 Day 6 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-6/resources/puzzle-input"

const Guard = "^"
const Wall = "#"
const Obstacle = "O"
const Walkway = "."
const PassMarker = "X"

type coord struct {
	x int
	y int
}

type puzzleInput struct {
	guardMap   [][]string
	guardStart coord
	lineNum    int
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, puzzleInput puzzleInput) puzzleInput {
		puzzleInput.lineNum++
		puzzleInput.guardMap = append(puzzleInput.guardMap, strings.Split(line, ""))
		if strings.Contains(line, Guard) {
			puzzleInput.guardStart = coord{strings.Index(line, Guard), puzzleInput.lineNum}
		}
		return puzzleInput
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	infinitumObstacle := 0
	pos := puzzleInput.guardStart
	walkVec := coord{x: 0, y: -1}
	dc := aocslice.DeepCopy2D(puzzleInput.guardMap)

	printMap(puzzleInput.guardMap, pos)
	visited := map[coord]coord{{pos.x, pos.y}: {walkVec.x, walkVec.y}}
	_, visited = walk(dc, pos, walkVec, visited)

	// NOTE: This could be a good candidate for multi-threading
	for place := range visited {
		puzzleInput.guardMap[place.y][place.x] = Obstacle

		printMap(puzzleInput.guardMap, pos)

		success, _ := walk(puzzleInput.guardMap, coord{pos.x, pos.y}, walkVec, map[coord]coord{{0, 0}: {0, 0}})
		if !success {
			infinitumObstacle++
		}

		puzzleInput.guardMap[place.y][place.x] = Walkway
	}
	fmt.Printf(SolutionFormat, infinitumObstacle)
}

// NOTE: more efficent here would be to calculate jump tables at the start of the program (precalculating on which pos I will end up)
// It could be something like a map, containing the position+walkVec after turning left because of an obstacle and with the value of the jumping destination
// Only tricky part would be to not use the jump table at the path where the Obstacle resides (probably not too hard?)
// With that I could save n-1 calls between each obstacle
func walk(guardMap [][]string, pos coord, vec coord, visited map[coord]coord) (bool, map[coord]coord) {
	nextY, nextX := pos.y+vec.y, pos.x+vec.x

	// infinite loop exit condition (if we are at a position we already visited, facing the same direction, we are in an infinite loop)
	if visited[coord{nextX, nextY}] == vec {
		return false, visited
	}

	// success exit condition
	if pos.x == 0 || pos.y == 0 || pos.x == len(guardMap[0])-1 || pos.y == len(guardMap)-1 {
		return true, visited
	}
	if guardMap[nextY][nextX] == Wall || guardMap[nextY][nextX] == Obstacle {
		vec = turnLeft(vec)
	} else {
		pos.y, pos.x = nextY, nextX
		visited[pos] = coord{vec.x, vec.y}
	}

	printPassMarker(vec)

	return walk(guardMap, pos, vec, visited)
}

// The puzzle tells us to turn right, but because the map was read in backwards,
// we also need to do everything the other way round.
// NOTE: while it is certainly I programatic way, something like a switch case is just better for this case
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

func printMap(guardMap [][]string, pos coord) {
	if visual {
		VisualRun++
		if VisualRun >= showFromRun {
			cursor.HorizontalAbsolute(0)
			cursor.UpAndClear(128)
			for y := 0; y < len(guardMap); y++ {
				for x := 0; x < len(guardMap[0]); x++ {
					if guardMap[y][x] == Obstacle {
						fmt.Print("\033[41m", Obstacle)
					} else {
						fmt.Print("\033[0m", guardMap[y][x])
					}
				}
				fmt.Print("\n")
			}
			cursor.Move(pos.x+1, len(guardMap)-pos.y-1)
			cursor.Hide()
		}
	}
}

func printPassMarker(vec coord) {
	if visual {
		if VisualRun >= showFromRun {
			time.Sleep(timeBetweenRenderInMicroSeconds * time.Microsecond)
			cursor.Move(vec.x-1, -vec.y)
			fmt.Print("\033[46m", PassMarker)
			fmt.Print("\033[0m")
		}
	}
}
