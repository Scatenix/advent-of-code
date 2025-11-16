use aoc;

#[macro_use(defer)]
extern crate scopeguard;

const FALLBACK_PUZZLE_INPUT_PATH: &str = "2019/Day-1/resources/puzzle-input";
const DAY_PART: &str = "2019 Day 1 - Part 1";
const SOLUTION_FORMAT: &str = ">>> The solution is: ";

// Usage: app <PATH_TO_PUZZLE_FILE>
fn main() {
    defer!{aoc::perf::time_tracker(std::time::Instant::now(), "main")}
    let file = aoc::util::aoc_setup(DAY_PART, FALLBACK_PUZZLE_INPUT_PATH);

    let reading_handler = |line: String, mut col: Vec<isize>| -> Vec<isize> {
        col.push(line.parse::<isize>().unwrap());
        return col
    };
w
    let module_masses = aoc::io::read_puzzle_file(file, reading_handler);
    let total_fuel: isize = module_masses.iter().map(|mass| mass / 3 - 2).sum();
    println!("{SOLUTION_FORMAT}{}", total_fuel);
}
