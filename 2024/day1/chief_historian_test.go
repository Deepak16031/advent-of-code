package main

import (
	"log"
	"testing"
)

func TestFindDistance(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	actual := FindDistance(a, b)
	expected := 11

	if actual != expected {
		log.Fatalf("expected %v actual %v", expected, actual)
	}

}

func TestFreqMap(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	actual := FindSimilarity(a, b)
	expected := 31
	if actual != expected {
		log.Fatalf("expected %v actual %v", expected, actual)
	}
}
