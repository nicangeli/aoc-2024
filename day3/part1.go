package main

import (
	"regexp"
	"strconv"
)

func part1(input string) int {
	// fmt.Println(input)
	// mul\(\d+,\d+\)

	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	allMatched := r.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range allMatched {
		left := match[1]
		right := match[2]

		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)

		sum += leftInt * rightInt
	}

	return sum

}
