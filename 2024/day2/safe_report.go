package main

import (
	"reflect"
	"slices"
)

//func CountSafeReportsPar2(arr [][]int) int {
//
//	safeReportsCount := 0
//
//	for _, report := range arr {
//		isPositiveReport := report[1] > report[0]
//		if isValidReportPart2(report, isPositiveReport) {
//			safeReportsCount++
//		}
//	}
//
//	return safeReportsCount
//}

func CountSafeReports(arr [][]int) int {

	safeReportsCount := 0

	for _, report := range arr {
		isPositiveReport := report[1] > report[0]
		if isValidReport(report, isPositiveReport) {
			safeReportsCount++
		}
	}

	return safeReportsCount
}

func isValidReport(arr []int, isPositiveReport bool) bool {
	originalCopy := make([]int, len(arr))
	copy(originalCopy, arr)

	slices.Sort(arr)
	if !isPositiveReport {
		slices.Reverse(arr)
	}
	if !reflect.DeepEqual(originalCopy, arr) {
		return false
	}

	for i, _ := range originalCopy {
		if i == 0 {
			continue
		}
		if !withInRange(originalCopy[i], originalCopy[i-1], 1, 3) {
			return false
		}
	}
	return true
}

//
//func isValidReportPart2(arr []int) bool {
//	//dampenseUsed := false
//	for i := 0; i < len(arr); i++ {
//		if i == 0 {
//			continue
//		}
//
//	}
//}

func withInRange(a, b, min, max int) bool {
	distance := b - a
	if distance < 0 {
		distance *= -1
	}
	return distance >= min && distance <= max
}
