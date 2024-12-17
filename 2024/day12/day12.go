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
			dfs(arr, i, j, &area, &peri)
			cost += area * peri
		}
	}
	return cost
}

func dfs(arr [][]string, i int, j int, area *int, perimeter *int) {
	point := Point{i, j}
	if visited[point] {
		return
	}
	visited[point] = true
	*area++

	good := func(dir [2]int) bool {
		r2 := i + dir[0]
		c2 := j + dir[1]
		return util.IsValid(r2, c2, ROWS, COLS) && arr[r2][c2] == arr[i][j]
	}

	for i, dir := range dirs {
		dir2 := dirs[(i+1)%4]
		if !good(dir) && !good(dir2) {
			*perimeter++
		}
		if good(dir) && good(dir2) && !good([2]int{dir[0] + dir2[0], dir[1] + dir2[1]}) {
			*perimeter++
		}
	}

	for _, dir := range dirs {
		a1 := i + dir[0]
		b1 := j + dir[1]
		if !visited[Point{a1, b1}] && good(dir) {
			dfs(arr, a1, b1, area, perimeter)
		}
	}
}
