package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() int {
	fmt.Println("Advent of code day 1")

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	listA := make([]int, 100)
	listB := make([]int, 100)

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
		listB = append(listB, bInt)

	}

	sort.Ints(listA)
	sort.Ints(listB)

	if len(listA) != len(listB) {
		log.Fatal("not same length")
	}

	sum := 0

	for i := range listA {
		diff := math.Abs(float64(listA[i] - listB[i]))

		sum += int(diff)
	}

	return sum
}
