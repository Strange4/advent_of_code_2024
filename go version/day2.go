package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := readLines("../inputs/day2.txt")
	safeReports := 0
	for _, line := range lines {
		numbersAsInt := stringsToInt(strings.Split(line, " "))
		isSafe := numbersAreSafe(numbersAsInt)
		if isSafe {
			safeReports++
		}
	}
	fmt.Printf("Part 1: %v\n", safeReports)
}

func part2() {
	lines := readLines("../inputs/day2.txt")
	safeReports := 0
	for _, line := range lines {
		numbersAsInt := stringsToInt(strings.Split(line, " "))
		isSafe := numbersAreSafe(numbersAsInt)
		if isSafe {
			safeReports++
		} else {
			// since there are only 5 more possibilities to check
			// lets just make them all and check for all of them
			possibilities := createPossibleRemovals(numbersAsInt)
			for _, numbers := range possibilities {
				// this means that there is at least one possibility that
				// is safe
				if numbersAreSafe(numbers) {
					safeReports++
					break
				}
			}
		}
	}
	fmt.Printf("Part 2: %v\n", safeReports)
}

func createPossibleRemovals(numbers []int64) [][]int64 {
	possibilities := make([][]int64, 0, len(numbers))
	for i := 0; i < len(numbers); i++ {
		newCopy := make([]int64, len(numbers))
		copy(newCopy, numbers)
		// from begining to the one we want to remove (not inclusive), and then from the next to the end
		withOneRemoved := append(newCopy[:i], newCopy[i+1:]...)
		possibilities = append(possibilities, withOneRemoved)
	}
	return possibilities
}

func numbersAreSafe(numbers []int64) bool {
	var isIncreasing = (numbers[0] - numbers[1]) > 0

	isSafe := true
	for i := 0; i < len(numbers)-1; i++ {
		this := numbers[i]
		next := numbers[i+1]
		diff := this - next
		absDiff := absInt(diff)
		if isIncreasing != (diff > 0) || absDiff == 0 || absDiff > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func stringsToInt(strings []string) []int64 {
	asInt := make([]int64, 0, len(strings))
	for _, s := range strings {
		number, err := strconv.ParseInt(s, 10, 8)
		check(err)
		asInt = append(asInt, number)
	}
	return asInt
}

func readLines(path string) []string {
	data, err := os.ReadFile(path)
	check(err)
	asString := string(data)
	lines := strings.Split(strings.ReplaceAll(asString, "\r\n", "\n"), "\n")
	return lines
}

func absInt(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
