package day7

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Run() {
	part1And2()
}

func part1And2() {
	lines := util.ReadLines("./inputs/day7.txt")

	var part1 uint64 = 0
	var part2 uint64 = 0
	for _, line := range lines {
		sections := strings.Split(line, ": ")
		result, err := strconv.ParseUint(sections[0], 10, 64)
		util.Check(err)
		operands := util.Map(strings.Split(sections[1], " "), func(s string) uint16 {
			number, err := strconv.ParseUint(s, 10, 16)
			util.Check(err)
			return uint16(number)
		})

		result1 := calculate(result, operands)
		// if only addition and multiplication don't work
		// try using the third operator
		if result1 == 0 {
			part2 += calculate2(result, operands)
		}
		part1 += result1
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2+part1)
}

func calculate(result uint64, operands []uint16) uint64 {

	// 0 in the list represents add, 1 represents multiply
	operationList := 0

	numberOfOperations := len(operands) - 1

	// there are 2^(numbers - 1) possible operations
	possibleOperations := int(math.Pow(2, float64(numberOfOperations)))

	for operationList < possibleOperations {
		currentIndex := 1
		currentResult := uint64(operands[0])
		for pos := numberOfOperations - 1; pos >= 0; pos-- {

			// getting the next operation in the list
			// ex: 01 >> 1 -> 0 which means that the first operation is an addition
			//     01 >> 0 -> 1 which means that the second operation is a multiplication
			operation := operationList >> pos
			operand := uint64(operands[currentIndex])

			if operation&1 == 0 {
				currentResult += operand
			} else {
				currentResult *= operand
			}
			currentIndex++
		}
		if result == currentResult {
			return currentResult
		}
		operationList++
	}
	return 0
}

func calculate2(result uint64, operands []uint16) uint64 {
	var countingBase uint8 = 3
	numberOfOperations := len(operands) - 1
	operationList := make([]uint8, numberOfOperations)
	currentIncrementIndex := len(operationList) - 1

	// there are 3^(numbers - 1) possible operations
	possibleOperations := int(math.Pow(float64(countingBase), float64(numberOfOperations)))

	for i := 0; i < possibleOperations; i++ {

		currentResult := uint64(operands[0])
		for index, operation := range operationList {
			// the operand is always the next one on the list
			// since we start with the first one
			operand := uint64(operands[index+1])

			if operation == 0 {
				currentResult += operand
			} else if operation == 1 {
				currentResult *= operand
			} else {
				// contatenation
				numberOfDigits := uint64(math.Floor(math.Log10(float64(operand)) + 1))
				currentResult *= uint64(math.Pow10(int(numberOfDigits)))
				currentResult += operand
			}
		}
		if uint64(result) == currentResult {
			return currentResult
		}

		// we need to increment the next one

		rollOverNumber := countingBase - 1
		if operationList[numberOfOperations-1] == rollOverNumber {
			for j := numberOfOperations - 1; j >= 0; j-- {
				if operationList[j] == rollOverNumber {
					operationList[j] = 0
				} else {
					operationList[j]++
					break
				}
			}
		} else {
			operationList[currentIncrementIndex]++
		}
	}
	return 0
}
