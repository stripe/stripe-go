#!/bin/bash

# go fmt supports ./... but gofmt does not.
# but go fmt always runs gofmt with -w, which replaces
# which we don't desire
# hence the need for this annoying script
  

# go fmt -n prints the command that *would*
# be run by go fmt ./...
# so we just mangle that a bit
bad_files=$(gofmt -l $(go fmt -n ./... | cut -d ' ' -f4-))
if [[ -n "${bad_files}" ]]; then
    echo "!!! gofmt -w needs to be run on the following files: "
    echo "${bad_files}"
    exit 1
fi
