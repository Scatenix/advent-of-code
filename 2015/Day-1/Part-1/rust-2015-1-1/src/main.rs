use aoc;

const FALLBACK_PUZZLE_INPUT_PATH: &str = "2015/Day-1/resources/puzzle-input";
const DAY_PART: &str = "2015 Day 1 - Part 1";
const SOLUTION_FORMAT: &str = ">>> The solution is: ";

// Usage: app <PATH_TO_PUZZLE_FILE>
fn main() {
    let file = aoc::util::aoc_setup(DAY_PART, FALLBACK_PUZZLE_INPUT_PATH);

    let reading_hanlder = |line: String, mut col: String| -> String {
        col.push_str(&line);
        col.push('\n');
        return col
    };

    let puzzle_input = aoc::io::read_puzzle_file(file, reading_hanlder);

    println!("{SOLUTION_FORMAT}{}", puzzle_input.len());
}
