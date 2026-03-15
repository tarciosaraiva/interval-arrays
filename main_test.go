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
