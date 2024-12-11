use aoc;

const FALLBACK_PUZZLE_INPUT_PATH: &str = "2015/Day-1/resources/puzzle-input";
const DAY_PART: &str = "2015 Day 1 - Part 1";
const SOLUTION_FORMAT: &str = ">>> The solution is: ";

// Usage: app <PATH_TO_PUZZLE_FILE>
fn main() {
    let file = aoc::util::aoc_setup(DAY_PART, FALLBACK_PUZZLE_INPUT_PATH);

    let reading_handler = |line: String, mut floor: isize| -> isize {
        for ch in line.chars() {
            if ch == '(' {
                floor += 1;
            } else {
                floor -= 1;
            }
        }
        return floor
    };

    let puzzle_input = aoc::io::read_puzzle_file(file, reading_handler);

    println!("{SOLUTION_FORMAT}{}", puzzle_input);
}
