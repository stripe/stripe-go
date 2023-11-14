#!/usr/bin/env bash

cd "$(dirname $0)" || exit 1
cd .. || exit 1
echo "Running sorbet typecheck"
bundle exec srb tc || exit 1
echo "Running rubocop"
bundle exec rubocop || exit 1