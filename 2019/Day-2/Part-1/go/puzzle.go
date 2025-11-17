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

const DayPart = "2019 Day 2 - Part 1"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2019/Day-2/resources/puzzle-input"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col []int) []int {
		strList := strings.Split(line, ",")
		var intList []int

		for _, v := range strList {
			i, _ := strconv.Atoi(v)
			intList = append(intList, i)
		}

		return intList
	}

	program, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	// Prepare program for puzzle, as instructed
	program[1] = 12
	program[2] = 2

	pptr := 0
	for program[pptr] != PStop {
		pExec(pptr, program)

		pptr += 4
		if pptr >= len(program) {
			println("No stop code 99 reached in the program stack.")
			break
		}
	}

	fmt.Printf(SolutionFormat, program[0])
}

const PStop int = 99
const ADD int = 1
const MUL int = 2

func pExec(pptr int, program []int) {
	op := program[pptr]
	src1 := program[pptr+1]
	src2 := program[pptr+2]
	dst := program[pptr+3]

	if op == ADD {
		program[dst] = program[src1] + program[src2]
	}
	if op == MUL {
		program[dst] = program[src1] * program[src2]
	}
}
