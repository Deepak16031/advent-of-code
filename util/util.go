package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr1 []int
	var arr2 []int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		number, _ := strconv.Atoi(split[0])
		arr1 = append(arr1, number)
		number, _ = strconv.Atoi(split[1])
		arr2 = append(arr2, number)
	}

	return arr1, arr2
}
