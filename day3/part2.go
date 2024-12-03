package main

import (
	"slices"
	"strconv"
	"strings"
)

func tokenize(input string) []string {
	var multiplications []string
	i := 0
	enabled := true
	for i < len(input) {
		if i+7 > len(input) {
			break
		}

		doToken := input[i : i+4]
		if doToken == "do()" {
			enabled = true
			i = i + len(doToken)
		}

		donotToken := input[i : i+7]
		if donotToken == "don't()" {
			enabled = false
			i = i + len(donotToken)
		}

		mulToken := input[i : i+4]
		if mulToken == "mul(" {
			i = i + 4
			inner := ""

			validNums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
			innerCount := 0
			for {
				char := string(input[i+innerCount])

				if slices.Contains(validNums, char) {
					inner += char
					innerCount++
					continue
				} else if char == "," && !strings.Contains(inner, ",") {
					inner += char
					innerCount++
					continue
				} else if char == ")" && strings.Contains(inner, ",") && strings.Index(inner, ",") != len(inner)-1 {
					// successfully parsed
					if enabled {
						multiplications = append(multiplications, inner)
					}
					i = i + len(inner) // 1 is the final )
					break
				} else {
					i = i + len(inner) // 1 is the final char we tried to read
					break
				}
			}
		}
		i++

	}

	return multiplications
}

func part2(input string) int {
	tokens := tokenize(input)

	sum := 0

	for _, elem := range tokens {
		parts := strings.Split(elem, ",")
		left := parts[0]
		right := parts[1]

		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)

		sum += leftInt * rightInt

	}

	return sum
}
