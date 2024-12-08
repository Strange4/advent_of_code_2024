package main

import (
	"Strange4/adventofcode2024/days"
	"Strange4/adventofcode2024/days/util"
	"flag"
	"fmt"
	"os"
)

func main() {
	day := flag.Int("day", 4, "the day to run")
	flag.Parse()
	if *day <= 0 {
		fmt.Println("The days must be 1 days max")
		os.Exit(1)
	}
	fmt.Println("Running day ", *day)
	duration := util.TestFunc(func() {
		days.RunDay(*day)
	})
	fmt.Println("Running day", day, "took: ", duration)
}
