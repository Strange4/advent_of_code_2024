package day10

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
)

func Run() {
	topoMap := util.ReadLines("./inputs/day10.txt")
	part1Sum := 0
	part2Sum := 0
	for y, line := range topoMap {
		for x, char := range line {
			if char == '0' {
				// no direction
				d := util.Direction(-1)
				finishes := map[util.Position]bool{}
				part2Sum += exploreMap(util.Position{X: x, Y: y}, d, finishes, topoMap)
				part1Sum += len(finishes)
			}
		}
	}
	fmt.Println("Part 1:", part1Sum)
	fmt.Println("Part 2:", part2Sum)
}

func exploreMap(position util.Position, noGo util.Direction, finishes map[util.Position]bool, topoMap []string) int {
	currentHeight := topoMap[position.Y][position.X]
	if currentHeight == '9' {
		finishes[position] = true
		return 1
	}

	width := len(topoMap[0])
	height := len(topoMap)
	foundHeights := 0
	directions := []util.Direction{util.Up, util.Down, util.Left, util.Right}
	for _, d := range directions {
		if d == noGo {
			continue
		}
		newPosition := position.MoveAndCopy(d)
		if !newPosition.InBounds(width, height) {
			continue
		}
		if topoMap[newPosition.Y][newPosition.X] == currentHeight+1 {
			foundHeights += exploreMap(newPosition, d.OppositeDirection(), finishes, topoMap)
		}
	}
	return foundHeights
}
