package main

import "slices"

func FindDistance(arr1, arr2 []int) int {
	var distance []int

	slices.Sort(arr1)
	slices.Sort(arr2)
	for i := 0; i < len(arr1); i++ {
		diff := absDiff(arr1[i], arr2[i])
		distance = append(distance, diff)
	}

	return sum(distance)
}

func sum(arr []int) int {
	var s int

	for _, a := range arr {
		s += a
	}
	return s
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
