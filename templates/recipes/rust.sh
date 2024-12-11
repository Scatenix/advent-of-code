#!/bin/bash

cd $5/..
rmdir rust
cargo new rust-"$2-$3-$4"

cd rust-"$2-$3-$4"/src
cp $AOC_HOME/templates/template-rust ./main.rs
sed -i s/'<<<YEAR>>>'/"$2"/ ./main.rs
sed -i s/'<<<DAY>>>'/"$3"/ ./main.rs
sed -i s/'<<<PART>>>'/"$4"/ ./main.rs

printf 'aoc = { path = "../../../../aocutil/rust/aoc" }\n' >> ../Cargo.toml
printf 'scopeguard = "1.2.0"\n' >> ../Cargo.toml
