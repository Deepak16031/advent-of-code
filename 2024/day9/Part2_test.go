package main

import (
	"strings"
	"testing"
)

func TestSolution2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			"basic test case",
			"2333133121414131402",
			2858,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := Solution2(strings.Split(tt.input, "")); got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}
		})
	}
}
