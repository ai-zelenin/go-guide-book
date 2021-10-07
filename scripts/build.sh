#!/bin/bash

go install ./...

resolver -src=ru -dst=./docs/
cp README.md docs/index.md

resolver -src=ru -dst=./go-guide-book.wiki/
cp README.md go-guide-book.wiki/index.md