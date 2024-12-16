package main

import (
	"fmt"
	"strconv"
)

type Pair [2]int

func main() {
	arr := []int{2701, 64945, 0, 9959979, 93, 781524, 620, 1}
	fmt.Println(Solution2(arr, 75))
}

func Solution2(arr []int, iterations int) int {
	cache := make(map[Pair]int)
	ans := 0
	for i := 0; i < len(arr); i++ {
		ans += solve(cache, arr[i], iterations)
	}
	return ans
}

func solve(cache map[Pair]int, val, iterations int) int {
	pair := Pair{val, iterations}
	if cache[pair] != 0 {
		return cache[pair]
	}
	if iterations == 0 {
		return 1
	}
	if val == 0 {
		cache[pair] = solve(cache, 1, iterations-1)
		return cache[pair]
	}
	strVal := strconv.Itoa(val)
	digits := len(strVal)
	if digits%2 == 0 {
		leftHalf, _ := strconv.Atoi(strVal[0 : digits/2])
		rightHalf, _ := strconv.Atoi(strVal[digits/2:])
		cache[pair] = solve(cache, leftHalf, iterations-1) + solve(cache, rightHalf, iterations-1)
	} else {
		cache[pair] = solve(cache, val*2024, iterations-1)
	}
	return cache[pair]
}
