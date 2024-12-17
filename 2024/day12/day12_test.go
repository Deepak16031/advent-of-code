package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		arr [][]string
	}
	tests := []struct {
		name string
		args [][]string
		want int
	}{
		{
			"basic test",
			[][]string{
				{"A", "A", "A", "A"},
				{"B", "B", "C", "D"},
				{"B", "B", "C", "C"},
				{"E", "E", "E", "C"},
			},
			140,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
