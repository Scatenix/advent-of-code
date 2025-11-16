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
Math is needed here! Everyone is talking about "Linear algebra".
Need to look into this.
The calculatePrice function is from some reddit user. I think this solution is doing linear algebra?
*/

const DayPart = "2024 Day 13 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-13/resources/puzzle-input"

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
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col []string) []string {
		col = append(col, line)
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	machines := inputToMachines(puzzleInput)

	tokens := 0
	for _, m := range machines {
		A, B := calculatePrice(m)
		tokens += A*3 + B
	}

	fmt.Printf(SolutionFormat, tokens)
}

func calculatePrice(m machine) (int, int) {
	A, B := 0, 0
	if (m.Ay*m.Bx - m.Ax*m.By) != 0 {
		B = ((m.Ay * m.PrizeX) - (m.Ax * m.PrizeY)) / (m.Ay*m.Bx - m.Ax*m.By)
	}
	if m.Ax != 0 {
		A = (m.PrizeX - (B * m.Bx)) / m.Ax
	}

	if ((m.Ax*A)+(m.Bx*B)) == m.PrizeX && ((m.Ay*A)+(m.By*B)) == m.PrizeY {
		return A, B
	}

	return 0, 0
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
			prizeX, _ := strconv.Atoi(strings.TrimRight(strings.Split(strings.Split(line, " ")[1], "=")[1], ","))
			prizeY, _ := strconv.Atoi(strings.Split(strings.Split(line, " ")[2], "=")[1])
			machines[len(machines)-1].PrizeX = 10000000000000 + prizeX
			machines[len(machines)-1].PrizeY = 10000000000000 + prizeY
		} else if len(line) <= 0 {
			machines = append(machines, machine{})
		}
	}
	return machines
}
