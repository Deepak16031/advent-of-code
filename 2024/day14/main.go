package main

import (
	"fmt"

	"advent-of-code-in-go/util"
)

type Coordinate [2]int

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day14/input.txt"
	arr := util.ReadNumbersTo2DArrWithSep(filename, ",")

	ROWS := 103
	COLS := 101
	var grid [103][101]string

	cache := make(map[Coordinate]bool)

	var q1, q2, q3, q4 int

	for seconds := 1; seconds < 10000; seconds++ {
		cache = make(map[Coordinate]bool)
		grid = [103][101]string{}
		for i, _ := range arr {
			a := arr[i][0]
			b := arr[i][1]
			c := arr[i][2]
			d := arr[i][3]

			x, y := updatePosisition(a, b, c, d, seconds, ROWS, COLS)
			cache[Coordinate{x, y}] = true
			grid[x][y] = "*"

			if x == ROWS/2 || y == COLS/2 {
				continue
			}
			if x < ROWS/2 {
				if y < COLS/2 {
					q1++
				} else {
					q3++
				}
			} else {
				if y < COLS/2 {
					q2++
				} else {
					q4++
				}
			}
		}
		if len(cache) == len(arr) {
			fmt.Println(seconds)
			for i := 0; i < 103; i++ {
				for j := 0; j < 101; j++ {
					if grid[i][j] == "" {
						grid[i][j] = "."
					}
					fmt.Print(grid[i][j])
				}
				fmt.Println()
			}
			break
		}
	}

}

func updatePosisition(a, b, c, d, seconds, ROWS, COLS int) (int, int) {
	a = a + ((c * seconds) % COLS)
	b = b + ((d * seconds) % ROWS)
	a = (a + COLS) % COLS
	b = (b + ROWS) % ROWS
	return b, a
}
