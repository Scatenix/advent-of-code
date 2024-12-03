package main

import (
	aocio "advent-of-code/aocutil/go/aoc/io"
	aocutil "advent-of-code/aocutil/go/aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

const DayPart = "Day 3 - Part 2"
const SolutionFormat = "The sum of all conditional multiplications is: %d\n"

func main() {
	puzzleFile := aocutil.AocSetup(DayPart)

	puzzleLineHandler := func(line string, ret string) string {
		ret += line
		return ret
	}

	corruptMem, err := aocio.ReadPuzzleFile[string](puzzleFile, puzzleLineHandler)
	aocutil.Check(err)

	sumOfAllMultiplications := 0;
	rgxMulCommand := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|don't()|do())`)
	// the param n means how many matches should be made at max. -1 means infinite matches are allowed.
	mulCommands := rgxMulCommand.FindAllString(corruptMem, -1)

	rgxMulNumbers := regexp.MustCompile(`\d{1,3}`)
	rgxDo := regexp.MustCompile(`do()`)
	rgxDont := regexp.MustCompile(`don't()`)
	mulEnabled := true
	for _, mul := range mulCommands {
		if (rgxDont.MatchString(mul)) {
			mulEnabled = false
		} else if (rgxDo.MatchString(mul)) {
			mulEnabled = true
		} else if (mulEnabled) {
			mulNumbers := rgxMulNumbers.FindAllString(mul, -1)
			n1, err := strconv.Atoi(mulNumbers[0])
			aocutil.Check(err)
			n2, err := strconv.Atoi(mulNumbers[1])
			aocutil.Check(err)

			sumOfAllMultiplications += n1 * n2
		}
	}

	fmt.Printf(SolutionFormat, sumOfAllMultiplications)
}
