package main

import (
	"fmt"
	"math"

	"advent-of-code-in-go/util"
)

type Coordinate [2]int

var cache = make(map[Coordinate]int)

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day13/input.txt"
	arr := util.ReadNumbersTo2DArrWithSep(filename, ",")

	totalTokens := 0
	for i := 0; i < len(arr); i += 3 {
		tokens := Solution(arr[i][0], arr[i+1][0], arr[i][1], arr[i+1][1], 10000000000000+arr[i+2][0], 10000000000000+arr[i+2][1], true)
		totalTokens += tokens
	}
	fmt.Println(totalTokens)
}

func Solution(a, b, c, d, x, y int, part2Enabled bool) int {
	xSolve, ySolve := solveLinAlg(a, b, c, d, x, y)
	if !part2Enabled {
		if xSolve > 100 || ySolve > 100 {
			return 0
		}
	}
	if math.Round(xSolve) == xSolve && math.Round(ySolve) == ySolve {
		return 3*int(math.Round(xSolve)) + int(math.Round(ySolve))
	}
	return 0
}

func det(a, b, c, d int) int {
	return a*d - b*c
}

func solveLinAlg(a, b, c, d, x, y int) (float64, float64) {
	det1 := det(a, b, c, d)

	xdet := float64(det(x, b, y, d))
	ydet := float64(det(a, x, c, y))
	xSolve := xdet / float64(det1)
	ySolve := ydet / float64(det1)

	return xSolve, ySolve
}
