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

	sum := SumOfMiddleOfValidUpdates(updates, rules)
	fmt.Println("sum of all correct updates - ", sum)
}
