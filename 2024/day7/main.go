package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code-in-go/util"
)

func main() {

	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day7/input.txt"
	stringInputArr := util.ReadAllAsArr(filename)
	var arr [][]int

	for i, s := range stringInputArr {
		split := strings.Split(s, ":")
		result, _ := strconv.Atoi(split[0])
		values := util.ConvertStringToIntSlice(split[1], " ")
		arr = append(arr, []int{})
		arr[i] = append(arr[i], result)
		arr[i] = append(arr[i], values...)
	}
	fmt.Println("Output - ", Part1(arr))
}

func Part1(input [][]int) int {
	sum := 0
	var ch = make(chan int)
	for _, test := range input {
		go func(arr []int) {
			ch <- solve(arr)
		}(test)
	}
	for i := 0; i < len(input); i++ {
		sum += <-ch
	}
	return sum
}

func solve(input []int) int {
	result := input[0]
	values := input[1:]
	if solveUtil(values, 1, values[0], result) {
		return result
	}
	return 0
}

func solveUtil(values []int, indx, currValue, target int) bool {
	if indx == len(values) {
		return currValue == target
	}
	return solveUtil(values, indx+1, currValue+values[indx], target) ||
		solveUtil(values, indx+1, currValue*values[indx], target)

}
