package main

import (
	"fmt"

	"advent-of-code-in-go/util"
)

type Point [2]int

var visited = make(map[Point]bool)
var ROWS, COLS int
var dirs = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day12/input.txt"
	arr := util.ReadFileAs2DArr(filename, "")
	fmt.Println(Solution(arr))
}

func Solution(arr [][]string) int {
	cost := 0
	ROWS = len(arr)
	COLS = len(arr[0])
	for i, _ := range arr {
		for j := 0; j < COLS; j++ {
			area := 0
			peri := 0
			pointCost := dfs(arr, i, j, &area, &peri)
			cost += pointCost
		}
	}
	return cost
}

func dfs(arr [][]string, i int, j int, area *int, perimeter *int) int {
	point := Point{i, j}
	if visited[point] {
		return (*area) * (*perimeter)
	}
	visited[point] = true
	*area++
	for _, dir := range dirs {
		a1 := i + dir[0]
		b1 := j + dir[1]
		if !util.IsValid(a1, b1, ROWS, COLS) || arr[i][j] != arr[a1][b1] {
			*perimeter++
		}
		if util.IsValid(a1, b1, ROWS, COLS) && arr[i][j] == arr[a1][b1] {
			dfs(arr, a1, b1, area, perimeter)
		}
	}
	return (*area) * (*perimeter)
}
