package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocmath "advent-of-code/aocutil/go/aoc/math"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const DayPart = "Day 1 - Part 1"
const SolutionFormat = "The sum of the distances is: %d\n"

const Left = 0
const Right = 1

type distance struct {
	leftDistances []int
	rightDistances []int
}

func main() {
	puzzleFile := aocutil.AocSetup(DayPart)

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

	sort.Ints(distances.leftDistances)
	sort.Ints(distances.rightDistances)

	sumDistance := 0;
	for i := range distances.leftDistances {
		sumDistance += aocmath.Abs(distances.rightDistances[i] - distances.leftDistances[i])
	}

	fmt.Printf(SolutionFormat, sumDistance)
}
