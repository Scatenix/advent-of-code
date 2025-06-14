#!/usr/bin/awk -f

# Advend of Code 2024 - Day 1 Part 2 (of 2)
# Usage: ./similarity-score.awk ../resources/distances

{
    col1[$1]        # Collect unique values from column 1
    col2[$2]++      # Count occurrences of each value in column 2
	# The trick here is to use the value of column 2 as index and every time we have this "index" on a line, we increment the value in there (which just starts at 0 implicitly).
	# In the END function, we simply search for the value of col1 in the col2 array by accessing the col1 value as index of col2. If something is present, we have the count of it in the col2 array.
}

END {
    for (value in col1) {
    	similarityScore += value*(value in col2 ? col2[value] : 0)
	}
	print similarityScore
}
