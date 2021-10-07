#!/bin/bash

go install ./...

resolver -src=ru -dst=./docs/
cp README.md docs/readme.md
cp README.md docs/index.md