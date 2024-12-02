package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

func part1(file string) (int, error) {
	safe := 0
	lines := strings.Split(file, "\n")

	for _, line := range lines {
		levelsStr := strings.Split(line, " ")
		levelsInt := make([]int, len(levelsStr))

		for i, levelStr := range levelsStr {
			num, err := strconv.Atoi(levelStr)
			if err != nil {
				return 0, err
			}
			levelsInt[i] = num
		}

		if len(levelsInt) < 2 {
			safe += 1
			break
		}

		isSafe := true
		isIncrementing := levelsInt[0] < levelsInt[1]
		for i, current := range levelsInt {
			if (i + 1) == len(levelsInt) {
				break
			}

			next := levelsInt[i+1]
			if next == current {
				isSafe = false
				break
			}

			if isIncrementing && (next < current) {
				isSafe = false
				break
			}

			if !isIncrementing && (next > current) {
				isSafe = false
				break
			}

			diff := int(math.Abs(float64(next - current)))

			if diff > 3 {
				isSafe = false
				break
			}

		}

		if isSafe {
			safe += 1
		}

	}

	return safe, nil
}
