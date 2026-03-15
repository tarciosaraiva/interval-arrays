# interval-arrays

A CLI tool that merges overlapping integer intervals.

## Usage

```bash
go run main.go -interval "1-5,3-7"
# Output: Original: [[1 5] [3 7]], Merged: [[1 7]]
```

Intervals are specified as a comma-separated string of `start-end` pairs. Spaces around `-` and `,` are tolerated.

## How it works

1. Parses the input string into `[][]int`
2. Sorts intervals by start value
3. Merges overlapping or adjacent intervals into a minimal set

## Development

```bash
go test ./...                          # run all tests
go test -run Test_mergeIntervals       # run a specific test
go build -o interval-arrays            # build binary
```
