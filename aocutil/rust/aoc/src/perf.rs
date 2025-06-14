use std::time::Instant;

// Track time of current scope
// usage: defer!{aoc::perf::time_tracker(std::time::Instant::now(), "main")}
pub fn time_tracker(start: Instant, name: &str) {
    println!("{} took {:.2?}", name, start.elapsed());
}
