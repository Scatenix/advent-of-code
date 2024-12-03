#!/bin/bash
set -eu # Abort script on error, avoid using unset variables

# Colors: 1=red, 2=green, 3=yellow, 4=blue, 5=pruple, 6=cyan, 7=gray, 8=dark_gray, 9-16=light-versions-of-colors
warn=$(tput setaf 3)
success=$(tput setaf 2)info=$(tput sgr0)

print() {
  printf '%b\n' "$1"
}

usage() {
  print "Advent of code: Script to create a new template directory for a specific technology, year, day and part"
	print "usage: createNewPuzzleTemplate.sh [TECHNOLOGY] [YEAR] [DAY] [PART]"
	print "Create both parts for todays AoC: createNewPuzzleTemplate.sh [TECHNOLOGY] today"

	print "\nUse 'createNewPuzzleTemplate.sh help' for more info"

	print "\nAOC_HOME is $AOC_HOME"
}

help() {
	usage

	print "\nTECHNOLOGY    Technology to create new puzzle template for"
	print "                 Currently supported within my AoC project: go, awk"
	print "                 But you can provide your own by reading the part about templates further down."
	print "YEAR          Year of AoC (2015 - current year)"
	print "DAY           Day of AoC (1 - 25)"
	print "PART          Part of AoC (1 or 2)"

  print "\nTemplates: Providing your own templates: create a recipe file with the name of the technology, followed by .sh (e.g. java.sh)."
  print "In there you can describe what needs to be done to setup the technology specific files"
  print "args supplied to recipe file in this order: [TECHNOLOGY] [YEAR] [DAY] [PART] [PATH_TO_CURRENT_PART]"
  print "${warn}CAUTION: Be careful about recipe files overwriting existing files! Having git set up and commiting regularly is advised!\n${info}"
	print "This script uses the AOC_HOME variable. It must be exported to your environment with the path to your advent-of-code project root."

	print "\nAOC_HOME is $AOC_HOME"
}

currentYear=$(date +%-Y)
reInt='^[0-9]+$' # check for value being an integer

ARGS_COUNT=4
ALT_ARGS_COUNT=2
if [ $# -eq 1 ]; then
	help
  exit
elif [ $# -eq "$ALT_ARGS_COUNT" ]; then
  if [ $2 = "today" ]; then
    tech=$1
    year="$currentYear"
    day=$(date +%-d)
    part=12
  else
    usage
    exit 1
  fi
elif [ $# -gt "$ARGS_COUNT" ] || [ $# -lt "$ARGS_COUNT" ]; then
	usage
  exit 1
else
  # User input arguments
  tech=$1
  year=$(sed 's/^0*//' <<< "$2")
  day=$(sed 's/^0*//' <<< "$3")
  part=$(sed 's/^0*//' <<< "$4")
fi

validateEnvironment() {
  if [ -z "$AOC_HOME" ]; then
    print "AOC_HOME not set. Please set AOC_HOME to the root directory of your advent-of-code project."
    print "Inside AoC dir do 'export AOC_HOME=$(pwd)'"
    print "Or put 'export AOC_HOME='<full path to AoC>'' into your shell rc file"
    exit 1
  fi
  if ! [ -e "$AOC_HOME"/templates/recipes/"$tech".sh ]; then
    print "$AOC_HOME/templates/recipes/$tech.sh does not exist, stopping..."
    exit 1
  fi
}

validateInputs() {
  if ! [[ "$year" =~ $reInt ]] || [ "$year" -lt 2015 ] || [ "$year" -gt "$currentYear" ]; then
      print "Provided year is in bad format (2015 - $currentYear)"
      exit 1
  fi
  if ! [[ "$day" =~ $reInt ]] || [ "$day" -lt 1 ] || [ "$day" -gt 25 ]; then
      print "Provided day is in bad format (1 - 25)"
      exit 1
  fi
  if ! [[ "$part" =~ $reInt ]] || { [ ! "$1" = "today" ] && { [ "$part" -lt 1 ] || [ "$part" -gt 2 ]; } }; then
      print "Provided part is in bad format (1 - 2)"
      exit 1
  fi
}

createDirStructure() {
  mkdir -p "$AOC_HOME"/"$year"/Day-"$day"/{Part-"$part"/"$tech",resources}
  print "${success}Created directory structure for $year $day $part.${info}"
  executeRecipe "$AOC_HOME/$year/Day-$day/Part-$part/$tech"
  print "${success}Executed recipe for $tech.${info}"
}

createDirStructures() {
  # if $part is 12, we know we are in the today mode
  if [ "$part" = 12 ]; then
    part=1
    dirAlreadyExists_skip "$AOC_HOME"/"$year"/Day-"$day"/Part-"$part"/"$tech"
    if [ ! "$skip" = true ]; then
      createDirStructure
    fi

    part=2
    dirAlreadyExists_proceed "$AOC_HOME"/"$year"/Day-"$day"/Part-"$part"/"$tech"
    createDirStructure
  else
    dirAlreadyExists_proceed "$AOC_HOME"/"$year"/Day-"$day"/Part-"$part"/"$tech"
    createDirStructure
  fi
}

executeRecipe() {
  sh $AOC_HOME/templates/recipes/$tech.sh $tech $year $day $part $1
}

dirAlreadyExists_proceed(){
  if [ -d "$1" ]; then
    read -p "${warn}WARNING: Directory $1 already exists. Proceed with executing recipe?(y) ${info}" -n 1 -r
    if [[ ! $REPLY =~ ^([Yy])$ ]]
    then
      print "\nStopping creation process..."
      exit 1
    fi
  fi
}

dirAlreadyExists_skip(){
  if [ -d "$1" ]; then
    read -p "${warn}WARNING: Directory $1 already exists. Proceed with executing recipe?(y) ${info}" -n 1 -r
    if [[ ! $REPLY =~ ^([Yy])$ ]]
    then
      print "\nSkipping Part-1"
      skip=true
    else
      skip=false
    fi
  fi
}

validateEnvironment
validateInputs $2
createDirStructures
