package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2() int {
	fmt.Println("Advent of code day 1")

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	listA := make([]int, 100)
	occurrences := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")

		if len(numbers) != 2 {
			log.Fatal("Bad parsing")
		}

		a := numbers[0]
		b := numbers[1]

		aInt, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal("Bad parsing")
		}

		bInt, err := strconv.Atoi(b)
		if err != nil {
			log.Fatal("Bad parsing")
		}

		listA = append(listA, aInt)
		occurrences[bInt]++
	}

	sum := 0

	for _, a := range listA {
		score := a * occurrences[a]

		sum += score
	}

	return sum
}
