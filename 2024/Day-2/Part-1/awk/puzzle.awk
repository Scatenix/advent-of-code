#!/usr/bin/awk -f

{
    levelsIncreasing = $1 < $2
    for(i=1; i<=NF; i++) {
        if(i < NF) {
            if (levelsIncreasing != ($(i) < $(i+1))) {
                break
            }

            diff = $(i) - $(i+1)
            diffAbs = (diff >= 0) ? diff : -diff
            if (diffAbs < 1 || diffAbs > 3) {
                break
            }
        } else {
            safeReports++
        }
    }
}

END {
    print safeReports
}
