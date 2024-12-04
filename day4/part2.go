package main

type XGroup struct {
	left    []Coordinate
	right   []Coordinate
	topLeft Coordinate
}

func (cw *Crossword) getCoordinatesOfMas(x int, y int) XGroup {
	rightDown := []Coordinate{Coordinate{x, y}, Coordinate{x + 1, y + 1}, Coordinate{x + 2, y + 2}}
	leftDown := []Coordinate{Coordinate{x + 2, y}, Coordinate{x + 1, y + 1}, Coordinate{x, y + 2}}

	return XGroup{left: rightDown, right: leftDown, topLeft: Coordinate{x, y}}
}

func (cw *Crossword) getMases(x int, y int) int {
	masCount := 0
	xgroupCoords := cw.getCoordinatesOfMas(x, y)

	leftStr := ""
	for _, coord := range xgroupCoords.left {
		if (coord.x < cw.width) && (coord.y < cw.height) && (coord.y >= 0 && coord.x >= 0) {
			leftStr += string(cw.data[coord.y][coord.x])
		}
	}

	rightStr := ""
	for _, coord := range xgroupCoords.right {
		if (coord.x < cw.width) && (coord.y < cw.height) && (coord.y >= 0 && coord.x >= 0) {
			rightStr += string(cw.data[coord.y][coord.x])
		}
	}
	if (leftStr == "MAS" || leftStr == "SAM") && (rightStr == "MAS" || rightStr == "SAM") {
		masCount++
	}

	return masCount
}

func part2(input string) int {
	crossword := NewCrossword(input)

	masCount := 0

	for y := range crossword.data {
		for x := range crossword.data[y] {
			char := string(crossword.data[y][x])
			if char == "M" || char == "S" {
				mases := crossword.getMases(x, y)
				masCount += mases
			}
		}
	}

	return masCount
}
