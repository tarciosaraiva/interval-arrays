package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r, _ = regexp.Compile(`(\d\s?-\s?\d,?\s?)+`)

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

func main() {
	var intervals = flag.String("interval", "", "a collection of intervals defined like 7-8,1-5 ...")
	flag.Parse()

	parsedIntervals, err := parseIntervalInput(*intervals)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(parsedIntervals)
}
