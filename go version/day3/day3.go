package main

import (
	"adventofcode2024/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// part1()
	// part25()
	part2()
}

func part25() {
	input := util.ReadLines("../../inputs/day3.txt")[0]
	invalid, err := regexp.Compile(`don't\(\).+?(do\(\)|\z)`)
	util.Check(err)
	valid := invalid.ReplaceAllString(input, "")
	regex, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	util.Check(err)
	matches := regex.FindAllStringSubmatch(valid, -1)
	var sum uint64 = 0
	for _, matchInfo := range matches {
		firstNumber, err := strconv.ParseUint(matchInfo[1], 10, 16)
		util.Check(err)
		secondNumber, err := strconv.ParseUint(matchInfo[2], 10, 16)
		util.Check(err)
		fmt.Println(matchInfo[0])
		sum += firstNumber * secondNumber
	}
	fmt.Println(sum)
}

func part2() {
	// input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	input := util.ReadLines("../../inputs/day3.txt")[0]
	mulRegex, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	util.Check(err)
	doRegex, err := regexp.Compile(`do\(\)`)
	util.Check(err)
	dontRegex, err := regexp.Compile(`don't\(\)`)
	util.Check(err)
	mulIndeces := mulRegex.FindAllStringSubmatchIndex(input, -1)
	doIndeces := doRegex.FindAllStringIndex(input, -1)
	dontIndeces := dontRegex.FindAllStringIndex(input, -1)
	doPositions := indecesToPositions(doIndeces)
	dontPositions := indecesToPositions(dontIndeces)
	var sum uint64 = 0

	for _, matchdata := range mulIndeces {
		startIndex := matchdata[0]
		previousDo := positionInSliceStartingFromRight(func(elem int) bool { return elem < startIndex }, doPositions)
		previousDont := positionInSliceStartingFromRight(func(elem int) bool { return elem < startIndex }, dontPositions)
		// the latest do is after the latest don't
		if previousDo == -1 && previousDont == -1 {
			// fmt.Println("case1", startIndex)
			sum += parseAndMultiply(input, matchdata)
			fmt.Println(input[matchdata[0]:matchdata[1]])

		} else if previousDo != -1 && previousDont == -1 {
			// fmt.Println("case2", startIndex)
			fmt.Println(input[matchdata[0]:matchdata[1]])

			sum += parseAndMultiply(input, matchdata)
		} else if previousDo != -1 && previousDont != -1 && doPositions[previousDo] > dontPositions[previousDont] {
			// fmt.Println("case3", startIndex, "dos and don'ts", doPositions[previousDo], dontPositions[previousDont])
			// fmt.Println("Parse and multiply")
			fmt.Println(input[matchdata[0]:matchdata[1]])
			sum += parseAndMultiply(input, matchdata)
		}
	}
	fmt.Println(sum)
}

func parseAndMultiply(input string, matchData []int) uint64 {
	// the match data always has the same format
	// 3 pairs of ranges that represent the match start-end, group 1 start-end, group 2 start-end
	// the first group is the first number, the second group is the second number
	firstGroupString := input[matchData[2]:matchData[3]]
	firstNumber, err := strconv.ParseUint(firstGroupString, 10, 16)
	util.Check(err)
	secondGroupString := input[matchData[4]:matchData[5]]
	secondNumber, err := strconv.ParseUint(secondGroupString, 10, 16)
	util.Check(err)
	return firstNumber * secondNumber
}

func positionInSliceStartingFromRight(predicate func(elem int) bool, slice []int) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return i
		}
	}
	return -1
}

func indecesToPositions(indeces [][]int) []int {
	positions := make([]int, 0, len(indeces))
	for _, element := range indeces {
		positions = append(positions, element[1])
	}
	return positions
}

func part1() {
	// input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	input := util.ReadLines("../../inputs/day3.txt")[0]
	regex, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	util.Check(err)
	matches := regex.FindAllStringSubmatch(input, -1)
	var sum uint64 = 0
	for _, matchInfo := range matches {
		firstNumber, err := strconv.ParseUint(matchInfo[1], 10, 16)
		util.Check(err)
		secondNumber, err := strconv.ParseUint(matchInfo[2], 10, 16)
		util.Check(err)
		sum += firstNumber * secondNumber
	}
	fmt.Println(sum)
}
