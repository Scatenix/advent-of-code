# Advent of Code

**This is my Advent of Code repository!**

Puzzles are structured to be in their respective `Year/Day-#/Part-#/Technology`.

Each Day-# directory is a resource directory, containing the puzzle-file.

Each Part-# directory contains a puzzle-text file containing
the copied puzzle from the official site including my personal solution.

## Puzzles I sucked at

### Repeat in the future and see if I can solve it then

- 2024 Day 6 Part 2: I was just not able to finish it (Recursive Pathfinding problem)
- 2024 Day 10 Part 1&2: Topographic map. Had no Idea how to do this (Recursive Pathfinding problem)

### Very hard imho, but finished. Sometimes with a few hints

- 2024 Day 11 Part 2: Performance optimization needed, using memoization
- 2024 Day 12 Part 1 and especially 2: Getting area, perimeter and number of sides for random shapes in 2D array

## Conclusions and Takeaway

Sometimes for problems I learned something valuable, I will write it down as comment
in the file. They start with "/* Conclusion" and end with the comment "*/"
Or "# Conclusion" and end with the last consecutive line starting with an "#"

They can be compibled to a more exhaustive list of conclusions and takeaways

## Used Technologies

### Go

Run with `go run <GO_FILE>.go <PUZZLE_FILE>`

Compile with `go build <GO_FILE>.go`(and run as `GO_FILE <PUZZLE_FILE>`)

A lot of util functions, to get started with the puzzles faster, where created at
`advent-of-code/aocutil/go/aoc/<UTIL_FILES>`


### AWK

All awk scripts are executable on unix
Run as normal AWK scripts would run `./<AWK_SCRIPT> <PUZZLE_FILE>`


## Testing

not yet implemented

## Templates

I've written a shell script to automatically create a puzzle directory structure
for every technology that has a recipe shell script in `advent-of-code/templates/recipes`.

Template files and recipes for Go and AWK are already provided.

The structure described is created through this script plus what ever the recipe creates.
Be careful with your recipe scripts, as you can easily overwrite important stuff.
Using Git and regularly commiting is always advised!

**The script can be executed in two ways:**

Create for any day
`sh advent-of-code/createNewPuzzleTemplate.sh <TECHNOLOGY(i.e. go)> <YEAR> <DAY> <PART#>`

---

Create for today 
`sh advent-of-code/createNewPuzzleTemplate.sh <TECHNOLOGY(i.e. go)> today`

---

For help information use
`sh advent-of-code/createNewPuzzleTemplate.sh help`

---

TECHNOLOGY: Currently `go` or `awk`, or just anything that is in the recipe directory.
            To add a java template, create a java.sh in there and you are set to use "java" as technology!
YEAR: Year of desired AoC puzzle.
DAY: Day of desired AoC puzzle.
PART#: Part number of desired AoC puzzle.

---

To use this script, the AOC_HOME variable needs to be exported in the environment.
This is typically done in your shell's RC file.

E.g. for bash in `~/.bashrc`
Add this somewhere in the file `export AOC_HOME="<PATH_TO_AOC_PROJECT>"`

Can also be done on every shell session where the script is run:
`export AOC_HOME=$(pwd)`
