package day8

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
)

func Run() {
	part1()
	part2()
}

func part1() {
	lines := util.ReadLines("./inputs/day8.txt")
	width := len(lines)
	height := len(lines[0])
	antennas := getAntennas(lines)

	antinodes := make(map[util.Position]bool, 0)

	for _, positions := range antennas {
		for i, p1 := range positions {
			for j := i; j < len(positions); j++ {
				p2 := positions[j]
				if i == j {
					continue
				}
				xDistance := p1.X - p2.X
				yDistance := p1.Y - p2.Y

				a1 := util.Position{X: p1.X + xDistance, Y: p1.Y + yDistance}
				a2 := util.Position{X: p2.X - xDistance, Y: p2.Y - yDistance}
				if positionInBounds(a1, width, height) {
					antinodes[a1] = true
				}
				if positionInBounds(a2, width, height) {
					antinodes[a2] = true
				}
			}
		}
	}
	fmt.Println("Part 1: ", len(antinodes))
}

func part2() {
	lines := util.ReadLines("./inputs/day8.txt")

	width := len(lines)
	height := len(lines[0])

	antennas := getAntennas(lines)
	antinodes := make(map[util.Position]bool, 0)

	for _, positions := range antennas {
		for i, p1 := range positions {
			for j := i; j < len(positions); j++ {
				p2 := positions[j]
				if i == j {
					continue
				}
				a1 := util.Position{X: p1.X, Y: p1.Y}
				a2 := util.Position{X: p2.X, Y: p2.Y}

				xDistance := p1.X - p2.X
				yDistance := p1.Y - p2.Y
				for positionInBounds(a1, width, height) {
					antinodes[a1] = true
					a1.X += xDistance
					a1.Y += yDistance
				}
				for positionInBounds(a2, width, height) {
					antinodes[a2] = true
					a2.X -= xDistance
					a2.Y -= yDistance
				}
			}
		}
	}

	fmt.Println("Part 2:", len(antinodes))
}

func positionInBounds(p util.Position, width, height int) bool {
	return (p.X >= 0) && (p.X <= width-1) && (p.Y >= 0) && (p.Y <= height-1)
}

func getAntennas(lines []string) map[rune][]util.Position {
	antennas := make(map[rune][]util.Position, 0)
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				pos := util.Position{X: x, Y: y}
				antennas[char] = append(antennas[char], pos)
			}
		}
	}
	return antennas
}
