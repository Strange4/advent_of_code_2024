package day4

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"strings"
)

func Run() {
	part2()
	part1()
}

func part1() {
	input := util.ReadFile("./inputs/day4.txt")
	lines := util.ReadLines("./inputs/day4.txt")
	v2 := util.TestFunc(func() {
		sum := actualyDoingIt(lines)
		fmt.Println()
		fmt.Println("Normal way:", sum)
	})
	fmt.Println("Execution time took: ", v2)
	v1 := util.TestFunc(func() {
		lineLength := len(lines[0])
		allPatterns := createAllPatterns(lineLength)
		sum := findOccurances(allPatterns, input)
		fmt.Println("With Funky Regexes:", sum)
	})
	fmt.Println("Execution time took: ", v1)
}

func actualyDoingIt(lines []string) int {
	lineMap := make([][]string, len(lines))
	mapOutput := make([][]string, len(lines))
	for i, line := range lines {
		lineMap[i] = strings.Split(line, "")
	}

	for y := 0; y < len(lineMap); y++ {
		l := len(lineMap[0])
		mapOutput[y] = make([]string, l)
		for x := 0; x < l; x++ {
			mapOutput[y][x] = "."
		}
	}
	sum := 0
	for x := 0; x < len(lineMap); x++ {
		for y := 0; y < len(lines); y++ {
			diagonal := checkDiagonal(x, y, len(lines[0]), lineMap, mapOutput)
			horizontal := checkHorizontal(x, y, len(lines[0]), lineMap, mapOutput)
			vertical := checkVertical(x, y, lineMap, mapOutput)
			sum += diagonal + horizontal + vertical
		}
	}
	// for x := 0; x < len(lineMap[0]); x++ {
	// 	for y := 0; y < len(lineMap); y++ {
	// 		fmt.Print(mapOutput[x][y])
	// 	}
	// 	fmt.Print("\n")
	// }
	return sum
}

func checkVertical(x, y int, lineMap [][]string, mapOutput [][]string) int {
	if y > len(lineMap)-len("XMAS") {
		return 0
	}
	indeces := [][]int{{x, y}, {x, y + 1}, {x, y + 2}, {x, y + 3}}
	xmas := ""
	m := lineMap
	for _, index := range indeces {
		xmas += m[index[1]][index[0]]
	}
	if checkBothXMAS(xmas) {
		for _, index := range indeces {
			char := m[index[1]][index[0]]
			mapOutput[index[1]][index[0]] = char
		}
		return 1
	}
	return 0
}

func checkDiagonal(x, y, lineLength int, lineMap [][]string, mapOutput [][]string) int {
	// we don't have any size underneath to check for diagonal
	if y > len(lineMap)-len("XMAS") {
		return 0
	}
	sum := 0
	// checking / diagonal
	if x >= len("XMAS")-1 {
		m := lineMap
		indeces := [][]int{{x, y}, {x - 1, y + 1}, {x - 2, y + 2}, {x - 3, y + 3}}
		xmas := ""
		for _, index := range indeces {
			xmas += m[index[1]][index[0]]
		}
		if checkBothXMAS(xmas) {
			for _, index := range indeces {
				char := m[index[1]][index[0]]
				mapOutput[index[1]][index[0]] = char
			}
			sum += 1
		}
	}
	// checking \ diagonal
	if x <= lineLength-len("XMAS") {
		m := lineMap
		indeces := [][]int{{x, y}, {x + 1, y + 1}, {x + 2, y + 2}, {x + 3, y + 3}}
		xmas := ""
		for _, index := range indeces {
			xmas += m[index[1]][index[0]]
		}

		if checkBothXMAS(xmas) {
			for _, index := range indeces {
				char := m[index[1]][index[0]]
				mapOutput[index[1]][index[0]] = char
			}
			sum += 1
		}
	}
	return sum
}

func checkHorizontal(x, y, lineLength int, lineMap [][]string, mapOutput [][]string) int {
	if x > lineLength-len("XMAS") {
		return 0
	}
	m := lineMap
	indeces := [][]int{{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y}}
	xmas := ""
	for _, index := range indeces {
		xmas += m[index[1]][index[0]]
	}
	if checkBothXMAS(xmas) {
		for _, index := range indeces {
			char := m[index[1]][index[0]]
			mapOutput[index[1]][index[0]] = char
		}
		return 1
	}
	return 0
}

func checkBothXMAS(xmasLike string) bool {
	if xmasLike == "XMAS" || xmasLike == "SAMX" {
		return true
	}
	return false
}
