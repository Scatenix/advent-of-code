package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
)

const DayPart = "Day 1 - Part 2"
const SolutionFormat = "The similarity score is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-4/resources/puzzle-input"

const Left = 0
const Right = 1

type distance struct {
	leftDistances  []int
	rightDistances []int
}

// Usage: go-app <PATH_TO_PUZZLE_FILE>
func main() {
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, ret distance) distance {
		distances := strings.Fields(line)
		leftDistance, err := strconv.Atoi(distances[Left])
		aocutil.Check(err)
		rightDistance, err := strconv.Atoi(distances[Right])
		aocutil.Check(err)
		ret.leftDistances = append(ret.leftDistances, leftDistance)
		ret.rightDistances = append(ret.rightDistances, rightDistance)
		return ret
	}

	distances, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	similarityScore := 0
	for _, vi := range distances.leftDistances {
		count := 0
		for _, vj := range distances.rightDistances {
			if vi == vj {
				count++
			}
		}
		similarityScore += vi * count
	}

	fmt.Printf(SolutionFormat, similarityScore)
}
