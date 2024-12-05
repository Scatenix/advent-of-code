#!/bin/bash

cp templates/template-go $5/puzzle.go
sed -i s/'<<<DAY>>>'/"$2"/ $5/puzzle.go
sed -i s/'<<<DAY>>>'/"$3"/ $5/puzzle.go
sed -i s/'<<<PART>>>'/"$4"/ $5/puzzle.go

cd $AOC_HOME
sed -e "s/<<<YEAR>>>/$2/g" -e "s/<<<DAY>>>/$3/g" -e "s/<<<PART>>>/$4/g" templates/idea-workspace-additions/go.ed > tmp.ed
ed .idea/workspace.xml < tmp.ed
rm tmp.ed
