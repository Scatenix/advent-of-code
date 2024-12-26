use aoc;

#[macro_use(defer)]
extern crate scopeguard;

const FALLBACK_PUZZLE_INPUT_PATH: &str = "2019/Day-2/resources/puzzle-input";
const DAY_PART: &str = "2019 Day 2 - Part 1";
const SOLUTION_FORMAT: &str = ">>> The solution is: ";

// Usage: app <PATH_TO_PUZZLE_FILE>
fn main() {
    defer!{aoc::perf::time_tracker(std::time::Instant::now(), "main")}
    let file = aoc::util::aoc_setup(DAY_PART, FALLBACK_PUZZLE_INPUT_PATH);

    let reading_handler = |line: String, mut col: Vec<isize>| -> Vec<isize> {
        col.extend(line.split(',')
            .map(|s| s.parse::<isize>().unwrap()));
        return col
    };

    let mut puzzle_input = aoc::io::read_puzzle_file(file, reading_handler);

    puzzle_input[1] = 12;
    puzzle_input[2] = 2;

    let mut addr = 0;
    while addr != -1 {
        addr = aoc::intcode::intcode_interpreter(&mut puzzle_input, addr as usize);
    }

    println!("{SOLUTION_FORMAT}{}", puzzle_input[0]);
}
