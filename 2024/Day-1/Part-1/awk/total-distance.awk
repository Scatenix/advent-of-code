#!/usr/bin/awk -f

{
	col1[NR] = $1
	col2[NR] = $2
}

END {
	asort(col1, sortedCol1)
	asort(col2, sortedCol2)

	for (i = 1; i <= NR; i++) {
        diff = sortedCol2[i] - sortedCol1[i]
		totalDistance += (diff >= 0) ? diff : -diff
	}
	print totalDistance
}
