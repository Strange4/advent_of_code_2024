package days

import (
	"Strange4/adventofcode2024/days/day1"
	"Strange4/adventofcode2024/days/day2"
	"Strange4/adventofcode2024/days/day3"
	"Strange4/adventofcode2024/days/day4"
	"Strange4/adventofcode2024/days/day5"
	"Strange4/adventofcode2024/days/day6"
)

func RunDay(day int) {
	daysMap := map[int]func(){
		1: day1.Run,
		2: day2.Run,
		3: day3.Run,
		4: day4.Run,
		5: day5.Run,
		6: day6.Run,
	}
	daysMap[day]()
}
