package main

import (
	"log"
	"testing"

	"advent-of-code-in-go/util"
)

func TestOrder(t *testing.T) {
	rules := map[int]map[int]bool{
		75: {
			47: true,
			61: true,
			53: true,
			29: true,
		},
		47: {
			61: true,
			53: true,
			29: true,
		},
		61: {
			53: true,
			29: true,
		},
		53: {
			70: true,
			79: true,
		},
		29: {},
		70: {
			79: true,
		},
		81: {
			90: true,
			70: true,
			79: true,
		},
	}

	t.Run("incorrect order - 1", func(t *testing.T) {
		arr := []int{75, 47, 61, 53, 29, 79, 70}
		assertCorrectOrder(arr, rules, false)
	})
	t.Run("correct order - 1", func(t *testing.T) {
		arr := []int{75, 47, 61, 53, 29, 79}
		assertCorrectOrder(arr, rules, true)
	})
	t.Run("incorrect order - 2", func(t *testing.T) {
		arr := []int{75, 47, 61, 53, 29, 80, 81, 70, 79, 53}
		assertCorrectOrder(arr, rules, false)
	})
}

func TestSumOfMiddleOfValidUpdates(t *testing.T) {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day5/test_input.txt"
	rulesInputArr := util.ReadAllAsArr(filename)
	indexOfEmptyLine := util.IndexOfEmptyLineInArr(rulesInputArr)
	if indexOfEmptyLine != 21 {
		log.Fatalf("expected indx of empty line - 21 actual - %v", indexOfEmptyLine)
	}
	rules := CreateRules(rulesInputArr[:indexOfEmptyLine])
	updates := rulesInputArr[indexOfEmptyLine:]

	actual := SumOfMiddleOfValidUpdates(updates, rules, false)
	expected := 143

	if expected != actual {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}

func TestSumOfMiddleOfValidUpdates2(t *testing.T) {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day5/test_input.txt"
	rulesInputArr := util.ReadAllAsArr(filename)
	indexOfEmptyLine := util.IndexOfEmptyLineInArr(rulesInputArr)
	if indexOfEmptyLine != 21 {
		log.Fatalf("expected indx of empty line - 21 actual - %v", indexOfEmptyLine)
	}
	rules := CreateRules(rulesInputArr[:indexOfEmptyLine])
	updates := rulesInputArr[indexOfEmptyLine:]

	actual := SumOfMiddleOfValidUpdates(updates, rules, false)
	expected := 143

	if expected != actual {
		log.Fatalf("expected %v actual %v", expected, actual)
	}

}

func assertCorrectOrder(arr []int, rules map[int]map[int]bool, expected bool) {
	actual := IsCorrectOrder(arr, rules)

	if expected != actual {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}
