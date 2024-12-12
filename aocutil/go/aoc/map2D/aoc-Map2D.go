package map2D

// NOTE: AoC 2024 Day 12 Part 1&2 has lots of potentially useful stuff for future competitive programming problems

import (
	"atomicgo.dev/cursor"
	"fmt"
	"time"
)

type Coord struct {
	X int
	Y int
}

// GetDirectionalVectors get vector list to us in for loops when trying to run in every direction from a point in a 2d map
func GetDirectionalVectors() []Coord {
	return []Coord{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
}

// RotateRight taking in a direction(vector) and right clockwise
func RotateRight(dir Coord) Coord {
	tmpVecX := dir.X
	if dir.X != 0 {
		dir.X = 0
	} else {
		dir.X = -dir.Y
	}
	if dir.Y != 0 {
		dir.Y = 0
	} else {
		dir.Y = tmpVecX
	}
	return dir
}

// SearchCorner This will effectively create a 2x2 window from the current coordinate, going in the dir vector and then right.
// The window is then used to get all corners of any shape within a 2d array.
// Not a perfect function... use like I used it for AoC 2024 Day 12 Part 2 (memoization to avoid searching to many corners)
func SearchCorner(puzzleInput [][]rune, symbol rune, currentCoordinate Coord, directionVector Coord) int {
	pos1 := AddVector(currentCoordinate, directionVector)
	pos3 := AddVector(pos1, RotateRight(directionVector))
	pos2 := AddVector(currentCoordinate, RotateRight(directionVector))
	s1, s2, s3 := '.', '.', '.'
	if WithinBounds(puzzleInput, pos1) {
		s1 = puzzleInput[pos1.Y][pos1.X]
	}
	if WithinBounds(puzzleInput, pos2) {
		s2 = puzzleInput[pos2.Y][pos2.X]
	}
	if WithinBounds(puzzleInput, pos3) {
		s3 = puzzleInput[pos3.Y][pos3.X]
	}
	if (s1 == symbol && s2 == symbol && s3 != symbol) || (s1 != symbol && s2 != symbol) {
		return 1
	}
	return 0
}

func WithinBounds[T any](antennaMap [][]T, pos Coord) bool {
	if pos.X < 0 || pos.Y < 0 || pos.X >= len(antennaMap[0]) || pos.Y >= len(antennaMap) {
		return false
	}
	return true
}

func OutOfBounds[T any](antennaMap [][]T, pos Coord) bool {
	if pos.X < 0 || pos.Y < 0 || pos.X >= len(antennaMap[0]) || pos.Y >= len(antennaMap) {
		return true
	}
	return false
}

func Vector(pos1, pos2 Coord) Coord {
	return Coord{pos2.X - pos1.X, pos2.Y - pos1.Y}
}

func AddVector(pos, vec Coord) Coord {
	return Coord{pos.X + vec.X, pos.Y + vec.Y}
}

func SubVector(pos, vec Coord) Coord {
	return Coord{pos.X - vec.X, pos.Y - vec.Y}
}

func PrintMap[T any](m [][]T) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			fmt.Print("\033[0m%s", m[y][x])
		}
		fmt.Print("\n")
	}
}

func PrintMapForVisual(m [][]string, cursorVector Coord, show bool) {
	if show {
		cursor.HorizontalAbsolute(0)
		cursor.UpAndClear(128)
		for y := 0; y < len(m); y++ {
			for x := 0; x < len(m[0]); x++ {
				fmt.Printf("\033[0m%s", m[y][x])
			}
			fmt.Print("\n")
		}
		cursor.Move(cursorVector.X+1, len(m)-cursorVector.Y-1)
		cursor.Hide()
	}
}

func PrintMarkerForVisual(vec Coord, marker string, show bool, timeBetweenRenderInMicroSeconds int) {
	if show {
		time.Sleep(time.Duration(timeBetweenRenderInMicroSeconds) * time.Microsecond)
		cursor.Move(vec.Y-1, -vec.Y)
		fmt.Printf("\033[46m%s", marker)
		fmt.Print("\033[0m")
	}
}
