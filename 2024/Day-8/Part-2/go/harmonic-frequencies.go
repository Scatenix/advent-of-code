package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	"advent-of-code/aocutil/go/aoc/map2D"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"time"
)

const DayPart = "2024 Day 8 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-8/resources/puzzle-input"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col [][]rune) [][]rune {
		col = append(col, []rune(line))
		return col
	}

	antennaMap, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	freqMap := getFreqMap(antennaMap)

	antinodes := calculateAntinodesCount(antennaMap, freqMap)
	printAntennaMap(antennaMap, antinodes)

	fmt.Printf(SolutionFormat, len(antinodes))
}

func getFreqMap(antennaMap [][]rune) map[rune][]map2D.Coord {
	freqMap := make(map[rune][]map2D.Coord)
	for y := 0; y < len(antennaMap); y++ {
		for x := 0; x < len(antennaMap); x++ {
			if antennaMap[y][x] != '.' {
				freqMap[antennaMap[y][x]] = append(freqMap[antennaMap[y][x]], map2D.Coord{x, y})
			}
		}
	}
	return freqMap
}

func calculateAntinodesCount(antennaMap [][]rune, freqMap map[rune][]map2D.Coord) map[map2D.Coord]bool {
	antinodes := make(map[map2D.Coord]bool)
	for _, freq := range freqMap {
		for _, pos1 := range freq {
			for _, pos2 := range freq {
				if pos1 != pos2 {
					vec := map2D.Vector(pos1, pos2)
					antinodes[pos1] = true
					antinodes[pos2] = true

					antiPos := pos1
					for map2D.WithinBounds(antennaMap, antiPos) {
						antinodes[antiPos] = true
						antiPos = map2D.SubVector(antiPos, vec)
					}
					antiPos = pos2
					for map2D.WithinBounds(antennaMap, antiPos) {
						antinodes[antiPos] = true
						antiPos = map2D.AddVector(antiPos, vec)
					}
				}
			}
		}
	}
	return antinodes
}

func printAntennaMap(antennaMap [][]rune, antinodes map[map2D.Coord]bool) {
	for y := 0; y < len(antennaMap); y++ {
		for x := 0; x < len(antennaMap[0]); x++ {
			if antennaMap[y][x] == '.' && antinodes[map2D.Coord{x, y}] {
				fmt.Printf("\033[32m#")
			} else if antinodes[map2D.Coord{x, y}] {
				fmt.Printf("\033[32m%s", string(antennaMap[y][x]))
			} else {
				fmt.Printf("\033[0m%s", string(antennaMap[y][x]))
			}
		}
		fmt.Print("\n")
	}
}
