#!/bin/sh
function main() {
  local NAMESPACE=$(echo $1 | sed '/\/$/! s|$|/|')
  local EXERCISE=$2
  local LAST_FILE=$(ls $NAMESPACE | sort -V | tail -n 1)
  local LAST_COUNT="${LAST_FILE%%_*}"
  local NEXT_COUNT=$(($LAST_COUNT+1))
  local DIRECTORY=$NAMESPACE${NEXT_COUNT}_$EXERCISE
  local SCRIPT="$DIRECTORY/$EXERCISE.go"

  mkdir $DIRECTORY
  touch $SCRIPT
  cat << EOF >> $SCRIPT
package $EXERCISE

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func New() {}
EOF

  cd $DIRECTORY
  ginkgo bootstrap
  ginkgo generate

  echo "Scaffolded $DIRECTORY"
}

main $*
