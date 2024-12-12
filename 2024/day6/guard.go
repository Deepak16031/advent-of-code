package main

import (
	"strings"

	"advent-of-code-in-go/util"
)

func FindVisitedPositions(arr []string) int {

	ROWS := len(arr)
	COLS := len(arr[0])
	var guardPosi [2]int
	for i, _ := range arr {
		str := arr[i]
		indx := strings.Index(str, "^")
		if indx != -1 {
			guardPosi[0] = i
			guardPosi[1] = indx
		}
	}

	dirs := map[[2]int][2]int{
		{-1, 0}: {0, 1},  // up to right
		{0, 1}:  {1, 0},  // right to down
		{1, 0}:  {0, -1}, // down to left
		{0, -1}: {-1, 0}, // left to up
	}
	res := make(map[[2]int]bool)
	currDir := [2]int{-1, 0}
	for util.IsValid(guardPosi[0], guardPosi[1], ROWS, COLS) {
		res[[2]int{guardPosi[0], guardPosi[1]}] = true
		newPosi := nextPosi(guardPosi, currDir)
		if util.IsValid(newPosi[0], newPosi[1], ROWS, COLS) &&
			arr[newPosi[0]][newPosi[1]:newPosi[1]+1] == "#" {
			currDir = dirs[currDir]
		}
		guardPosi = nextPosi(guardPosi, currDir)
	}
	return len(res)
}

func Part2(arr [][]string) int {
	ROWS := len(arr)
	COLS := len(arr[0])

	var start [2]int
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if arr[i][j] == "^" {
				start = [2]int{i, j}
				arr[i][j] = "."
			}
		}
	}

	dirs := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	res := 0
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if start != [2]int{i, j} && arr[i][j] == "." {
				arr[i][j] = "#"
				guardPosi := [2]int{start[0], start[1]}
				currDir := 0
				var visit [][3]int
				for {
					visit = append(visit, [3]int{guardPosi[0], guardPosi[1], currDir})
					if len(visit) > 6*ROWS*COLS {
						//fmt.Println("loop obstruction (", i, ",", j, ")")
						res++
						break
					}
					newPosi := nextPosi(guardPosi, dirs[currDir])
					if !util.IsValid(newPosi[0], newPosi[1], ROWS, COLS) {
						break
					} else if arr[newPosi[0]][newPosi[1]] == "." {
						guardPosi = newPosi
					} else {
						currDir = (currDir + 1) % 4
					}
				}
				arr[i][j] = "."
			}
		}
	}
	return res
}

func nextPosi(posi, dir [2]int) [2]int {
	a := posi[0] + dir[0]
	b := posi[1] + dir[1]
	return [2]int{a, b}
}
