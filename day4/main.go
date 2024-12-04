package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func main() {
	p1 := part1(input)

	fmt.Println(p1)

	p2 := part2(input)

	fmt.Println(p2)
}
