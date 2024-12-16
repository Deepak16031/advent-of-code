package main

import (
	"fmt"
	"testing"

	"advent-of-code-in-go/util"
)

func TestSolution(t *testing.T) {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day10/input_test.txt"
	arr := util.ReadFileAs2DArr(filename, "")

	actual := Solution(arr, false)
	expected := 36

	if actual != expected {
		fmt.Printf("expected %v actual %v", expected, actual)
	}

	actual = Solution(arr, true)
	expected = 81

	if actual != expected {
		fmt.Printf("Part 2: expected %v actual %v", expected, actual)
	}
}
