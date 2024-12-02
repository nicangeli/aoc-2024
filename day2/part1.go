package main

// 413

import (
	"math"
	"strconv"
	"strings"
)

func parseReport(report string) ([]int, error) {
	var levels []int

	for _, level := range strings.Split(report, " ") {
		num, err := strconv.Atoi(level)
		if err != nil {
			return nil, err
		}
		levels = append(levels, num)
	}

	return levels, nil
}

func isReportSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	isIncrementing := levels[0] < levels[1]
	for i, current := range levels {
		if i == len(levels)-1 {
			return true
		}

		next := levels[i+1]
		if isIncrementing && (next < current) {
			return false
		}

		if !isIncrementing && (next > current) {
			return false
		}

		diff := int(math.Abs(float64(next - current)))

		if diff > 3 || diff < 1 {
			return false
		}

	}

	return true
}

func part1(file string) (int, error) {
	safe := 0
	lines := strings.Split(file, "\n")

	for _, line := range lines {
		levels, err := parseReport(line)
		if err != nil {
			return 0, err
		}

		if isReportSafe(levels) {
			safe += 1
		}

	}

	return safe, nil
}
