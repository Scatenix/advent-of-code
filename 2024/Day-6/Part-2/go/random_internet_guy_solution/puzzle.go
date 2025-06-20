package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// https://adventofcode.com/2024/day/6

var chdir map[rune]rune = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

var dir map[rune][2]int = map[rune][2]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

type Field [][]rune

func read(file string) (f Field, ry, rx int) {
	in, _ := os.ReadFile(file)
	// f := Field{}
	for y, l := range strings.Split(strings.TrimSpace(string(in)), "\n") {
		f = append(f, []rune(l))
		if x := strings.IndexRune(l, '^'); x != -1 {
			ry, rx = y, x
		}
	}
	return f, ry, rx
}

func (field Field) walk(y, x int, d rune) (visited map[[2]int]rune, cycle bool) {
	visited = map[[2]int]rune{{y, x}: d}
	for range len(field) * len(field[0]) {
		ny, nx := y+dir[d][0], x+dir[d][1]
		// bounds check
		if ny < 0 || ny >= len(field) || nx < 0 || nx >= len(field[0]) {
			break
		}

		// cycle check
		if visited[[2]int{ny, nx}] == d {
			return visited, true
		}

		if field[ny][nx] == '#' {
			d = chdir[d]
			continue
		}

		y, x = ny, nx
		visited[[2]int{y, x}] = d
	}
	return
}

func solve(f string) (p1, p2 int) {
	field, y, x := read(f)
	d := '^'
	obstructions := [][2]int{}

	visited, _ := field.walk(y, x, d)
	for check := range visited {
		field[check[0]][check[1]] = '#'
		if _, cycle := field.walk(y, x, d); cycle {
			p2++
			obstructions = append(obstructions, [2]int{check[0], check[1]})
		}
		field[check[0]][check[1]] = '.'
	}

	return len(visited), p2
}

func main() {
	start := time.Now()
	fmt.Println("Day 06: Guard Gallivant")
	// p1, p2 := solve("input_pub.txt")
	p1, p2 := solve("resources/puzzle-input")
	fmt.Println("\tPart One:", p1) // 4647
	fmt.Println("\tPart Two:", p2) // 1723
	fmt.Printf("Done in %.3f seconds \n", time.Since(start).Seconds())
}