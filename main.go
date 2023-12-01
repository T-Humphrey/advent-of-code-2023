package main

import (
	"os"

	"github.com/t-humphrey/advent-of-code-2023/week_1"
)

func main() {
	week := os.Args[1]
	switch week {
	case "week1":
		if os.Args[2] == "1" {
			week_1.RunPart1()
		} else {
			week_1.RunPart2()
		}
	}
}
