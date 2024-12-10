package day6

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"slices"
)

func Run() {
	part1And2()
}

type Direction int

const (
	Up Direction = iota
	Right
	Left
	Down
)

type Guard = struct {
	x, y      int
	direction Direction
}

type Area = struct {
	hasBeenStepedOn bool
	direction       []Direction // if it was stepped on, what direction where you going?
}

func part1And2() {
	lines := util.ReadLines("./inputs/day6.txt")

	var currentPos Guard

	obstacles := make([][]bool, len(lines))
	visitedAreas := make([][]Area, len(lines))
	for y, line := range lines {
		obstacles[y] = make([]bool, len(line))
		visitedAreas[y] = make([]Area, len(line))

		for x, char := range line {
			if char == '#' {
				obstacles[y][x] = true
			} else if char == '^' {
				currentPos = Guard{x, y, Up}
			}
		}
	}

	possibleLoops := walkTheGuard(currentPos, currentPos, visitedAreas, obstacles, 1)

	sum := 1 // count the ending step
	for y := 0; y < len(visitedAreas); y++ {
		for x := 0; x < len(visitedAreas[y]); x++ {
			if visitedAreas[y][x].hasBeenStepedOn {
				sum++
			}
		}
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", possibleLoops)
}

func walkTheGuard(initialPosition Guard, currentPos Guard, visitedAreas [][]Area, obstacles [][]bool, obstaclesToPlace int) int {
	possibleLoops := 0
	for {
		for ; !obstacles[currentPos.y][currentPos.x] && !isAtEdge(currentPos, len(obstacles[0]), len(obstacles)); incrementDirection(&currentPos) {
			area := visitedAreas[currentPos.y][currentPos.x]
			if area.hasBeenStepedOn && slices.Contains(area.direction, currentPos.direction) {
				return possibleLoops + 1
			}

			if obstaclesToPlace > 0 {
				newLoop := canCreateLoop(initialPosition, currentPos, visitedAreas, obstacles, obstaclesToPlace-1)
				possibleLoops += newLoop
			}
			visitedAreas[currentPos.y][currentPos.x].hasBeenStepedOn = true
			if visitedAreas[currentPos.y][currentPos.x].direction == nil {
				visitedAreas[currentPos.y][currentPos.x].direction = make([]Direction, 0)
			}
			visitedAreas[currentPos.y][currentPos.x].direction = append(visitedAreas[currentPos.y][currentPos.x].direction, currentPos.direction)

		}
		if obstacles[currentPos.y][currentPos.x] {
			decrementDirection(&currentPos)
		} else if isAtEdge(currentPos, len(obstacles[0]), len(obstacles)) {
			return possibleLoops
		} else {
			panic("Exited loop without wanting to")
		}
		turnGuard(&currentPos)
	}
}

func incrementDirection(g *Guard) {
	if g.direction == Right {
		g.x++
	} else if g.direction == Left {
		g.x--
	} else if g.direction == Up {
		g.y--
	} else if g.direction == Down {
		g.y++
	}
}
func decrementDirection(g *Guard) {
	if g.direction == Right {
		g.x--
	} else if g.direction == Left {
		g.x++
	} else if g.direction == Up {
		g.y++
	} else if g.direction == Down {
		g.y--
	}
}

func turnGuard(g *Guard) {
	if g.direction == Right {
		g.direction = Down
	} else if g.direction == Left {
		g.direction = Up
	} else if g.direction == Up {
		g.direction = Right
	} else if g.direction == Down {
		g.direction = Left
	}
}

func canCreateLoop(initialPosition Guard, guard Guard, areas [][]Area, obstacles [][]bool, obstaclesToPlace int) int {
	// checking if we create an obstacle in front would create a loop
	incrementDirection(&guard)
	if guard.x == initialPosition.x && guard.y == initialPosition.y {
		// you can't place an obstacle on the guards initial position
		return 0
	}

	if areas[guard.y][guard.x].hasBeenStepedOn {
		// we can't add an obstacle where the guard has already stepped on
		return 0
	}

	newArea := make([][]Area, len(areas))
	for i, row := range areas {
		newRow := make([]Area, len(row))
		copy(newRow, row)
		newArea[i] = newRow
	}

	newObstacles := make([][]bool, len(obstacles))
	for i, row := range obstacles {
		newRow := make([]bool, len(row))
		copy(newRow, row)
		newObstacles[i] = newRow
	}
	newObstacles[guard.y][guard.x] = true
	decrementDirection(&guard)

	return walkTheGuard(initialPosition, guard, newArea, newObstacles, obstaclesToPlace)
}

func isAtEdge(p Guard, width, height int) bool {
	return (p.x <= 0 && p.direction == Left) || (p.x >= width-1 && p.direction == Right) || (p.y <= 0 && p.direction == Up) || (p.y >= height-1 && p.direction == Down)
}
