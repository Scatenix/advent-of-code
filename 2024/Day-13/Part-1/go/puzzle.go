package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

/* Conclusion
 */

const DayPart = "2024 Day 13 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"

type machine struct {
	Ax     int
	Ay     int
	Bx     int
	By     int
	PrizeX int
	PrizeY int
}

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col []string) []string {
		col = append(col, line)
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	machines := inputToMachines(puzzleInput)

	tokens := 0
	for _, machine := range machines {
		maxBxPresses := machine.PrizeX / machine.Bx
		maxByPresses := machine.PrizeY / machine.By

		BxMap := make(map[int]int)
		ByMap := make(map[int]int)

		for i := 1; i < maxBxPresses+1; i++ {
			BxMap[machine.Bx*i] = i
		}
		for i := 1; i < maxByPresses+1; i++ {
			ByMap[machine.By*i] = i
		}

		xPrize := machine.PrizeX
		yPrize := machine.PrizeY
		aCount := 0
		for xPrize >= 0 || yPrize >= 0 {
			//if BxMap[xPrize] > 0 && BxMap[xPrize] == ByMap[yPrize] && aCount+BxMap[xPrize] <= 100 {
			if BxMap[xPrize] > 0 && BxMap[xPrize] == ByMap[yPrize] {
				tokens += aCount*3 + BxMap[xPrize]
			} else if xPrize == 0 && yPrize == 0 {
				tokens += aCount * 3
			}
			xPrize -= machine.Ax
			yPrize -= machine.Ay
			aCount++
		}
	}

	fmt.Printf(SolutionFormat, tokens)
}

func inputToMachines(puzzleInput []string) []machine {
	machines := make([]machine, 0)
	machines = append(machines, machine{})
	for _, line := range puzzleInput {
		if strings.Contains(line, "Button A") {
			machines[len(machines)-1].Ax, _ = strconv.Atoi(strings.TrimRight(strings.Split(strings.Split(line, " ")[2], "+")[1], ","))
			machines[len(machines)-1].Ay, _ = strconv.Atoi(strings.Split(strings.Split(line, " ")[3], "+")[1])
		} else if strings.Contains(line, "Button B") {
			machines[len(machines)-1].Bx, _ = strconv.Atoi(strings.TrimRight(strings.Split(strings.Split(line, " ")[2], "+")[1], ","))
			machines[len(machines)-1].By, _ = strconv.Atoi(strings.Split(strings.Split(line, " ")[3], "+")[1])
		} else if strings.Contains(line, "Prize") {
			machines[len(machines)-1].PrizeX, _ = strconv.Atoi(strings.TrimRight(strings.Split(strings.Split(line, " ")[1], "=")[1], ","))
			machines[len(machines)-1].PrizeY, _ = strconv.Atoi(strings.Split(strings.Split(line, " ")[2], "=")[1])
		} else if len(line) <= 0 {
			machines = append(machines, machine{})
		}
	}
	return machines
}
