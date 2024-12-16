package main

import (
	"fmt"
	"strconv"

	"advent-of-code-in-go/util"
)

type Cell [2]int

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day10/input.txt"
	arr := util.ReadFileAs2DArr(filename, "")

	fmt.Println(Solution(arr, true))
}

func Solution(arr [][]string, part2Enabled bool) int {
	var queue [][2]Cell
	for i, _ := range arr {
		for j, _ := range arr[i] {
			s := arr[i][j]
			if s == "0" {
				queue = append(queue, [2]Cell{{i, j}, {i, j}})
			}
		}
	}

	getCellVal := func(cell Cell) int {
		number, _ := strconv.Atoi(arr[cell[0]][cell[1]])
		return number
	}

	dirs := [][2]int{
		{0, -1},
		{-1, 0},
		{0, 1},
		{1, 0},
	}

	ROWS := len(arr[0])
	COLS := len(arr[1])

	for i := 0; i < 9; i++ {
		var temp [][2]Cell
		for _, cell := range queue {
			for _, dir := range dirs {
				nextCell := [2]int{cell[0][0] + dir[0], cell[0][1] + dir[1]}
				if util.IsValid(nextCell[0], nextCell[1], ROWS, COLS) && getCellVal(nextCell) == 1+getCellVal(cell[0]) {
					temp = append(temp, [2]Cell{nextCell, cell[1]})
				}
			}
		}
		queue = temp
	}

	if part2Enabled {
		return len(queue)
	}

	set := make(map[[2]Cell]bool)
	for _, cells := range queue {
		set[cells] = true
	}
	return len(set)
}
