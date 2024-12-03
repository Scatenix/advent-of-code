#!/bin/bash

cp templates/template-go $5/puzzle.go
sed -i s/'<<<DAY>>>'/"$3"/ $5/puzzle.go
sed -i s/'<<<PART>>>'/"$4"/ $5/puzzle.go
