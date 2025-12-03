#!/bin/bash

# cabal must be installed to use this script!

# Arguments passed to this recipe from 'createNewPuzzleTemplate.sh' $1:tech $2:year $3:day $4:part $5:full-path-to-solution-dir
cd $5
cabal init puzzle --non-interactive

cp templates/template-haskell $5/puzzle.hs
