package day3

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	input := util.ReadLines("./inputs/day3.txt")[0]
	p1 := part1(input)
	fmt.Println(p1)
	part2()
}

func part1(input string) uint64 {
	mulRegex, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	util.Check(err)
	var sum uint64 = 0
	matches := mulRegex.FindAllStringSubmatch(input, -1)
	for _, matchInfo := range matches {
		firstNumber, err := strconv.ParseUint(matchInfo[1], 10, 16)
		util.Check(err)
		secondNumber, err := strconv.ParseUint(matchInfo[2], 10, 16)
		util.Check(err)
		sum += firstNumber * secondNumber
	}
	return sum
}

func part2() {
	input := util.ReadLines("./inputs/day3.txt")[0]
	goodMul, err := regexp.Compile(`^.+?don't\(\)|do\(\).+?don't\(\)`)
	util.Check(err)
	matches := goodMul.FindAllString(input, -1)
	var sum uint64 = 0
	for _, match := range matches {
		sum += part1(match)
	}
	fmt.Println(sum)
}
