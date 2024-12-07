package day3

import (
	"regexp"
	"strings"

	"advent-of-code-in-go/util"
)

func MulItOver(str string) int {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	allMatches := regex.FindAllString(str, -1)

	numberRegex := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	sum := 0
	for _, match := range allMatches {
		numberString := numberRegex.FindString(match)
		numbersSplit := strings.Split(numberString, ",")
		sum += util.MultiplyString(numbersSplit[0], numbersSplit[1])
	}
	return sum

}

func MulItOverPart2(str string) int {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	allMatches := regex.FindAllString(str, -1)

	numberRegex := regexp.MustCompile(`\d{1,3},\d{1,3}`)
	sum := 0
	enabled := true
	for _, match := range allMatches {
		if match == "don't()" {
			enabled = false
			continue
		} else if match == "do()" {
			enabled = true
			continue
		}
		if !enabled {
			continue
		}
		numberString := numberRegex.FindString(match)
		numbersSplit := strings.Split(numberString, ",")
		sum += util.MultiplyString(numbersSplit[0], numbersSplit[1])
	}
	return sum
}
