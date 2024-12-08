package map2D

import (
	"atomicgo.dev/cursor"
	"fmt"
	"time"
)

type Coord struct {
	X int
	Y int
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
	for y := 0; y<len(m); y++ {
		for x := 0; x<len(m[0]); x++ {
			fmt.Print("\033[0m%s",m[y][x])
		}
		fmt.Print("\n")
	}
}

func PrintMapForVisual(m [][]string, cursorVector Coord, show bool) {
	if show {
		cursor.HorizontalAbsolute(0)
		cursor.UpAndClear(128)
		for y := 0; y<len(m); y++ {
			for x := 0; x<len(m[0]); x++ {
				fmt.Printf("\033[0m%s",m[y][x])
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
		fmt.Print("\033[0m", )
	}
}