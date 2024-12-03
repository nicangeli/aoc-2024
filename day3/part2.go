package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func tokenize(input string) []string {
	var multiplications []string
	i := 0
	for i < len(input) {
		char := string(input[i])
		if char == "m" {
			token := input[i : i+4]
			if token == "mul(" {
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
						multiplications = append(multiplications, inner)
						i = i + len(inner) // 1 is the final )
						break
					} else {
						i = i + len(inner) // 1 is the final char we tried to read
						break
					}
				}
			}
		}
		i++

	}

	return multiplications
}

func part2(input string) {
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

	fmt.Println(sum)
}
