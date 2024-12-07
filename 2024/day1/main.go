package main

import (
	"fmt"

	"advent-of-code-in-go/util"
)

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day1/day1_inputA.txt"
	arr1, arr2 := util.ReadFile(filename)
	fmt.Printf("Distance %v", FindDistance(arr1, arr2))
}
