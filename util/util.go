package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string, separator string) ([]int, []int) {
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
		split := strings.Split(line, separator)
		number, _ := strconv.Atoi(split[0])
		arr1 = append(arr1, number)
		number, _ = strconv.Atoi(split[1])
		arr2 = append(arr2, number)
	}

	return arr1, arr2
}

func ReadAllAsArr(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr [][]string
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, "")
		arr = append(arr, lineArr)
	}
	return arr
}

func ReadNumbersTo2DArr(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var res [][]int
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		var convertedToInt []int
		for _, s := range split {
			number, _ := strconv.Atoi(s)
			convertedToInt = append(convertedToInt, number)
		}
		res = append(res, convertedToInt)
	}
	return res
}

func MultiplyString(str1, str2 string) int {
	a, _ := strconv.Atoi(str1)
	b, _ := strconv.Atoi(str2)
	return a * b
}
