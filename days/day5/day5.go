package day5

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"strconv"
	"strings"
)

// go will have generic type aliases in 2025, this is advent of code 2024 :(
// type Set[K comparable] = map[K]bool
type Set = map[int]bool

func Run() {
	part1()
	part2()
}

func part1() {
	input := util.ReadFile("./inputs/day5.txt")

	rules, updates := getRulesAndUpdates(input)
	sum := 0
	for _, numbers := range updates {
		if updateIsValid(numbers, rules) {
			middle := len(numbers) / 2
			sum += numbers[middle]
		}
	}
	fmt.Println("Part 1: ", sum)
}

func part2() {
	input := util.ReadFile("./inputs/day5.txt")
	rules, updates := getRulesAndUpdates(input)

	sum := 0
	for _, update := range updates {
		if !updateIsValid(update, rules) {

			rightOrder := makeTheRightOrder(update, rules)

			middle := len(rightOrder) / 2
			sum += rightOrder[middle]
		}
	}
	fmt.Println("Part 2:", sum)
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

func makeTheRightOrder(numbers []int, rules map[int]Set) []int {
	rightOrder := make([]int, 0, len(numbers))

	numbersThatMustGoBeforePreviousNumber := make(Set)
	previousNumber := -1
	for _, n := range numbers {

		// initializing the map
		numbersThatMustGoBeforePreviousNumber[n] = false

		rulesForThisNumber := rules[n]

		hasOneAfterContraint := false
		for _, anotherNumber := range numbers {
			if n == anotherNumber {
				continue
			}
			// this "another number" must go after the previous number
			if rulesForThisNumber[anotherNumber] {
				hasOneAfterContraint = true
				break
			}
		}

		// this number has no other number that should be after it
		if !hasOneAfterContraint {
			previousNumber = n
		}
	}

	if previousNumber == -1 {
		panic("Couldn't find a last number")
	}

	rightOrder = append(rightOrder, previousNumber)
	delete(numbersThatMustGoBeforePreviousNumber, previousNumber)

	for len(numbersThatMustGoBeforePreviousNumber) > 1 {

		noneOfTheRemainingNumbersHaveRules := true
		for number := range numbersThatMustGoBeforePreviousNumber {
			// this number must go before previous number
			if rules[number][previousNumber] {
				numbersThatMustGoBeforePreviousNumber[number] = true
				noneOfTheRemainingNumbersHaveRules = false
			}
		}

		// all the numbers that remain have no rules about
		// the previous number, so all of them could potentially be
		// before the previous number
		if noneOfTheRemainingNumbersHaveRules {
			for number := range numbersThatMustGoBeforePreviousNumber {
				numbersThatMustGoBeforePreviousNumber[number] = true
			}
		}

		allOfThemMustGoBeforeEachOther := true
		for number, itShouldGoBefore := range numbersThatMustGoBeforePreviousNumber {
			if itShouldGoBefore {
				for anotherNumber, itShouldAlsoGoBefore := range numbersThatMustGoBeforePreviousNumber {
					if number == anotherNumber {
						continue
					}
					// if this other number must go after this number
					// the other number has priority
					if itShouldAlsoGoBefore && rules[number][anotherNumber] {
						numbersThatMustGoBeforePreviousNumber[number] = false
						allOfThemMustGoBeforeEachOther = false
						break
					}
				}
			}
		}

		if allOfThemMustGoBeforeEachOther {
			panic("A solution is not possible since all of them must go before each other")
		}

		for number, itShouldGoBefore := range numbersThatMustGoBeforePreviousNumber {
			if itShouldGoBefore {
				rightOrder = append(rightOrder, number)
				delete(numbersThatMustGoBeforePreviousNumber, number)
				previousNumber = number
			}
		}
	}

	// there should only be 1 left
	for number := range numbersThatMustGoBeforePreviousNumber {
		rightOrder = append(rightOrder, number)
	}

	return rightOrder
}

func getRulesAndUpdates(input string) (map[int]Set, [][]int) {
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

	updateLines := strings.Split(updatesSections, "\n")
	updates := make([][]int, len(updateLines))
	for i, line := range updateLines {
		numbers := util.Map(strings.Split(line, ","), func(n string) int {
			parsed, err := strconv.Atoi(n)
			util.Check(err)
			return parsed
		})
		updates[i] = numbers
	}

	return rules, updates
}
