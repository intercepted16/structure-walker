#!/bin/bash
watchman-make -p '*.c' '*.go' '*/**/*.c' '*/**/*.go' -r 'go build . && ./data-structures'

