package main

import (
	"fmt"

	"advent-of-code-in-go/util"
)

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day2/day2_input.txt"
	arr := util.ReadNumbersTo2DArr(filename)
	fmt.Printf("Safe Reports count - %v", CountSafeReports(arr))
}
