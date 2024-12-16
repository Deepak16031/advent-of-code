package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ConvertStringToIntSlice(str string, sep string) []int {
	str = strings.Trim(str, " ")
	return ConvertToIntSlice(strings.Split(str, sep))
}

func ConvertToIntSlice(strArr []string) []int {
	var res []int
	for _, str := range strArr {
		converted, _ := strconv.Atoi(str)
		res = append(res, converted)
	}
	return res
}

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

func ReadAllAsArr(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []string
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}
	return arr
}

func RealAllLineAsStringUntilBreak(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var res []string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
		if line == "" {
			break
		}
		i++
	}
	return res
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

func ReadFileAs2DArr(filename, separator string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr [][]string
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, separator)
		arr = append(arr, lineSplit)
	}
	return arr
}

func MultiplyString(str1, str2 string) int {
	a, _ := strconv.Atoi(str1)
	b, _ := strconv.Atoi(str2)
	return a * b
}

func IndexOfEmptyLineInArr(rulesInputArr []string) int {
	indexOfEmptyLine := func(stringArr []string) int {
		for i, arr := range stringArr {
			if arr == "" {
				return i
			}
		}
		return -1
	}(rulesInputArr)
	return indexOfEmptyLine
}

func IsValid(a1, b1, ROWS, COLS int) bool {
	return a1 >= 0 && a1 < ROWS && b1 >= 0 && b1 < COLS
}

func CountDigits(n int) (count int) {
	count = 1
	for n > 9 {
		n = n / 10
		count++
	}
	return
}
