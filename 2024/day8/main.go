package main

import (
	"fmt"
	"log"

	"advent-of-code-in-go/util"
)

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day8/input.txt"
	arr := util.ReadFileAs2DArr(filename, "")
	fmt.Println("Number of anti nodes - ", Day8(arr, true))
}

func Day8(arr [][]string, part2Enabled bool) any {
	freqMap := make(map[string][][2]int)

	ROWS := len(arr)
	COLS := len(arr[0])
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			s := arr[i][j]
			freqMap[s] = append(freqMap[s], [2]int{i, j})
		}
	}
	delete(freqMap, ".")

	isValidAntiNode := func(node [2]int) bool {
		return node[0] >= 0 && node[0] < ROWS && node[1] >= 0 && node[1] < COLS
	}

	antiNodePosition := func(node1, node2 [2]int) [2]int {
		return [2]int{node1[0] + 2*(node2[0]-node1[0]), node1[1] + 2*(node2[1]-node1[1])}
	}

	antiNodes := make(map[[2]int]bool)
	for _, value := range freqMap {
		for i := 0; i < len(value); i++ {
			for j := 0; j < len(value); j++ {
				if i == j {
					continue
				}
				node1 := value[i]
				node2 := value[j]
				if part2Enabled {
					antiNodes[node1] = true
					antiNodes[node2] = true
				}
				k := 0
				for ; k < 10000000; k++ {
					antiNode := antiNodePosition(node1, node2)
					if !isValidAntiNode(antiNode) {
						break
					}
					antiNodes[antiNode] = true
					node1 = node2
					node2 = antiNode
					if !part2Enabled {
						break
					}
				}
				if k == 10000000 {
					log.Fatalf("infinite loop created")
				}
			}
		}
	}
	return len(antiNodes)
}
