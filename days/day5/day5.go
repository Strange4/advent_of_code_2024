package day5

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	part1()
}

// go will have generic type aliases in 2025, this is advent of code 2024 :(
// type Set[K comparable] = map[K]bool
type Set = map[int]bool

func part1() {
	// 	input := `47|53
	// 97|13
	// 97|61
	// 97|47
	// 75|29
	// 61|13
	// 75|53
	// 29|13
	// 97|29
	// 53|29
	// 61|53
	// 97|53
	// 61|29
	// 47|13
	// 75|47
	// 97|75
	// 47|61
	// 75|61
	// 47|29
	// 75|13
	// 53|13

	// 75,47,61,53,29
	// 97,61,53,29,13
	// 75,29,13
	// 75,97,47,61,53
	// 61,13,29
	// 97,13,75,29,47`
	// input = util.CleanInput(input)
	input := util.ReadFile("./inputs/day5.txt")
	sections := strings.Split(input, "\n\n")
	rulesSection := sections[0]
	updatesSections := sections[1]
	rules := make(map[int]Set, len(rulesSection))
	for _, rule := range strings.Split(rulesSection, "\n") {
		split := strings.Split(rule, "|")
		x, err := strconv.Atoi(split[0])
		util.Check(err)
		y, err := strconv.Atoi(split[1])
		util.Check(err)
		xRules, ok := rules[x]
		if !ok {
			xRules = make(Set)
			rules[x] = xRules
		}
		xRules[y] = true
	}

	sum := 0
	for _, update := range strings.Split(updatesSections, "\n") {
		numbers := util.Map(strings.Split(update, ","), func(n string) int {
			parsed, err := strconv.Atoi(n)
			util.Check(err)
			return parsed
		})

		if updateIsValid(numbers, rules) {
			middle := len(numbers) / 2
			sum += numbers[middle]
		}
	}
	fmt.Println("Part 1: ", sum)
}

func updateIsValid(update []int, rules map[int]Set) bool {
	// we're checking in reverse to see if there are any numbers
	// that are out of place
	for i := len(update) - 1; i >= 0; i-- {
		numberRules := rules[update[i]]
		if numberRules == nil {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			// out of place
			if numberRules[update[j]] {
				return false
			}
		}
	}
	return true
}
