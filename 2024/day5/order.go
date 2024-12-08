package main

import (
	"log"
	"strings"

	"advent-of-code-in-go/util"
)

func IsCorrectOrder(arr []int, rules map[int]map[int]bool) bool {
	for i, val := range arr {
		subArr := arr[i+1:]
		for _, subVal := range subArr {
			if rules[subVal][val] {
				return false
			}
		}
	}
	return true
}

func CreateRules(rulesInputArr []string) map[int]map[int]bool {
	rules := make(map[int]map[int]bool)
	for i, rule := range rulesInputArr {
		ruleSplit := strings.Split(rule, "|")
		numberArr := util.ConvertToIntSlice(ruleSplit)
		if len(ruleSplit) == 1 {
			log.Fatalf("ruplesSplit %v prev %v", ruleSplit, rulesInputArr[i-1])
		}
		if rules[numberArr[0]] == nil {
			rules[numberArr[0]] = make(map[int]bool)
		}
		rules[numberArr[0]][numberArr[1]] = true
	}
	return rules
}

func SumOfMiddleOfValidUpdates(updates []string, rules map[int]map[int]bool) int {
	sum := 0
	for _, update := range updates {
		updateIntArr := util.ConvertToIntSlice(strings.Split(update, ","))
		if IsCorrectOrder(updateIntArr, rules) {
			sum += updateIntArr[len(updateIntArr)/2]
		}
	}
	return sum
}
