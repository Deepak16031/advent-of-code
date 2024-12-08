package main

import (
	"fmt"

	"advent-of-code-in-go/util"
)

func main() {
	filename := "/Users/deepak/GolandProjects/advent-of-code/2024/day5/input.txt"
	rulesInputArr := util.ReadAllAsArr(filename)
	indexOfEmptyLine := util.IndexOfEmptyLineInArr(rulesInputArr)
	rules := CreateRules(rulesInputArr[:indexOfEmptyLine])
	updates := rulesInputArr[indexOfEmptyLine:]

	part := true
	sum := SumOfMiddleOfValidUpdates(updates, rules, part)
	fmt.Printf("sum of all updates with part2 ? %v is %v\n", part, sum)
	sum = SumOfMiddleOfValidUpdates(updates, rules, !part)
	fmt.Printf("sum of all updates with part2 ? %v is %v\n", !part, sum)
}
