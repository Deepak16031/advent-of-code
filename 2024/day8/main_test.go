package main

import (
	"log"
	"testing"

	"advent-of-code-in-go/util"
)

func TestPart1(t *testing.T) {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day8/input_test.txt"
	arr := util.ReadFileAs2DArr(filename, "")
	actual := Day8(arr, false)
	expected := 14
	if actual != expected {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day8/input_test_part2.txt"
	arr := util.ReadFileAs2DArr(filename, "")
	actual := Day8(arr, true)
	expected := 9
	if actual != expected {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}
