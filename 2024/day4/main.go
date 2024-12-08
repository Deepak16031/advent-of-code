package main

import (
	"fmt"

	"advent-of-code-in-go/util"
)

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day4/input.txt"
	arr := util.ReadAllAsArr(filename)
	fmt.Printf("Word Count - %v\n", FindWordCount(arr, "XMAS"))
	arr = util.ReadAllAsArr(filename)
	fmt.Printf("X-MAS count %v", FindWordCountPart2(arr))
}
