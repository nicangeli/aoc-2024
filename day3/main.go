package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

//go:embed inputEnabled.txt
var inputEnabled string

func main() {
	sum := part1(input)

	fmt.Println(sum)

	// 169140488
	// 178794710
	p2 := part2(input)
	fmt.Println(p2)
}
