use std::env;
use std::fs::File;

pub fn aoc_setup(day_part: &str, fallback_puzzle_path: &str) -> File {
    let mut args: Vec<String> = env::args().collect();
    println!("{:?}", env::current_dir());
    if args.len() < 2 {
        args.push(fallback_puzzle_path.to_owned());
    }

    let arg_error_msg = &format!("Please specify a puzzle-input file for AoC {day_part}");
    let filename = args.get(1).expect(arg_error_msg);
    return File::open(filename).expect(arg_error_msg)
}
