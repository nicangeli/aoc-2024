package main

import (
	"strings"
)

func part2(input string) int {
	dos := strings.Split(input, "do()")

	sum := 0
	for _, do := range dos {
		donots := strings.Split(do, "don't()")

		goodDos := donots[0]
		sum += part1(goodDos)
	}

	return sum
}
