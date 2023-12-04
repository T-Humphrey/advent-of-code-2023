package main

import (
	"os"

	"github.com/t-humphrey/advent-of-code-2023/day_1"
	"github.com/t-humphrey/advent-of-code-2023/day_2"
	"github.com/t-humphrey/advent-of-code-2023/day_3"
)

func main() {
	day := os.Args[1]
	switch day {
	case "day1":
		if os.Args[2] == "1" {
			day_1.RunPart1()
		} else {
			day_1.RunPart2()
		}
	case "day2":
		if os.Args[2] == "1" {
			day_2.RunPart1()
		} else {
			day_2.RunPart2()
		}
	case "day3":
		if os.Args[2] == "1" {
			day_3.RunPart1()
		}
	}
}
