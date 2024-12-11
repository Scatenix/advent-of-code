package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocperf "advent-of-code/aocutil/go/aoc/perf"
	aocslice "advent-of-code/aocutil/go/aoc/slice"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DayPart = "2024 Day 11 - Part 2"
const SolutionFormat = ">>> The solution is: %d\n"

// Usage: app <PATH_TO_PUZZLE_FILE>
func main() {
	defer aocperf.TimeTracker(time.Now(), "Main")
	defer aocperf.PrintMemUsage(aocperf.KB, "Main")
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, col []int) []int {
		col = append(col, aocslice.Atoi(strings.Split(line, " "))...)
		return col
	}

	puzzleInput, err := aocio.ReadPuzzleFile(puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	stones := 0

	for _, v := range puzzleInput {
		l := list.New()
		l.PushFront(v)

		blink(l, 40)
		stones += l.Len()
		fmt.Printf(SolutionFormat, stones)

		s := linkedListToSlice(l)
		print(s)
		for x, v := range s {
			l2 := list.New()
			l2.PushFront(v)

			blink(l2, 35)
			stones += l2.Len() - 1
			//fmt.Printf(SolutionFormat, stones)
			println("At step of second part ", x, " of ", len(s))
		}
	}

	//fmt.Printf(SolutionFormat, puzzleInput.Len())
	fmt.Printf("--------------")
	fmt.Printf(SolutionFormat, stones)
}

func blink(l *list.List, times int) {
	for i := 0; i < times; i++ {
		for e := l.Front(); e != nil; e = e.Next() {
			strVal := strconv.Itoa(e.Value.(int))
			if e.Value == 0 {
				e.Value = 1
			} else if len(strVal)%2 == 0 {
				firstHalf, _ := strconv.Atoi(strVal[0 : len(strVal)/2])
				secondHalf, _ := strconv.Atoi(strVal[len(strVal)/2:])
				e.Value = firstHalf
				l.InsertAfter(secondHalf, e)
				e = e.Next()
			} else {
				e.Value = e.Value.(int) * 2024
			}
		}
		//println("finished ", i, " of ", times)
	}
}

func linkedListToSlice(ll *list.List) []int {
	s := make([]int, 0)
	for e := ll.Front(); e != nil; e = e.Next() {
		s = append(s, e.Value.(int))
	}
	return s
}
