package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DayPart = "2024 Day 11 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"
const FallbackPuzzleInputPath = "/home/sca/Programming/advent-of-code/2024/Day-11/resources/puzzle-input"
const Blinks = 75

/* Conclusion
2024 Day 11 - Part 2
You can see the evolution of my attempt to get a good runtime for 75 blinks
1. Started with what is still on part 1. A slice which gets expanded til the end. VERY high memory consumption
2. Tried to get results faster using multithreading. Horrible code. Still VERY high mem consumption. And not even fast
3. Replaced slice with linked list. Better performance but the same mem consumption
4. Splitting tasks up for better mem consumption. Basically create smaller sub-lists for the GC to collect old ones
5. Changing to a recursive function. Naturally good mem consumption and even faster than the linked list approach
6. Current: Introducing memoization to compute every number only once and remember the stone count per number. Superfast, very good mem consumption

Takeaway: 1. If recursion is possible, it might very well be the fastest way, as there is no overhead for handling the datastructures
          2. memoization is king when doing the exact same calculation multiple times
*/

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart, FallbackPuzzleInputPath)

	puzzleLineHandler := func(line string, col []int) []int {
		col = append(col, aocslice.Atoi(strings.Split(line, " "))...)
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	stones := 0

	for _, v := range puzzleInput {
		stones += blink(v, 0, make(map[[2]int]int))
	}

	fmt.Printf(SolutionFormat, stones)
}

func blink(val, blinked int, memoizeMap map[[2]int]int) int {
	if memoizeMap[[2]int{val, blinked}] != 0 {
		return memoizeMap[[2]int{val, blinked}]
	}
	if blinked == Blinks {
		return 1
	}
	strVal := strconv.Itoa(val)
	if val == 0 {
		return blink(1, blinked+1, memoizeMap)
	} else if len(strVal)%2 == 0 {
		firstHalf, _ := strconv.Atoi(strVal[0 : len(strVal)/2])
		secondHalf, _ := strconv.Atoi(strVal[len(strVal)/2:])
		tmp1 := blink(firstHalf, blinked+1, memoizeMap)
		tmp2 := blink(secondHalf, blinked+1, memoizeMap)
		memoizeMap[[2]int{val, blinked}] += tmp1 + tmp2
		return tmp1 + tmp2
	} else {
		return blink(val*2024, blinked+1, memoizeMap)
	}
}
