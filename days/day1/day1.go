package day1

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	part1()
	part2()
}

func part1() {
	leftList, rightList := getList()
	slices.Sort(leftList)
	slices.Sort(rightList)
	var sum int64 = 0
	for i := 0; i < len(leftList); i++ {
		diff := leftList[i] - rightList[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	fmt.Println("Part 1: ", sum)
}

func part2() {
	leftList, rightList := getList()
	occurances := make(map[int64]int64)
	for _, number := range rightList {
		occurances[number] += 1
	}
	var sum int64 = 0
	for _, number := range leftList {
		sum += number * occurances[number]
	}
	fmt.Println("Part 2: ", sum)
}

func getList() ([]int64, []int64) {
	lines := util.ReadLines("./inputs/day1.txt")
	leftList := make([]int64, len(lines))
	rightList := make([]int64, len(lines))
	for i, line := range lines {
		numbers := strings.Split(line, "   ")
		a, err := strconv.ParseInt(numbers[0], 10, 64)
		util.Check(err)
		b, err := strconv.ParseInt(numbers[1], 10, 64)
		util.Check(err)
		leftList[i] = a
		rightList[i] = b
	}
	return leftList, rightList
}
