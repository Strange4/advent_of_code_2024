package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

func main() {
	// 	input := `MMMSXXMASM
	// MSAMXMSMSA
	// AMXSXMAAMM
	// MSAMASMSMX
	// XMASAMXAMM
	// XXAMMXXAMA
	// SMSMSASXSS
	// SAXAMASAAA
	// MAMMMXMMMM
	// MXMXAXMASX`
	// input = strings.ReplaceAll(input, "\r\n", "\n")
	// lines := strings.Split(input, "\n")
	input := util.ReadFile("../../inputs/day4.txt")
	lines := util.ReadLines("../../inputs/day4.txt")
	lineLength := len(lines[0])
	start := time.Now()
	allPatterns := createAllPatterns(lineLength)
	end := time.Now()
	fmt.Println("Compiled", len(allPatterns), "patterns for line length", lineLength, "in", end.Sub(start))
	v2 := util.TestFunc(func() {
		sum := actualyDoingIt(lines)
		fmt.Println()
		fmt.Println("V2:", sum)
	})
	fmt.Println("Execution time took: ", v2)
	v1 := util.TestFunc(func() {
		sum := findOccurances(allPatterns, input)
		fmt.Println("V1:", sum)
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
	for x := 0; x < len(lineMap[0]); x++ {
		for y := 0; y < len(lineMap); y++ {
			fmt.Print(mapOutput[x][y])
		}
		fmt.Print("\n")
	}
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
		// xmas := m[x][y] + m[x-1][y+1] + m[x-2][y+2] + m[x-3][y+3]
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

func findOccurances(patterns []*regexp2.Regexp, input string) int {
	input = strings.Replace(input, "\r\n", "\n", -1)
	sum := 0
	emptyMap := ""
	for _, line := range strings.Split(input, "\n") {
		emptyMap += strings.Repeat(".", len(line)) + "\n"
	}
	for _, pattern := range patterns {
		fmt.Println(pattern)
		// pattern.
		match, err := pattern.FindNextMatch()
		util.Check(err)
		
		groups := match.Groups()
		for x := 1; x < len(groups); x++ {
			group := groups[x]
			a := emptyMap[:group.Index]
			b := string(input[group.Index])
			c := emptyMap[group.Index+1:]
			emptyMap = a + b + c
		}
		// for _, pairs := range match.Groups() {
		// 	for x := 1; x < len(pairs); x += 2 {
		// 		groupIndex := pairs[x]
		// 		a := emptyMap[:groupIndex]
		// 		b := string(input[groupIndex])
		// 		c := emptyMap[groupIndex+1:]
		// 		emptyMap = a + b + c
		// 	}
		// }
		// occurances := pattern.FindAllString(input, -1)
		sum += len(match.)
	}
	fmt.Print(emptyMap)
	return sum
}

func createAllPatterns(lineLength int) []*regexp2.Regexp {
	options := regexp2.RegexOptions(regexp2.Multiline | regexp2.Compiled)
	upDownPatternLength := lineLength * 2
	diagonalPatternLength := lineLength * (lineLength - len("XMAS")*2 - 1)
	allPaternsLength := 1 + 1 + upDownPatternLength + diagonalPatternLength
	allPatterns := make([]*regexp2.Regexp, 0, allPaternsLength)

	leftToRight := regexp2.MustCompile("(X)(M)(A)(S)", options)
	rightToLeft := regexp2.MustCompile("(S)(A)(M)(X)", options)
	allPatterns = append(allPatterns, leftToRight, rightToLeft)
	for i := 0; i < lineLength; i++ {
		nbLettersBefore := i
		nbLettersAfter := lineLength - i - 1

		// we have to make this any to pass a variadic argument
		beforeAndAfter := make([]any, 0, len("XMAS")*2)
		for x := 0; x < len("XMAS"); x++ {
			beforeAndAfter = append(beforeAndAfter, nbLettersBefore, nbLettersAfter)
		}
		topDownPatternTemplate := `.{%v}(X)(?=.{%v}\n.{%v}(M).{%v}\n.{%v}(A).{%v}\n.{%v}(S).{%v})`
		bottomUpPatternTemplate := `.{%v}(S)(?=.{%v}\n.{%v}(A).{%v}\n.{%v}(M).{%v}\n.{%v}(X).{%v})`
		topDown := regexp2.MustCompile(fmt.Sprintf(topDownPatternTemplate, beforeAndAfter...), options)
		bottomUp := regexp2.MustCompile(fmt.Sprintf(bottomUpPatternTemplate, beforeAndAfter...), options)
		allPatterns = append(allPatterns, topDown, bottomUp)
		// diagonals

		// top right diagonal
		if nbLettersBefore+1-len("XMAS") >= 0 {

			trDiagonalBeforeAndAfter := make([]any, 0, len("XMAS")*2)
			for x := 0; x < len("XMAS"); x++ {
				trDiagonalBeforeAndAfter = append(trDiagonalBeforeAndAfter, nbLettersBefore-x, nbLettersAfter+x)
			}

			trToBL := regexp2.MustCompile(fmt.Sprintf(topDownPatternTemplate, trDiagonalBeforeAndAfter...), options)
			sameButBackwards := regexp2.MustCompile(fmt.Sprintf(bottomUpPatternTemplate, trDiagonalBeforeAndAfter...), options)
			allPatterns = append(allPatterns, trToBL, sameButBackwards)
		}

		// top left diagonal
		if nbLettersAfter+1-len("XMAS") >= 0 {
			tlDiagonalBeforeAndAfter := make([]any, 0, len("XMAS")*2)
			for x := 0; x < len("XMAS"); x++ {
				tlDiagonalBeforeAndAfter = append(tlDiagonalBeforeAndAfter, nbLettersBefore+x, nbLettersAfter-x)
			}
			tlToBR := regexp2.MustCompile(fmt.Sprintf(bottomUpPatternTemplate, tlDiagonalBeforeAndAfter...), options)
			sameButBackwards := regexp2.MustCompile(fmt.Sprintf(topDownPatternTemplate, tlDiagonalBeforeAndAfter...), options)
			allPatterns = append(allPatterns, tlToBR, sameButBackwards)
		}
	}
	return allPatterns
}
