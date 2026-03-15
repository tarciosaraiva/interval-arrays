package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var r, _ = regexp.Compile(`(\d\s?-\s?\d,?\s?)+`)

// parseIntervalInput takes an input string and parse it to match
// the regex specified by r - throws an error if string does not conform
func parseIntervalInput(input string) ([][]int, error) {
	m := r.MatchString(input)
	if !m {
		return nil, errors.New("invalid pattern - should be n-n,n1-n1")
	}

	var parsedIntervals [][]int
	for i := range strings.SplitSeq(input, ",") {
		var singleInterval []int
		for j := range strings.SplitSeq(strings.TrimSpace(i), "-") {
			n, _ := strconv.Atoi(strings.TrimSpace(j))
			singleInterval = append(singleInterval, n)
		}
		parsedIntervals = append(parsedIntervals, singleInterval)
	}

	return parsedIntervals, nil
}

// mergeIntervals takes an array of intervals and
// 1. sort them
// 2.
func mergeIntervals(intervals [][]int) [][]int {
	var merged = [][]int{}
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	for _, currentInterval := range intervals {
		var lastMergedInterval []int
		var mergedLength = len(merged)

		// add the first item to the merged interval
		// and always set lastMergedInterval
		if mergedLength == 0 {
			merged = append(merged, currentInterval)
			lastMergedInterval = currentInterval
		} else {
			lastMergedInterval = merged[mergedLength-1]
		}

		// append to merged when the start of the current interval
		// is greater than the end of the previous interval
		if currentInterval[0] > lastMergedInterval[1] {
			merged = append(merged, currentInterval)
		} else if currentInterval[1] > lastMergedInterval[1] {
			var newInterval = []int{lastMergedInterval[0], currentInterval[1]}

			if lastMergedInterval[0] < currentInterval[0] {
				newInterval = []int{lastMergedInterval[0], currentInterval[1]}
			}
			merged = slices.Replace(merged, 0, 1, newInterval)
		}

		// fmt.Printf("Current interval: %v - Merged: %v\n", currentInterval, merged)
	}

	return merged
}

func main() {
	var intervals = flag.String("interval", "", "a collection of intervals defined like 7-8,1-5 ...")
	flag.Parse()

	parsedIntervals, err := parseIntervalInput(*intervals)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	merged := mergeIntervals(parsedIntervals)

	fmt.Printf("Original: %v, Merged: %v", parsedIntervals, merged)
}
