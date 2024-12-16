package main

import (
	"fmt"
	"math"
	"slices"

	"advent-of-code-in-go/util"
)

func main() {
	arr := []int{2701, 64945, 0, 9959979, 93, 781524, 620, 1}
	fmt.Println(Solution(arr))
}

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
