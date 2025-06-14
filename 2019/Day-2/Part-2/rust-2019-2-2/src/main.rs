use aoc;

#[macro_use(defer)]
extern crate scopeguard;

const FALLBACK_PUZZLE_INPUT_PATH: &str = "2019/Day-2/resources/puzzle-input";
const DAY_PART: &str = "2019 Day 2 - Part 2";
const SOLUTION_FORMAT: &str = ">>> The solution is: ";

// Usage: app <PATH_TO_PUZZLE_FILE>
fn main() {
    defer!{aoc::perf::time_tracker(std::time::Instant::now(), "main")}
    let file = aoc::util::aoc_setup(DAY_PART, FALLBACK_PUZZLE_INPUT_PATH);
    let prog_stack = aoc::intcode::prog_stack_reader(file);

    let (noun, verb) = search_pair(prog_stack);
    println!("{SOLUTION_FORMAT}{}", 100 * noun + verb);
}

fn search_pair(prog_stack: Vec<isize>) -> (isize, isize) {
    for noun in 0..100 {
        for verb in 0..100 {
            let mut stack_clone = prog_stack.clone();
            stack_clone[1] = noun;
            stack_clone[2] = verb;
            let mut inst_pointer = 0;
            while inst_pointer != -1 {
                inst_pointer = aoc::intcode::intcode_interpreter(&mut stack_clone, inst_pointer as usize);
            }
            if stack_clone[0] == 19690720 {
                return (noun, verb);
            }
        }
    }
    panic!("No pair found that leads to address 0 being 19690720");
}