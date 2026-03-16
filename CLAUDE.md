# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the CLI
go run main.go -interval "1-5,3-7"

# Build
go build -o interval-arrays

# Run all tests
go test ./...

# Run a specific test
go test -run Test_mergeIntervals
go test -run Test_parseIntervalString
```

## Architecture

Single-package Go CLI (`package main`) with two core functions:

- `parseIntervalInput(input string) ([][]int, error)` — validates the input string against a regex pattern (`n-n,n1-n1` format, spaces allowed) and parses it into a slice of `[start, end]` int pairs.
- `mergeIntervals(intervals [][]int) [][]int` — sorts intervals by start value, then merges overlapping/adjacent ones into a minimal set.

The `main` function wires these together via the `-interval` flag.

## Input format

Intervals are passed as a string: `"1-5, 3-7, 10-15"`. Spaces around `-` and `,` are tolerated. Only single-digit bounds are supported by the current regex.
