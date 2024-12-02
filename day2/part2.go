package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func permutations(levels []int) [][]int {
	var ps [][]int

	for i := 0; i < len(levels); i++ {
		var permutation []int
		for j := 0; j < len(levels); j++ {
			if j == i {
				continue
			}
			permutation = append(permutation, levels[j])
		}
		ps = append(ps, permutation)
	}

	return ps
}

func isSafe(levels []int) bool {
	isSafe := true
	isIncrementing := levels[0] < levels[1]
	for i, current := range levels {
		if (i + 1) == len(levels) {
			break
		}

		next := levels[i+1]
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

	return isSafe
}

func part2() (int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)

	safe := 0

	for scanner.Scan() {
		line := scanner.Text()

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

		isSfe := isSafe(levelsInt)

		isPermutationSafe := false
		if !isSfe {
			for _, permutation := range permutations(levelsInt) {
				if isSafe(permutation) {
					isPermutationSafe = true
					break
				}
			}
		}

		if isSfe || isPermutationSafe {
			safe += 1
		}

	}

	return safe, nil
}
