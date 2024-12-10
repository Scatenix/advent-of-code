use std::fs::File;
use std::io::{BufRead, BufReader};

/// # Read puzzle file to a collection
///
/// Read the puzzle file to a collection line by line through reading_handler closure
///
/// ## Parameters
///
/// file: std::fs::File - puzzle input file
///
/// reading_handler: fn(line: String, collection: T) -> T - closure that transforms the stream of lines into the defined collection
///
/// ## Return value
///
/// Returns the collection of the defined type to be used for the next line
///
/// ## Examples for reading_handler
///
/// ```
///     let reading_hanlder = |line: String, mut col: String| -> String {
///         col.push_str(&line);
///         col.push('\n');
///         return col
///     };
///
///     let reading_hanlder = |line: String, mut col: Vec<String>| -> Vec<String> {
///         col.push(line.split(' ').collect());
///         return col
///     };
///
///     let reading_hanlder = |line: String, mut col: Vec<String>| -> Vec<String> {
///         col.push(line.chars().collect());
///         return col
///     };
///
///     let reading_hanlder = |line: String, mut col: Vec<Vec<u32>>| -> Vec<Vec<u32>> {
///         let digits: Vec<u32> = line.chars()
///             .filter_map(|c| c.to_digit(10))
///             .collect();
///         col.push(digits);
///         return col
///     };
/// ```
///
pub fn read_puzzle_file<T:Default>(file: File, reading_handler: fn(line: String, collection: T) -> T) -> T {
    let reader = BufReader::new(file);

    let mut col: T = T::default();

    for line in reader.lines() {
        col = reading_handler(line.unwrap(), col);
    }
    return col;
}