package main

import (
	"log"
	"slices"
	"testing"
)

func TestBfs(t *testing.T) {

	t.Run("big test", func(t *testing.T) {
		arr := [][]string{
			{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
			{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
			{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
			{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
			{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
			{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
			{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
			{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
			{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
			{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
		}

		testInput(arr, 18)
	})

	t.Run("small test case", func(t *testing.T) {

		arr := [][]string{
			{"X", "M", "A", "S"},
		}

		testInput(arr, 1)
	})

	t.Run("reverse case", func(t *testing.T) {
		arr := [][]string{
			{"X", "M", "A", "S"},
		}

		slices.Reverse(arr[0])
		testInput(arr, 1)
	})

	t.Run("test Part 2", func(t *testing.T) {
		arr := [][]string{
			{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
			{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
			{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
			{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
			{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
			{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
			{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
			{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
			{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
			{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
		}

		testInputPart2(arr, 9)

	})

}

func testInput(arr [][]string, expected int) {
	actual := FindWordCount(arr, "XMAS")
	if actual != expected {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}

func testInputPart2(arr [][]string, expected int) {
	actual := FindWordCountPart2(arr)
	if actual != expected {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}
