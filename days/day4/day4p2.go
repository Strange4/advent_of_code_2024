package day4

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"strings"
)

type Coordinate struct {
	x, y int
}

func part2() {
	lines := util.ReadLines("./inputs/day4.txt")
	lineMap := make([][]string, len(lines))
	for i, line := range lines {
		lineMap[i] = strings.Split(line, "")
	}

	sum := 0

	for y := 0; y < len(lines); y++ {
		lineLength := len(lineMap[y])
		for x := 0; x < lineLength; x++ {
			// can't form an x
			if x > lineLength-len("MAS") || y > len(lines)-len("MAS") {
				break
			}
			tlIndeces := []Coordinate{{x, y}, {x + 1, y + 1}, {x + 2, y + 2}}
			trIndeces := []Coordinate{{x + 2, y}, {x + 1, y + 1}, {x, y + 2}}
			if checkIndeces(tlIndeces, lineMap) && checkIndeces(trIndeces, lineMap) {
				sum++
			}
		}
	}
	fmt.Println("Part 2: ", sum)
}

func checkIndeces(indeces []Coordinate, lineMap [][]string) bool {
	mas := ""
	for _, coordinate := range indeces {
		mas += lineMap[coordinate.y][coordinate.x]
	}
	return checkForMas(mas)
}

func checkForMas(masLike string) bool {
	if masLike == "MAS" || masLike == "SAM" {
		return true
	}
	return false
}
