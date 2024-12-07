package main

import (
	"log"
	"slices"
	"testing"
)

func TestCountSafeReports(t *testing.T) {
	//
	//a := [][]int{
	//	{7, 6, 4, 2, 1},
	//	{1, 2, 7, 8, 9},
	//	{9, 7, 6, 2, 1},
	//	{1, 3, 2, 4, 5},
	//	{8, 6, 4, 4, 1},
	//	{1, 3, 6, 7, 9}}
	//
	//actual := CountSafeReports(a)
	//expected := 2
	//
	//if actual != expected {
	//	log.Fatalf("expected %v actual %v", expected, actual)
	//}

	b := [][]int{{63, 66, 67, 66, 69, 72}}
	actual := CountSafeReports(b)
	expected := 1

	if actual != expected {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}

func TestIsValidReport(t *testing.T) {

	t.Run("positive report", func(t *testing.T) {
		b := []int{56, 57, 58, 61, 64, 66, 68}
		if !isValidReport(b, true) {
			log.Fatalf("wrong interpreration 1")
		}
		if isValidReport(b, false) {
			log.Fatalf("wrong interpreration 2")
		}
	})

	t.Run("negative report", func(t *testing.T) {
		b := []int{56, 57, 58, 61, 64, 66, 68}
		if isValidReport(b, false) {
			log.Fatalf("wrong interpreration 3")
		}
		slices.Reverse(b)
		if !isValidReport(b, true) {
			log.Fatalf("wrong interpreration 4")
		}
	})

}
