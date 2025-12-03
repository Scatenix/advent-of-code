#!/bin/bash

# cabal must be installed to use this script!

# Arguments passed to this recipe from 'createNewPuzzleTemplate.sh' $1:tech $2:year $3:day $4:part $5:full-path-to-solution-dir
cd $5
cabal init puzzle --non-interactive
/bin/rm puzzle/CHANGELOG.md
/bin/rm puzzle/app/Main.hs

# copy template and substitute placeholders
cp $AOC_HOME/templates/template-haskell $5/puzzle/app/Main.hs
sed -i s/'<<<YEAR>>>'/"$2"/ $5/puzzle/app/Main.hs
sed -i s/'<<<DAY>>>'/"$3"/ $5/puzzle/app/Main.hs
sed -i s+'<<<AOC_HOME>>>'+"$AOC_HOME"+ $5/puzzle/app/Main.hs

# Set up aocutil libraries
sed -i '/^[[:space:]]*build-depends:/s/$/, puzzle/' $5/puzzle/puzzle.cabal
printf "\nlibrary
\ths-source-dirs: ../../../../../aocutil/haskell/
\texposed-modules: AocUtils
\tbuild-depends:       base >=4.14 && <5
\tdefault-language:    Haskell2010" >> $5/puzzle/puzzle.cabal

# This is a small hack that is needed for the Haskell LSP to be able to properly parse the sub-projects
sed -i '/^library$/{
    n
    /^[[:space:]]*hs-source-dirs:/s|$|\n        '"$2"'/Day-'"$3"'/Part-'"$4"'/haskell/puzzle/app,|
}' "$AOC_HOME/advent-of-code.cabal"

# Perform clean build, mainly to ensure a more stable Haskell LSP
cd puzzle
cabal clean
cabal build