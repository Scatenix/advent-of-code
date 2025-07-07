# Advent of Code

**This is my Advent of Code repository!**

Personal disclaimer: Old git commits prior to June 2025 for this project are gone, since I had puzzle inputs and solutions in my repo before making it public!

## Structure

Puzzles are structured to be in their respective `Year/Day-#/Part-#/Technology` directory.

In each Day-# directory is a resource directory, containing the puzzle-input where your puzzle input **needs** to go.

Each Part-# directory contains a puzzle-text file containing the copied puzzle from the official site.

Note: Both the puzzle input and the solution are **not** stored in this repo as we are told by the creator to not do that!

## Puzzles I had troubles with

### Repeat in the future and see if I can solve it then

- 2024 Day 6 Part 2: I was just not able to finish it (Recursive Pathfinding problem)
- 2024 Day 10 Part 1&2: Topographic map. Did not know how to solve this (Recursive Pathfinding problem)

### Very hard imho, but finished. Sometimes with a few hints

- 2024 Day 11 Part 2: Performance optimization needed, using memoization
- 2024 Day 12 Part 1 and especially 2: Getting area, perimeter and number of sides for random shapes in 2D array

## Used Technologies

Make sure to use the `advent-of-code/createNewPuzzleTemplate.sh` script to create the file structure
with everything at the right place.
For more information use `sh advent-of-code/createNewPuzzleTemplate.sh help`
or go to the Templates section of this readme.

### Go

Run with `go run <GO_FILE>.go <PUZZLE_FILE>` or the automatically created IntelliJ run configuration.

Compile with `go build <GO_FILE>.go` and run as `GO_FILE <PUZZLE_FILE>`.

A lot of util functions, to get started with the puzzles faster were created at
`advent-of-code/aocutil/go/aoc/`.

### Rust

Run with `cargo run <RUST_FILE> <PUZZLE_FILE>` or just `cargo run <RUST_FILE>`.
Run in release mode for actual rust runtime speed and memory allocation `cargo run --release <RUST_FILE>`.

If no PUZZLE_FILE is supplied, a fallback to the automatically generated `path-to-solution/resources/puzzle-input` will be used.

### AWK

All awk scripts should be executable on all unix systems.
Run as normal AWK scripts would run `./<AWK_SCRIPT> <PUZZLE_FILE>`

### Languages I want to try in the future

- R
- Ocaml
- Haskel
- Zig

## Testing

not yet implemented

## Templates

I've written a bash script to automatically create a puzzle directory structure
for every technology that has a recipe bash script in `advent-of-code/templates/recipes`.

### How to create your own templates

The recipe is executed to copy or generate the code files to quickly start solving the AoC puzzle.
Place any resources to be used in a recipe somewhere in `advent-of-code/templates/`.
Recipe scripts start in the `advent-of-code/` directory.
Use the existing recipes as reference. They heavily utilize sed and even a bit of ed.

Template files and recipes for Go, Rust and AWK are already provided.

The structure described is created through this script plus what ever the recipe creates.
Be careful with your recipe scripts, as you can easily overwrite important stuff.
Using Git and regularly commiting is always advised!

To add your own templates, simply create a new recipe shell script in `advent-of-code/templates/recipes` as `TECHNOLOGY.sh`. The `createNewPuzzleTemplate` will simply execute the first argument as `<TECHNOLOGY>.sh` in the recipes directory.

### Usage of templates (User perspective)

To use this script, the AOC_HOME variable needs to be exported in the environment.
This is typically done in your shell's RC file.
The script will also tell you if it is missing and what to do.

E.g. for bash in `~/.bashrc` add this somewhere in the file `export AOC_HOME="<PATH_TO_AOC_PROJECT>"`.

Can also be done on every shell session where the script is run:
`export AOC_HOME=$(pwd)`

---

**The script can be executed in two ways:**

Create for any day
`sh advent-of-code/createNewPuzzleTemplate.sh <TECHNOLOGY> <YEAR> <DAY> <PART#>`

Create for today 
`sh advent-of-code/createNewPuzzleTemplate.sh <TECHNOLOGY> today`

For help information use
`sh advent-of-code/createNewPuzzleTemplate.sh help`

---

TECHNOLOGY: Currently `go`, `rust` or `awk`, or just anything that is in the recipe directory.
            E.g. to add a java template, create a java.sh in there and you are set to use "java" as technology!
YEAR: Year of desired AoC puzzle. (e.g. 2024)
DAY: Day of desired AoC puzzle. (e.g. 24)
PART#: Part number of desired AoC puzzle. (1 or 2)
