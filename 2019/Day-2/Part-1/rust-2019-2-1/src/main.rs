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
    let mut prog_stack = aoc::intcode::prog_stack_reader(file);

    prog_stack[1] = 12;
    prog_stack[2] = 2;

    let mut inst_pointer = 0;
    while inst_pointer != -1 {
        inst_pointer = aoc::intcode::intcode_interpreter_v1(&mut prog_stack, inst_pointer as usize);
    }

    println!("{SOLUTION_FORMAT}{}", prog_stack[0]);
}
