package main

import (
	"fmt"
	"log"
	"strconv"

	"advent-of-code-in-go/util"
)

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day9/input.txt"
	input := util.ReadFileAs2DArr(filename, "")
	arr := input[0]
	fmt.Println("Checksum - ", Solution(arr))
}

func Solution(input []string) int {
	image := DiskImage(input)
	image = Compaction(image)
	return checksum(image)
}

func checksum(image []string) (sum int) {
	for i, val := range image {
		if val == "." {
			break
		}
		fileId, _ := strconv.Atoi(val)
		sum += i * fileId
	}
	return sum
}

func DiskImage(arr []string) []string {
	var createDiskArr []string

	fileId := 0
	for i := 0; i < len(arr); i++ {
		val, _ := strconv.Atoi(arr[i])
		for j := 0; j < val; j++ {
			createDiskArr = append(createDiskArr, strconv.Itoa(fileId))
		}
		fileId++
		i++
		if i >= len(arr) {
			break
		}
		val, _ = strconv.Atoi(arr[i])
		for j := 0; j < val; j++ {
			createDiskArr = append(createDiskArr, ".")
		}
	}
	return createDiskArr
}

func Compaction(arr []string) []string {
	i := 0
	j := len(arr) - 1
	k := 0
	for k < 900000 {
		i = findFirFreeSpace(arr, i)
		j = findFirstUsedSpace(arr, j)

		if i == -1 || j == -1 || i > j {
			return arr
		}
		arr[i] = arr[j]
		arr[j] = "."
		i++
		j--

		k++
	}
	if k == 9000000 {
		log.Fatalf("infite loop created")
	}
	return arr
}

func findFirstUsedSpace(arr []string, j int) int {
	for i := j; i > -1; i-- {
		if arr[i] != "." {
			return i
		}
	}
	return -1
}

func findFirFreeSpace(arr []string, i int) int {
	for j := i; j < len(arr); j++ {
		if arr[j] == "." {
			return j
		}
	}
	return -1
}
