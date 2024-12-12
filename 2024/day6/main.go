package main

import (
	"fmt"
	"strings"

	"advent-of-code-in-go/util"
)

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day6/input.txt"
	arr := util.ReadAllAsArr(filename)
	unqiuePosisCount := FindVisitedPositions(arr)

	fmt.Println("Unique positions visited by guard - ", unqiuePosisCount)
	arr2 := make([][]string, len(arr))
	for i := 0; i < len(arr); i++ {
		split := strings.Split(arr[i], "")
		arr2[i] = split
	}
	obstructions := Part2(arr2)
	fmt.Println("Obstruction count for infinte loop - ", obstructions)
}
