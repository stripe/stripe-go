#!/bin/bash

find_files() {
  find . -name '*.go' -not -path "./vendor/*" -not -path "./.git/*"
}

bad_files=$(find_files | xargs gofmt -s -l)

# if first argument is check
if [[ "$1" == "check" ]]; then
    if [[ -n "${bad_files}" ]]; then
        echo "!!! gofmt -s needs to be run on the following files: "
        echo "${bad_files}"
        exit 1
    fi
fi

for file in ${bad_files}; do
    gofmt -s -w "${file}"
done
exit 0
