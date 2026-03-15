package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseIntervalString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"parse valid input", args{input: "1-5,3-7"}, [][]int{{1, 5}, {3, 7}}},
		{"parse valid input with space between ranges", args{input: "2-5, 4-8 , 10-15"}, [][]int{{2, 5}, {4, 8}, {10, 15}}},
		{"parse valid input with space between range items", args{input: "3 - 5,7 -9, 8- 12"}, [][]int{{3, 5}, {7, 9}, {8, 12}}},
		{"parse valid input with large numbers", args{input: "10 - 25,37 -43, 23- 41, 66-191,77-212"}, [][]int{{10, 25}, {37, 43}, {23, 41}, {66, 191}, {77, 212}}},
		{"throw error on alpha collection", args{input: "a-c,f-7"}, nil},
		{"throw error on alpha", args{input: "error"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseIntervalInput(tt.args.input)
			if tt.want == nil {
				assert.ErrorContains(t, err, "invalid pattern - should be n-n,n1-n1")
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.want, result)
			}
		})
	}
}

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"merge interval", args{input: [][]int{{1, 3}, {2, 4}, {6, 8}, {9, 10}}}, [][]int{{1, 4}, {6, 8}, {9, 10}}},
		{"merge out of order interval", args{input: [][]int{{7, 8}, {1, 5}, {2, 4}, {4, 6}}}, [][]int{{1, 6}, {7, 8}}},
		{"merge another out of order interval", args{input: [][]int{{10, 15}, {5, 7}, {6, 9}, {12, 16}, {9, 18}}}, [][]int{{5, 18}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeIntervals(tt.args.input)
			assert.Equal(t, tt.want, result)
		})
	}
}
