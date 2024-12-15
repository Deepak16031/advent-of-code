package main

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestDiskImage(t *testing.T) {
	arr := strings.Split("12345", "")
	actual := DiskImage(arr)
	expected := strings.Split("0..111....22222", "")

	if !reflect.DeepEqual(actual, expected) {
		log.Fatalf("expected %v actual %v", expected, actual)
	}

}

func TestCompaction(t *testing.T) {
	type args []string
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"basic test",
			args(strings.Split("0..111....22222", "")),
			strings.Split("022111222......", ""),
		},
		{
			"advances test",
			args(strings.Split("00...111...2...333.44.5555.6666.777.888899", "")),
			strings.Split("0099811188827773336446555566..............", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compaction(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checksum(t *testing.T) {
	type args []string
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{"basic test",
			args(strings.Split("0099811188827773336446555566..............", "")),
			1928,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := checksum(tt.args); gotSum != tt.wantSum {
				t.Errorf("checksum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestSolution(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			"basic test case",
			"2333133121414131402",
			1928,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := Solution(strings.Split(tt.input, "")); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
