#!/bin/sh

set -e

cd $(dirname $0)
cd ..

NAME=${1}
shift

echo "$NAME"
rm -fr test/output

docker run --rm \
           -v "$(pwd)/test/input:/input:ro" \
           -v "$(pwd)/test/output:/output" \
           "${NAME}" \
           "$@" \
           /input/input.txt

echo
echo '## output file content ##'
echo
cat test/output/*
