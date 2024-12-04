package main

import (
	"strings"
)

type Crossword struct {
	data   []string
	width  int
	height int
}

type Coordinate struct {
	x int
	y int
}

func NewCrossword(input string) *Crossword {
	lines := strings.Split(input, "\n")

	return &Crossword{
		data:   lines,
		width:  len(lines),
		height: len(lines),
	}
}

func (cw *Crossword) getCoordinatesOfXmas(x int, y int) [][]Coordinate {
	var coords [][]Coordinate

	right := []Coordinate{Coordinate{x, y}, Coordinate{x + 1, y}, Coordinate{x + 2, y}, Coordinate{x + 3, y}}
	down := []Coordinate{Coordinate{x, y}, Coordinate{x, y + 1}, Coordinate{x, y + 2}, Coordinate{x, y + 3}}
	left := []Coordinate{Coordinate{x, y}, Coordinate{x - 1, y}, Coordinate{x - 2, y}, Coordinate{x - 3, y}}
	up := []Coordinate{Coordinate{x, y}, Coordinate{x, y - 1}, Coordinate{x, y - 2}, Coordinate{x, y - 3}}
	topRight := []Coordinate{Coordinate{x, y}, Coordinate{x + 1, y - 1}, Coordinate{x + 2, y - 2}, Coordinate{x + 3, y - 3}}
	bottomRight := []Coordinate{Coordinate{x, y}, Coordinate{x + 1, y + 1}, Coordinate{x + 2, y + 2}, Coordinate{x + 3, y + 3}}
	bottomLeft := []Coordinate{Coordinate{x, y}, Coordinate{x - 1, y + 1}, Coordinate{x - 2, y + 2}, Coordinate{x - 3, y + 3}}
	topLeft := []Coordinate{Coordinate{x, y}, Coordinate{x - 1, y - 1}, Coordinate{x - 2, y - 2}, Coordinate{x - 3, y - 3}}

	coords = append(coords, right)
	coords = append(coords, down)
	coords = append(coords, left)
	coords = append(coords, up)
	coords = append(coords, topRight)
	coords = append(coords, bottomRight)
	coords = append(coords, bottomLeft)
	coords = append(coords, topLeft)
	return coords
}

func (cw *Crossword) getXmases(x int, y int) int {
	xmasCount := 0
	coords := cw.getCoordinatesOfXmas(x, y)

	for _, coordGroup := range coords {
		str := ""
		for _, coord := range coordGroup {
			if (coord.x < cw.width) && (coord.y < cw.height) && (coord.y >= 0 && coord.x >= 0) {
				str += string(cw.data[coord.y][coord.x])
			}
		}
		if str == "XMAS" {
			xmasCount++
		}
	}

	return xmasCount

}

func part1(input string) int {
	crossword := NewCrossword(input)

	xmasCount := 0

	for y := range crossword.data {
		for x := range crossword.data[y] {
			char := string(crossword.data[y][x])
			if char == "X" {
				xmases := crossword.getXmases(x, y)
				xmasCount += xmases
			}
		}
	}

	return xmasCount
}
