use aoc;

#[macro_use(defer)]
extern crate scopeguard;

const FALLBACK_PUZZLE_INPUT_PATH: &str = "2015/Day-2/resources/puzzle-input";
const DAY_PART: &str = "2015 Day 2 - Part 1";
const SOLUTION_FORMAT: &str = ">>> The solution is: ";

#[derive(Debug)]
struct Present {
    h: i32,
    l: i32,
    w: i32
}

// Usage: app <PATH_TO_PUZZLE_FILE>
fn main() {
    defer!{aoc::perf::time_tracker(std::time::Instant::now(), "main")}
    let file = aoc::util::aoc_setup(DAY_PART, FALLBACK_PUZZLE_INPUT_PATH);

    let reading_handler = |line: String, mut col: Vec<Present>| -> Vec<Present> {
        let hlw: Vec<&str> = line.split('x').collect();
        let present: Present = Present{
            h: hlw[0].parse::<i32>().unwrap(),
            l: hlw[1].parse::<i32>().unwrap(),
            w: hlw[2].parse::<i32>().unwrap()
        };
        col. push(present);
        return col
    };

    let puzzle_input = aoc::io::read_puzzle_file(file, reading_handler);

    let mut total_wrapping_paper = 0;
    for present in puzzle_input.iter() {
        let side1 = present.w*present.l;
        let side2 = present.h*present.l;
        let side3 = present.h*present.w;
        let smallest = side1.min(side2).min(side3);
        total_wrapping_paper += side1*2 + side2*2 + side3*2 + smallest;
    }

    println!("{SOLUTION_FORMAT}{}", total_wrapping_paper);
}
