#!/usr/bin/sh

# Advent of Code 2024, Day 1, Part 1 (of 2)

leftDistances=$(awk '{print $1}' ../resources/distances | sort -g)
rightDistances=$(awk '{print $2}' ../resources/distances | sort -g)

sortedDistances=$(paste <(printf "$leftDistances") <(printf "$rightDistances"))

awk '{diff=$2-$1; totalDistance += (diff >= 0) ? diff : -diff} END{print totalDistance}' <<< "$sortedDistances"
