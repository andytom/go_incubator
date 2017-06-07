#!/bin/bash
set -euo pipefail

# -- Find our test directories --
dirs=$(glide novendor)

echo "-- Running Tests --"
go test -cover ${dirs}

echo "-- Running Vet --"
go vet ${dirs}

echo "-- Running Lint --"
golint ${dirs}

echo "-- All Done --"
