package day4

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
)

func findOccurances(patterns []*regexp2.Regexp, input string) int {
	input = strings.Replace(input, "\r\n", "\n", -1)
	sum := 0
	emptyMap := ""
	for _, line := range strings.Split(input, "\n") {
		emptyMap += strings.Repeat(".", len(line)) + "\n"
	}
	for _, pattern := range patterns {
		for match, err := pattern.FindStringMatch(input); match != nil; match, err = pattern.FindNextMatch(match) {
			util.Check(err)
			sum++
			groups := match.Groups()
			for x := 1; x < len(groups); x++ {
				group := groups[x]
				a := emptyMap[:group.Index]
				b := string(input[group.Index])
				c := emptyMap[group.Index+1:]
				emptyMap = a + b + c
			}
		}
	}
	// fmt.Print(emptyMap)
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
		// we need to use positive look ahead so that the characters aren't consumed
		// this is why we need the other package
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
