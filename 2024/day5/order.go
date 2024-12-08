package main

import (
	"log"
	"slices"
	"strings"

	"advent-of-code-in-go/util"
)

type Rules map[int]map[int]bool

func IsCorrectOrder(arr []int, rules Rules) bool {
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

func CreateRules(rulesInputArr []string) Rules {
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

func SumOfMiddleOfValidUpdates(updates []string, rules Rules, part2 bool) int {
	sum := 0
	for _, update := range updates {
		updateIntArr := util.ConvertToIntSlice(strings.Split(update, ","))
		isOrderCorrect := IsCorrectOrder(updateIntArr, rules)
		if !part2 && isOrderCorrect {
			sum += updateIntArr[len(updateIntArr)/2]
		} else if part2 && !isOrderCorrect {
			correctTheOrder(updateIntArr, rules)
			if !IsCorrectOrder(updateIntArr, rules) {
				log.Fatalf("order is not corrected")
			}
			sum += updateIntArr[len(updateIntArr)/2]
		}
	}
	return sum
}

func correctTheOrder(arr []int, rules Rules) {
	slices.SortFunc(arr, func(a, b int) int {
		if !rules[a][b] {
			return 1
		}
		return -1
	})
}
