package main

import (
	"fmt"
)

const LOG_ENABLED = false

func FindWordCount(arr [][]string, word string) int {
	ROWS := len(arr)
	COLS := len(arr[0])
	isValid := func(a1, b1 int) bool {
		return a1 >= 0 && a1 < ROWS && b1 >= 0 && b1 < COLS
	}

	log := make(map[[2]int]bool)

	formWord := func(dir [2]int, x, y, size int) (string, [][2]int) {
		newWord := ""
		var indexes [][2]int
		for i := 0; i < size; i++ {
			if isValid(x, y) {
				newWord = newWord + arr[x][y]
				indexes = append(indexes, [2]int{x, y})
			} else {
				break
			}
			x = x + dir[0]
			y = y + dir[1]
		}
		return newWord, indexes
	}

	posis := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	res := 0

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			for k := 0; k < len(posis); k++ {
				dir := posis[k]
				wordInDir, indexes := formWord(dir, i, j, len(word))
				if wordInDir == word {
					for i := 0; i < len(indexes); i++ {
						indx := indexes[i]
						log[[2]int{indx[0], indx[1]}] = true
					}
					res++
				}
			}
		}
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if !log[[2]int{i, j}] {
				arr[i][j] = "."
			}
		}
		Log(arr[i])
	}
	return res
}

func FindWordCountPart2(arr [][]string) int {
	ROWS := len(arr)
	COLS := len(arr[0])
	isValid := func(a1, b1 int) bool {
		return a1 >= 0 && a1 < ROWS && b1 >= 0 && b1 < COLS
	}

	getLefDiagonalWord := func(i, j int) string {
		if isValid(i-1, j-1) && isValid(i+1, j+1) {
			return fmt.Sprintf("%v%v%v", arr[i-1][j-1], arr[i][j], arr[i+1][j+1])
		}
		return ""
	}

	getRightDiagonalWord := func(i, j int) string {
		if isValid(i+1, j-1) && isValid(i-1, j+1) {
			return fmt.Sprintf("%v%v%v", arr[i-1][j+1], arr[i][j], arr[i+1][j-1])
		}
		return ""
	}

	res := 0

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			sortedLeftDiag := getLefDiagonalWord(i, j)
			sortedRightDiag := getRightDiagonalWord(i, j)
			if !(sortedLeftDiag == "MAS") && !(sortedLeftDiag == "SAM") {
				continue
			}
			if !(sortedRightDiag == "MAS") && !(sortedRightDiag == "SAM") {
				continue
			}
			res++
		}
	}

	return res
}

func Log(s []string) {
	if LOG_ENABLED {
		fmt.Println(s)
	}
}
