#!/bin/bash

cd "$(dirname "$0")/.."

VERBOSE=0
STATS=0
DEBUG=0
SHOW_GOOD=1
SHOW_PROBLEMS=1

while getopts ':dsvgp' flag; do
  case "${flag}" in
    g) SHOW_PROBLEMS=0 ;;
    p) SHOW_GOOD=0 ;;
    d) DEBUG=1 ;;
    s) STATS=1 ;;
    v) VERBOSE=1 ;;
    *) error "Unexpected option ${flag}" ;;
  esac
done

if [ ! -z "$CI" ]; then
  SHOW_GOOD=0
  SHOW_PROBLEMS=1
fi


function getObjectName() {
  local path=$1
  local filename=${path##*/}
  local object=${filename%%.*}
  echo $object
}

function debugOutput() {
  if [ $DEBUG -eq 1 ]; then
    echo "$@"
  fi
}

function problem() {
  if [ $SHOW_PROBLEMS -eq 1 ]; then
    echo "❌ $*"
  fi
}

function good() {
  if [ $SHOW_GOOD -eq 1 ]; then
    echo "👍 $*"
  fi
}

function stat() {
  echo "📊 $*"
}

BASEPATH=force-app/main/default/objects
OLD_IFS=$IFS
IFS=$'\n'

MISSING_HELP_TEXT=0
MISSING_HELP_TEXT_FILES=""
MISSING_DESCRIPTION=0
MISSING_DESCRIPTION_FILES=""
WITH_HELP_TEXT=0
WITH_DESCRIPTION=0
WITHOUT_HELP_TEXT=0
WITHOUT_DESCRIPTION=0

DIRS=`find $BASEPATH/* -type d -maxdepth 0`
for dir in $DIRS; do
  object=`getObjectName $dir`
  echo "🔁 Checking $object"
  FILES=`find $dir -type f -name '*.field-meta.xml'`
  for file in $FILES; do
    fieldname=`getObjectName $file`
    help=`grep "<inlineHelpText>" $file | sed "s/.*>\\(.*\\)<.*/\1/"`
    debugOutput $help
    if [ -z "$help" ] || [[ "$help" == "" ]]; then
      problem "$object.$fieldname is missing inline help text."
      WITHOUT_HELP_TEXT=$((WITHOUT_HELP_TEXT+1))
      MISSING_HELP_TEXT=1
      MISSING_HELP_TEXT_FILES="$MISSING_HELP_TEXT_FILES $file"
    else
      WITH_HELP_TEXT=$((WITH_HELP_TEXT+1))
      good "$object.$fieldname has inline help text."
    fi
    desc=`grep "<description>" $file | sed "s/.*>\\(.*\\)<.*/\1/"`
    debugOutput $desc
    if [ -z "$desc" ] || [[ "$desc" == "" ]]; then
      WITHOUT_DESCRIPTION=$((WITHOUT_DESCRIPTION+1))
      problem "$object.$fieldname is missing a description."
      MISSING_DESCRIPTION=1
      MISSING_DESCRIPTION_FILES="$MISSING_DESCRIPTION_FILES $file"
    else
      WITH_DESCRIPTION=$((WITH_DESCRIPTION+1))
      good "$object.$fieldname has a description."
    fi
  done
done

EXITVAL=0
if [ $MISSING_HELP_TEXT -eq 1 ] || [ $MISSING_DESCRIPTION -eq 1 ]; then
  EXITVAL=1
fi

if [ $VERBOSE -eq 1 ]; then
  if [ $EXITVAL -eq 1 ]; then
    problem "There are problems to address."
    IFS=$' '
    if [ $MISSING_HELP_TEXT -eq 1 ]; then
      problem "Missing inline help text:"
      for file in $MISSING_HELP_TEXT_FILES; do
        echo $file
      done
    fi
    if [ $MISSING_DESCRIPTION -eq 1 ]; then
      problem "Missing description:"
      for file in $MISSING_DESCRIPTION_FILES; do
        echo $file
      done
    fi
  else
    good "Everything is good!"
  fi
fi

if [ $STATS -eq 1 ]; then
  TOTAL=$((WITH_HELP_TEXT+WITHOUT_HELP_TEXT))
  stat "$WITH_HELP_TEXT/$TOTAL fields have inline help text."
  TOTAL=$((WITH_DESCRIPTION+WITHOUT_DESCRIPTION))
  stat "$WITH_DESCRIPTION/$TOTAL fields have a description."
fi

exit $EXITVAL