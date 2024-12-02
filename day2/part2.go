package main

import (
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

func isOnePermutationSafe(reports []int) bool {
	for _, permutation := range permutations(reports) {
		if isReportSafe(permutation) {
			return true
		}
	}

	return false
}

func part2(file string) (int, error) {
	safe := 0
	lines := strings.Split(file, "\n")

	for _, line := range lines {
		levels, err := parseReport(line)
		if err != nil {
			return 0, err
		}

		isSafe := isReportSafe(levels)

		if isSafe || isOnePermutationSafe(levels) {
			safe += 1
		}

	}

	return safe, nil
}
