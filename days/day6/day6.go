package day6

import (
	"fmt"
	"strings"
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
	direction       Direction // if it was stepped on, what direction where you going?
}

func part1And2() {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	lines := strings.Split(input, "\n")
	// lines := util.ReadLines("./inputs/day6.txt")

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

	possibleLoops := walkTheGuard(currentPos, visitedAreas, obstacles, 1)

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

func walkTheGuard(currentPos Guard, visitedAreas [][]Area, obstacles [][]bool, obstaclesToPlace int) int {
	possibleLoops := 0

	for {
		// for ;!obstacles[currentPos.y][currentPos.x] && !isAtEdge(currentPos, len(obstacles[0]), len(obstacles));
		if currentPos.direction == Right {
			for ; !obstacles[currentPos.y][currentPos.x] && !isAtEdge(currentPos, len(obstacles[0]), len(obstacles)); currentPos.x++ {
				area := visitedAreas[currentPos.y][currentPos.x]

				// loop detected
				if area.hasBeenStepedOn && area.direction == currentPos.direction {
					return possibleLoops + 1
				}

				if obstaclesToPlace > 0 {
					newLoops := canCreateLoop(currentPos, visitedAreas, obstacles, obstaclesToPlace-1)

					possibleLoops += newLoops
				}
				visitedAreas[currentPos.y][currentPos.x] = Area{true, Right}
			}
			if !obstacles[currentPos.y][currentPos.x] && isAtEdge(currentPos, len(obstacles[0]), len(obstacles)) {
				return possibleLoops
			}
			currentPos.x--
			currentPos.direction = Down
			currentPos.y++
		} else if currentPos.direction == Left {
			for ; !obstacles[currentPos.y][currentPos.x] && !isAtEdge(currentPos, len(obstacles[0]), len(obstacles)); currentPos.x-- {
				area := visitedAreas[currentPos.y][currentPos.x]

				// loop detected
				if area.hasBeenStepedOn && area.direction == Left {
					return possibleLoops + 1
				}
				if obstaclesToPlace > 0 {
					newLoops := canCreateLoop(currentPos, visitedAreas, obstacles, obstaclesToPlace-1)
					possibleLoops += newLoops
				}

				visitedAreas[currentPos.y][currentPos.x] = Area{true, Left}
			}
			if !obstacles[currentPos.y][currentPos.x] && isAtEdge(currentPos, len(obstacles[0]), len(obstacles)) {
				return possibleLoops
			}
			currentPos.x++
			currentPos.direction = Up
			currentPos.y--
		} else if currentPos.direction == Up {
			for ; !obstacles[currentPos.y][currentPos.x] && !isAtEdge(currentPos, len(obstacles[0]), len(obstacles)); currentPos.y-- {
				area := visitedAreas[currentPos.y][currentPos.x]

				// loop detected
				if area.hasBeenStepedOn && area.direction == Up {
					return possibleLoops + 1
				}
				if obstaclesToPlace > 0 {
					newLoops := canCreateLoop(currentPos, visitedAreas, obstacles, obstaclesToPlace-1)

					possibleLoops += newLoops
				}
				visitedAreas[currentPos.y][currentPos.x] = Area{true, Up}
			}
			if !obstacles[currentPos.y][currentPos.x] && isAtEdge(currentPos, len(obstacles[0]), len(obstacles)) {
				return possibleLoops
			}
			currentPos.y++
			currentPos.direction = Right
			currentPos.x++
		} else if currentPos.direction == Down {
			for ; !obstacles[currentPos.y][currentPos.x] && !isAtEdge(currentPos, len(obstacles[0]), len(obstacles)); currentPos.y++ {
				area := visitedAreas[currentPos.y][currentPos.x]

				// loop detected
				if area.hasBeenStepedOn && area.direction == Down {
					return possibleLoops + 1
				}
				if obstaclesToPlace > 0 {
					newLoops := canCreateLoop(currentPos, visitedAreas, obstacles, obstaclesToPlace-1)

					possibleLoops += newLoops
				}
				visitedAreas[currentPos.y][currentPos.x] = Area{true, Down}
			}
			if !obstacles[currentPos.y][currentPos.x] && isAtEdge(currentPos, len(obstacles[0]), len(obstacles)) {
				return possibleLoops
			}
			currentPos.y--
			currentPos.direction = Left
			currentPos.x--
		}
	}
}

// func incrementDirection(g *Guard) {
// 	if g.direction == Right {
// 		g.x++
// 	} else
// }

func canCreateLoop(p Guard, areas [][]Area, obstacles [][]bool, obstaclesToPlace int) int {
	// checking if we create an obstacle in front would create a loop

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
	if p.direction == Left {
		// if there is already an obstacle there
		// this means that we can't place one here
		if obstacles[p.y][p.x-1] {
			return 0
		}
		newObstacles[p.y][p.x-1] = true
		return walkTheGuard(p, newArea, newObstacles, obstaclesToPlace)
	} else if p.direction == Right {
		if obstacles[p.y][p.x+1] {
			return 0
		}
		newObstacles[p.y][p.x+1] = true
		return walkTheGuard(p, newArea, newObstacles, obstaclesToPlace)
	} else if p.direction == Up {
		if obstacles[p.y-1][p.x] {
			return 0
		}
		newObstacles[p.y-1][p.x] = true
		return walkTheGuard(p, newArea, newObstacles, obstaclesToPlace)
	} else if p.direction == Down {
		if obstacles[p.y+1][p.x] {
			return 0
		}
		newObstacles[p.y+1][p.x] = true
		return walkTheGuard(p, newArea, newObstacles, obstaclesToPlace)
	}

	panic("This a direction that I can't handle")
}

func isAtEdge(p Guard, width, height int) bool {
	return (p.x <= 0 && p.direction == Left) || (p.x >= width-1 && p.direction == Right) || (p.y <= 0 && p.direction == Up) || (p.y >= height-1 && p.direction == Down)
}
