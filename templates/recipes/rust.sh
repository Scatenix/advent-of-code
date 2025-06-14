#!/bin/bash

# Arguments passed to this recipe from 'createNewPuzzleTemplate.sh' $1:tech $2:year $3:day $4:part $5:full-path-to-solution-dir

# Deleting the created rust directory to let cargo recreate it
cd $5/..
rmdir rust
cargo new rust-"$2-$3-$4"

# Copy the rust template to the solution directory as main.rs and replace some placeholders with year, day and part
cd rust-"$2-$3-$4"/src
cp $AOC_HOME/templates/template-rust ./main.rs
sed -i s/'<<<YEAR>>>'/"$2"/ ./main.rs
sed -i s/'<<<DAY>>>'/"$3"/ ./main.rs
sed -i s/'<<<PART>>>'/"$4"/ ./main.rs

# Make my rust utils available for the new solution
printf 'aoc = { path = "../../../../aocutil/rust/aoc" }\n' >> ../Cargo.toml

# Install the scopeguard dependency, needed for the defer! macro which is used to meassure exeuction times.
printf 'scopeguard = "1.2.0"\n' >> ../Cargo.toml
