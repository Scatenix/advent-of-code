use aoc;

#[macro_use(defer)]
extern crate scopeguard;

const FALLBACK_PUZZLE_INPUT_PATH: &str = "2019/Day-1/resources/puzzle-input";
const DAY_PART: &str = "2019 Day 1 - Part 2";
const SOLUTION_FORMAT: &str = ">>> The solution is: ";

// Usage: app <PATH_TO_PUZZLE_FILE>
fn main() {
    defer!{aoc::perf::time_tracker(std::time::Instant::now(), "main")}
    let file = aoc::util::aoc_setup(DAY_PART, FALLBACK_PUZZLE_INPUT_PATH);

    let reading_handler = |line: String, mut col: Vec<isize>| -> Vec<isize> {
        col.push(line.parse::<isize>().unwrap());
        return col
    };
    let module_masses = aoc::io::read_puzzle_file(file, reading_handler);

    let mut total_fuel: isize = 0;

    for module in module_masses {
        let mut module_fuel: isize = module / 3 - 2;
        let mut add_fuel = module_fuel;
        loop {
            add_fuel = add_fuel / 3 - 2;
            if add_fuel <= 0 {
                break;
            }
            module_fuel += add_fuel;
        }
        total_fuel += module_fuel
    }

    println!("{SOLUTION_FORMAT}{}", total_fuel);
}
