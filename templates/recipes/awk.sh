#!/bin/bash

# Arguments passed to this recipe from 'createNewPuzzleTemplate.sh' $1:tech $2:year $3:day $4:part $5:full-path-to-solution-dir

cp templates/template-awk $5/puzzle.awk
chmod +x $5/puzzle.awk
