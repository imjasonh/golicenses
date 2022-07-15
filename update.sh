#!/usr/bin/env bash

set -euxo pipefail

time cat query.txt | bq query --nouse_legacy_sql --format=csv --headless=true --max_rows=10000000 > licenses.csv

wc -l licenses.csv

gzip -f -k licenses.csv

go build ./cmd/golicenses-example && ls -lh golicenses-example licenses.*
rm golicenses-example
