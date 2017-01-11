#!/bin/bash

# This script supports the process of following the Granitic tutorials available
# at https://github.com/graniticio/granitic/tree/master/doc/tutorial
#
# The grantic-examples repo (where this script orignates) contains source code
# representing the end state of each tutorial.
#
# The tutorial documentation assumes that you are working through each tutorial in order but
# this script allows tutorials to be skipped.
# 
# For example, running 
#
#   prepare-tutorial.sh 5
#
# copies the source code for tutorial 004 to $GOPATH/src/granitic-tutorial
# which allows you to begin tutorial 005 with the correct code.

if [ -z "$BASH" ]; then
	"This script requires Bash" && exit 1
fi

type perl >/dev/null 2>&1 || { echo >&2 "This script requires Perl to be installed"; exit 1; }

USAGE="Usage: $( basename $0 ) tutorial-number"
[ -z "$1" ] && echo $USAGE && exit 1

[ -z "$GOPATH" ] && echo "Missing GOPATH environment variable" && exit 1

if [ ! -d "$GOPATH" ]; then
  echo "The directory that GOPATH is set to ($GOPATH) does not exist" && exit 1
fi

EXAMPLES_PATH="$GOPATH/src/github.com/graniticio/granitic-examples"

if [ ! -d "$EXAMPLES_PATH" ]; then
  echo "This script requires that the granitic-examples project be checked out to \$GOPATH/src/github.com/graniticio/granitic-examples" && exit 1
fi

if [ ! -d "$GOPATH/src/granitic-tutorial" ]; then
  echo "The directory \$GOPATH/src/granitic-tutorial does not exist, creating"
  mkdir -p $GOPATH/src/granitic-tutorial
fi

TUTORIAL=$1

re='^[0-9]+$'
if ! [[ $TUTORIAL =~ $re ]] ; then
   echo "The argument to this script must be a positive integer" && exit 1
fi

if [[ $TUTORIAL =~ ^0 ]]; then 
	echo "Please remove leading zeroes from the tutorial number" && exit 1 
fi

if (( $TUTORIAL < 2 )); then
    echo "This script only supports tutorial 2 and later" && exit 1
fi

TBASE=$[$TUTORIAL - 1]

TUTORIAL_FOLDER=$GOPATH/src/granitic-tutorial/recordstore

if [ -d "$TUTORIAL_FOLDER" ]; then
  echo "Removing existing tutorial folder"
  rm -rf $TUTORIAL_FOLDER
fi

if (( $TBASE < 10 )); then
	TINDEX="00$TBASE"
elif (( $TBASE < 100 )); then
	TINDEX="0$TBASE"
else
	TINDEX="$TBASE"
fi

TPATH="$EXAMPLES_PATH/tutorial/tutorial$TINDEX"

if [ ! -d "$TPATH" ]; then
	echo "No code for the previous tutorial ($TBASE) available" && exit 1
fi

cp -R $TPATH $TUTORIAL_FOLDER

find $TUTORIAL_FOLDER -name '*.json' -o  -name '*.go' | while read filename; do
	PATTERN="s|github.com/graniticio/granitic-examples/tutorial/tutorial${TINDEX}|granitic-tutorial/recordstore|g"
	perl -pi -e $PATTERN $filename
done


