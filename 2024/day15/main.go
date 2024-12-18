package main

import (
	"fmt"

	"advent-of-code-in-go/util"
)

func main() {
	robotPosiFilename := "/Users/deepak/GolandProjects/advent-of-code/2024/day15/input_1_posi.txt"
	arr := util.ReadFileAs2DArr(robotPosiFilename, "")

	robotMovesFile := "/Users/deepak/GolandProjects/advent-of-code/2024/day15/input_2_moves.txt"
	robotMoves := util.ReadFileAs2DArr(robotMovesFile, "")

	var robotMoves2 []string
	for _, rm := range robotMoves {
		robotMoves2 = append(robotMoves2, rm...)
	}
	Solution(arr, robotMoves2)
}

var dirs = map[string][2]int{
	">": {0, 1},
	"v": {1, 0},
	"<": {0, -1},
	"^": {-1, 0},
}

func Solution(arr [][]string, robotMovements []string) {

	start := -1
	end := -1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == "@" {
				start = i
				end = j
			}
		}
	}
	//fmt.Println("start")
	//for i := 0; i < len(arr); i++ {
	//	for j := 0; j < len(arr[0]); j++ {
	//	}
	//	fmt.Println(arr[i])
	//}
	//fmt.Println("--------------------------------")
	for _, move := range robotMovements {
		dir := dirs[move]
		start, end = moveRobot(start, end, dir, arr)
		//fmt.Println("After move ", move, dir, start, end)
		//for i := 0; i < len(arr); i++ {
		//	for j := 0; j < len(arr[0]); j++ {
		//	}
		//	fmt.Println(arr[i])
		//}
		//fmt.Println("--------------------------------")

	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == "O" {
				sum += j + 100*i
			}
		}
		fmt.Println(arr[i])
	}

	fmt.Println(sum)

}

func moveRobot(i, j int, dir [2]int, arr [][]string) (int, int) {
	r2 := i + dir[0]
	c2 := j + dir[1]

	for {
		if arr[r2][c2] == "#" {
			break
		}
		if arr[r2][c2] == "." {

			for r2 != i || c2 != j {
				r3 := r2 - dir[0]
				c3 := c2 - dir[1]
				arr[r2][c2] = arr[r3][c3]
				r2 = r3
				c2 = c3
			}
			arr[i][j] = "."
			i += dir[0]
			j += dir[1]
			break
		}
		r2 = r2 + dir[0]
		c2 = c2 + dir[1]
	}

	arr[i][j] = "@"
	return i, j
}
