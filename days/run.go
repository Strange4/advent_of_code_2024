package days

import (
	"Strange4/adventofcode2024/days/day1"
	"Strange4/adventofcode2024/days/day2"
	"Strange4/adventofcode2024/days/day3"
	"Strange4/adventofcode2024/days/day4"
	"Strange4/adventofcode2024/days/day5"
	"Strange4/adventofcode2024/days/day6"
	"Strange4/adventofcode2024/days/day7"
	"Strange4/adventofcode2024/days/day8"
)

func RunDay(day int) {
	daysMap := map[int]func(){
		1: day1.Run,
		2: day2.Run,
		3: day3.Run,
		4: day4.Run,
		5: day5.Run,
		6: day6.Run,
		7: day7.Run,
		8: day8.Run,
	}
	daysMap[day]()
}
