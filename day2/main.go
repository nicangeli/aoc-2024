package main

import (
	_ "embed"
	"fmt"
	"log"
)

//go:embed input.txt
var file string

func main() {
	fmt.Println("Lets go day 2")

	safe, err := part1(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", safe)

	dampenedSafe, err := part2(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 2: ", dampenedSafe)
}
