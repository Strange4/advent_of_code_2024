package day9

import (
	"Strange4/adventofcode2024/days/util"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	input := util.ReadFile("./inputs/day9.txt")

	diskMap := util.Map(strings.Split(input, ""), func(s string) uint8 {
		number, err := strconv.Atoi(s)
		util.Check(err)
		return uint8(number)
	})
	part1(diskMap)
	part2(diskMap)
}

func part1(diskMap []uint8) {

	// disk data is 2 numbers to space space, first the id, then the length of the file
	var diskData []uint

	filePtr2 := len(diskMap) - 1

	for filePtr1 := 0; filePtr1 <= filePtr2; filePtr1++ {
		// add file to data
		if filePtr1%2 == 0 {
			id := filePtr1 / 2
			diskData = append(diskData, uint(id), uint(diskMap[filePtr1]))
		} else {
			// fill the space with the rightmost file
			for filePtr1 < filePtr2 {
				spaceAvailable := diskMap[filePtr1]
				fileLength := diskMap[filePtr2]
				availableToFill := util.Min(spaceAvailable, fileLength)

				diskData = append(diskData, uint(filePtr2/2), uint(availableToFill))
				spaceAvailable -= availableToFill
				fileLength -= availableToFill

				// the file is completely transfered
				if fileLength == 0 {
					filePtr2 -= 2
				} else {
					// write the remaining file to the disk map
					diskMap[filePtr2] = fileLength
				}
				if spaceAvailable == 0 {
					break
				} else {
					diskMap[filePtr1] = spaceAvailable
				}
			}
		}
	}
	fmt.Println("Part 1: ", computeCheckSum(diskData))
}

func part2(diskMap []uint8) {

	numberOfFiles := (len(diskMap) / 2) + 1

	// the disk data is id, length of file, length of space
	diskData := make([]uint, numberOfFiles*3)

	diskMap = append(diskMap, 0) // append a 0 length space at the end for parity
	for i := 0; i < len(diskMap); i += 2 {
		id := uint(i / 2)
		dataStart := id * 3
		diskData[dataStart] = id
		diskData[dataStart+1] = uint(diskMap[i])
		diskData[dataStart+2] = uint(diskMap[i+1])
	}

	for filePtr := len(diskData) - 3; filePtr > 0; filePtr -= 3 {
		// if this doesn't work try until the end
		for spacePtr := 2; spacePtr < filePtr; spacePtr += 3 {
			if diskData[spacePtr] >= diskData[filePtr+1] {
				moveWholeFile(diskData, spacePtr, filePtr)
				// since we moved the file the next file to move is
				// at the same position
				filePtr += 3 // setting up for the next loop
				break
			}
		}
	}
	fmt.Println("Part 2:", checkSumPart2(diskData))
}

func moveWholeFile(diskData []uint, spacePtr, filePtr int) {
	// add the space left after this file is moved to before this file
	diskData[filePtr-1] += diskData[filePtr+2] + diskData[filePtr+1]

	// the space after this one
	spaceLeftAfterShift := diskData[spacePtr] - diskData[filePtr+1]
	// the space before the shift should be 0 since it will be moved after
	diskData[spacePtr] = 0

	id := diskData[filePtr]
	size := diskData[filePtr+1]

	// deleting file at position
	copy(diskData[filePtr:], diskData[filePtr+3:])

	// shifting the slice to make space
	copy(diskData[spacePtr+3:], diskData[spacePtr:])

	// setting the file data
	diskData[spacePtr+1] = id
	diskData[spacePtr+2] = size
	diskData[spacePtr+3] = spaceLeftAfterShift
}

func computeCheckSum(diskData []uint) uint {
	var sum uint
	var position uint
	for i := 0; i < len(diskData); i += 2 {
		id := diskData[i]
		fileLength := diskData[i+1]
		// since 1 * 2 + 2 * 2 + 3 * 2 == 2 * (1 + 2 + 3)
		// we can just multiply the id by the sum of integers between the start and end positions
		start := position
		end := position + fileLength - 1
		sumOfIntegers := sumOfIntegers(start, end)
		sum += sumOfIntegers * id
		position = end + 1
	}
	return sum
}

func checkSumPart2(diskData []uint) uint {
	// almost like the first checksum but skip the position of the
	// empty spaces
	var sum uint
	var position uint
	for i := 0; i < len(diskData); i += 3 {
		id := diskData[i]
		fileLength := diskData[i+1]
		spaceAfter := diskData[i+2]
		start := position
		end := position + fileLength - 1
		sumOfIntegers := sumOfIntegers(start, end)
		sum += sumOfIntegers * id
		position = end + spaceAfter + 1
	}
	return sum
}

// of the integers between start to end inclusive
func sumOfIntegers(start, end uint) uint {
	return (end - start + 1) * (start + end) / 2
}
