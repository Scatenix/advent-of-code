#!/bin/bash

# Arguments passed to this recipe from 'createNewPuzzleTemplate.sh' $1:tech $2:year $3:day $4:part $5:full-path-to-solution-dir

# Copy the go template to the solution directory as puzzle.go and replace some placeholders with year, day and part
cp templates/template-go $5/puzzle.go
sed -i s/'<<<YEAR>>>'/"$2"/ $5/puzzle.go
sed -i s/'<<<DAY>>>'/"$3"/ $5/puzzle.go
sed -i s/'<<<PART>>>'/"$4"/ $5/puzzle.go

# Inject this solution as a run configuration to intelliJs workspace.xml
cd $AOC_HOME
sed -e "s/<<<YEAR>>>/$2/g" -e "s/<<<DAY>>>/$3/g" -e "s/<<<PART>>>/$4/g" templates/idea-workspace-additions/go.ed > tmp.ed
ed .idea/workspace.xml < tmp.ed
rm tmp.ed
