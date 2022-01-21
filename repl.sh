#!/bin/sh
set -e

go install github.com/traefik/yaegi/cmd/yaegi@v0.11.2

cd cmd/repl

cleanup() {
  rm -f repl
}
trap cleanup EXIT

while :; do
  yaegi extract -name main github.com/dottedmag/goreplhowto
  go build -o repl .

  if ./repl; then
    :
  else
    if [ $? -eq 100 ]; then
       exit 0
    fi
  fi
done
