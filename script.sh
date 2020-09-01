#!/bin/bash

curl https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv |
  iconv -f SJIS -t UTF-8 |
  tr -d "\\r" |
  tr ',/' ' ' |
  awk 'NR>1{
    printf "{year: %d, month: %d, day: %d, description: \"%s\"},\n", $1, $2, $3, $4
  }'