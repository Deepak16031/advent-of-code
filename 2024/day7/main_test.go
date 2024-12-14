package main

import (
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("test input with additional data", func(t *testing.T) {
		arr := [][]int{
			{190, 10, 19},
			{3267, 81, 40, 27},
			{83, 17, 5},
			{156, 15, 6},
			{7290, 6, 8, 6, 15},
			{161011, 16, 10, 13},
			{192, 17, 8, 14},
			{21037, 9, 7, 18, 13},
			{292, 11, 6, 16, 20},
			{11111, 11, 10, 1, 10, 1, 10, 1},
			{11111, 1, 1, 5, 1, 10, 90, 5, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100},
			{168, 7, 1, 6, 2, 2},
		}
		expected := 11387 + 11111*2 + 168
		assertValues(arr, expected)
	})

	t.Run("Part 2 small input", func(t *testing.T) {
		arr := [][]int{
			{156, 15, 6},
		}
		assertValues(arr, 156)
	})
}

func assertValues(arr [][]int, expected int) {
	actual := Part1(arr)
	if expected != actual {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}
