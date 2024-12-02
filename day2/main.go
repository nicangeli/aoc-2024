package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Lets go day 2")

	safe, err := part1()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", safe)

	dampenedSafe, err := part2()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 2: ", dampenedSafe)
}
