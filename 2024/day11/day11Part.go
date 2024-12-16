package main

import (
	"math"
	"slices"

	"advent-of-code-in-go/util"
)

func Solution(arr []int) int {

	for k := 0; k < 25; k++ {
		for i := 0; i < len(arr); i++ {
			val := arr[i]
			if val == 0 {
				arr[i] = 1
				continue
			}
			digits := util.CountDigits(val)
			if digits%2 == 0 {
				quotient := int(math.Pow(10, float64(digits/2)))
				leftHalf := val / quotient
				rightHalf := val % quotient
				arr[i] = leftHalf
				arr = slices.Insert(arr, i+1, rightHalf)
				i++
			} else {
				arr[i] = arr[i] * 2024
			}
		}
		//fmt.Println(arr)
	}
	return len(arr)
}
